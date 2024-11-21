package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/danexpress/golb"
)

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprint(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(args []string, Stdout io.Writer) error {
	mux := http.NewServeMux()

	postReader := golb.FileReader{}
	postTemplate := template.Must(template.ParseFiles("post.gohtml"))
	mux.HandleFunc("GET /posts/{slug}", golb.PostHandler(postReader, postTemplate))

	indexTemplate := template.Must(template.ParseFiles("index.gohtml"))
	mux.HandleFunc("GET /", golb.IndexHandler(postReader, indexTemplate))

	err := http.ListenAndServe(":3030", mux)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
