package main

import (
	"log"
	"net/http"
)

// Home handler
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from text.tk"))
}

func main() {
	mux := http.NewServeMux() // initialize new servemux
	mux.HandleFunc("/", home)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux) // start a new web server
	log.Fatal(err)
}
