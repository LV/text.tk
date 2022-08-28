package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))         // serve files; path specified is relative to directory root
	mux.Handle("/static/", http.StripPrefix("/static", fileServer)) // register file server as handler for all URL paths with "/static/"

	mux.HandleFunc("/", home)
	mux.HandleFunc("/t/view", textView)
	mux.HandleFunc("/t/create", textCreate)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
