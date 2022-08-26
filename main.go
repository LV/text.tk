package main

import (
	"log"
	"net/http"
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
	w.Write([]byte("Display a specific text..."))
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
