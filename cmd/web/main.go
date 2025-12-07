package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/create", createSnippet)

	err := http.ListenAndServe(":4000", mux)

	if err != nil {
		log.Fatal(err)
	}
}