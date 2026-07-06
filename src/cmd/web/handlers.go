package main

import (
    "fmt"
    "net/http"
    "strconv"
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

    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil || id <1 {
        http.NotFound(w, r)
        return
    }

    fmt.Fprintf(w, "Display a sepcific snippet with ID %d...", id)

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