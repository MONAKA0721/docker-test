package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {
    port := os.Getenv("PORT")
    fmt.Printf("Starting server at Port %s", port)
    if port == "" {
        log.Fatal("$PORT must be set")
    }

    http.HandleFunc("/", handler)
    http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello worldだよぉお！(´・ω・`)b\n")
}
