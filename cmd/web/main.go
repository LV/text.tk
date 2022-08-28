package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address") // default value of :4000 with short text description explaining what the flag controls

	flag.Parse() // parse command-line flag, must call this *before* using value of addr

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))         // serve files; path specified is relative to directory root
	mux.Handle("/static/", http.StripPrefix("/static", fileServer)) // register file server as handler for all URL paths with "/static/"

	mux.HandleFunc("/", home)
	mux.HandleFunc("/t/view", textView)
	mux.HandleFunc("/t/create", textCreate)

	// value returned from flag.String() is a pointer to the flag value
	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
