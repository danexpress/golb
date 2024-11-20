package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"

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

	postTemplate := template.Must(template.ParseFiles("post.gohtml"))

	mux.HandleFunc("GET /posts/{slug}", golb.PostHandler(
		golb.FileReader{},
		postTemplate,
	))

	err := http.ListenAndServe(":3030", mux)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
