package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	parser "github.com/axetroy/changelog/1_parser"
	extractor "github.com/axetroy/changelog/2_extractor"
	transformer "github.com/axetroy/changelog/3_transformer"
	generator "github.com/axetroy/changelog/4_generator"
	formatter "github.com/axetroy/changelog/5_formatter"
	"github.com/axetroy/changelog/internal/client"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/pkg/errors"
)

func generate(remote string, version string, templateStr string) ([]byte, error) {
	repo, err := git.CloneContext(context.Background(), memory.NewStorage(), nil, &git.CloneOptions{
		URL:      remote,
		Progress: os.Stderr,
	})

	if err != nil {
		return nil, err
	}

	c := &client.GitClient{
		Repository: repo,
	}

	scope, err := parser.Parse(c, version)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	splices, err := extractor.Extract(c, scope)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	ctxs, err := transformer.Transform(c, splices)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	output, err := generator.Generate(c, ctxs, "md", "default", templateStr)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	formattedOutput, err := formatter.Format(output, "md", "")

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return formattedOutput, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	// cors
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	query := r.URL.Query()

	username := query.Get("username")
	repo := query.Get("repo")
	version := query.Get("version")
	template := query.Get("template")

	url := fmt.Sprintf("https://github.com/%s/%s", username, repo)

	changelog, err := generate(url, version, template)

	if err != nil {
		b := []byte(fmt.Sprintf("%+v\n", err))
		w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
		w.WriteHeader(500)
		_, _ = w.Write(b)
	} else {
		w.Header().Set("Content-Type", "text/markdown; charset=UTF-8")
		w.WriteHeader(200)
		_, _ = w.Write(changelog)
	}
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", handler)
	print("Listen on port ", port, "\n")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
