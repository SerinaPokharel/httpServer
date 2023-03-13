package main

import (
    "fmt"
    "net/http"
)

func main() {
    // This function handles requests to the root ("/") URL path.
    // It responds with the message "You are in server".
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "You are in server")
    })

    // This function starts an HTTP server on port 80 and uses the default
    // http.ServeMux to handle incoming requests. Once started, the server will
    // continue running and listening for incoming requests until it is explicitly
    // stopped or terminated.
    http.ListenAndServe(":80", nil)
}
