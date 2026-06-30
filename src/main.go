package main

import (
    "net/http"
    "log"
)

// http://localhost:4000/
// Home handler function writes a byte slice containing 
// "Hello from Snippetbox" as the response body.
func home (w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello from Snippetbox"))
}

// http://localhost:4000//snippet/view
// Snippet viewing handler function writes a byte slice containing 
// "View your snippets" as the response body.
func view (w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("View your snippets"))
}

// http://localhost:4000//snippet/create
// Snippet creation handler function writes a byte slice containing 
// "Create your snippets" as the response body.
func create (w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Create a snippet"))
}

func main() {
    // NewServeMux function initialises a new http request, then registers 
    // the home function as the handler for the "/" URL pattern.
    mux := http.NewServeMux()
    mux.HandleFunc ("/", home)
    mux.HandleFunc ("/snippet/view", view)
    mux.HandleFunc ("/snippet/create", create)

    // http.ListenAndServe() function starts a new web server. Two parameters are passed
    // in: the TCP network address to listen on (":4000") and the servemux. If
    // http.ListenAndServe() returns an error the log.Fatal() function will log
    // the error message and exit.
    log.Println("Starting server on :4000")
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}