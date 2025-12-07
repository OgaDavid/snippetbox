package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}

	res.Write([]byte("Hello from SnippetBox"))
}

func showSnippet(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.URL.Query().Get("id"))

	if err != nil || id < 1 {
		http.NotFound(res, req)
		return
	}

	fmt.Fprintf(res, "Display a specific snippet with ID: %d", id)
}

func createSnippet(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		res.Header().Set("Allow", "POST")
		res.Header().Set("Content-Type", "application/json")
		http.Error(res, "Method not allowed!", http.StatusMethodNotAllowed)
		return
	}
	res.Write([]byte("Create a snippet"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)

	log.Fatal(err)
}
