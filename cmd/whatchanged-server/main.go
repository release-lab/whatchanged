package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/axetroy/whatchanged"
	"github.com/axetroy/whatchanged/option"
	"github.com/pkg/errors"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var (
		err    error
		output = bytes.NewBuffer([]byte{})
	)
	// cors
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "GET")

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

	username := query.Get("username")
	repo := query.Get("repo")
	version := query.Get("version")
	template := query.Get("template")
	preset := query.Get("preset")

	url := fmt.Sprintf("https://github.com/%s/%s", username, repo)

	if err = whatchanged.Generate(url, output, &option.Options{
		Version:  regexp.MustCompile(`\s+`).Split(version, -1),
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
