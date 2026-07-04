package main

import (
    "net/http"
    "log"
)

// home handler function
func home (w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    w.Write([]byte("Hello from Snippetbox"))
}

// snippet/view handler function
func view (w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("View your snippets"))
}

// snippet/create handler function
func create (w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.Header().Set("Allow", http.MethodPost)
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }

    w.Write([]byte("Create a snippet"))
}

// main function
func main() {
    mux := http.NewServeMux()
    mux.HandleFunc ("/", home)
    mux.HandleFunc ("/snippet/view", view)
    mux.HandleFunc ("/snippet/create", create)
    log.Println("Starting server on :4000")
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}