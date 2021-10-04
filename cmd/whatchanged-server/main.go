package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	"github.com/whatchanged-community/whatchanged"
	"github.com/whatchanged-community/whatchanged/option"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var (
		err    error
		output = bytes.NewBuffer([]byte{})
	)
	// cors
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", strings.Join([]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete}, " ,"))

	if r.Method == http.MethodOptions {
		w.WriteHeader(200)
		_, _ = w.Write([]byte{})
		return
	}

	defer func() {
		if r, ok := recover().(error); ok {
			err = r
		}

		if err != nil {
			b := []byte(fmt.Sprintf("%+v\n", err))
			w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write(b)
		} else {
			w.Header().Set("Content-Type", "text/markdown; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(output.Bytes())
		}
	}()

	query := r.URL.Query()

	repo := query.Get("repo")
	branch := query.Get("branch")
	version := query.Get("version")
	template := query.Get("template")
	preset := query.Get("preset")

	if templateDecoded, err := url.QueryUnescape(template); err == nil {
		template = templateDecoded
	}

	if branchDecoded, err := url.QueryUnescape(branch); err == nil {
		branch = branchDecoded
	}

	if err = whatchanged.Generate(r.Context(), repo, output, &option.Options{
		Version:  regexp.MustCompile(`\s+`).Split(version, -1),
		Branch:   branch,
		Template: template,
		Preset:   option.Preset(preset),
	}); err != nil {
		err = errors.WithStack(err)
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
