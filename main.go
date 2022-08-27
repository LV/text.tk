package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// --- HANDLERS ---
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from text.tk"))
}

func textView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id")) // Extract the value of the id parameter from the query string
	if err != nil || id <= 0 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Dislpay a specific text with ID %d...", id) // interpolate id value with response and write to http.ResponseWriter
}

func textCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new text..."))
}

// --- END HANDLERS ---

func main() {
	mux := http.NewServeMux() // initialize new servemux
	mux.HandleFunc("/", home)
	mux.HandleFunc("/t/view", textView)
	mux.HandleFunc("/t/create", textCreate)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux) // start a new web server
	log.Fatal(err)
}
