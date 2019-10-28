package main

import (
	"fmt"
	"net/http"
)

var counter int

func main() {
	http.Handle("/", http.FileServer(http.Dir("client")))
	http.HandleFunc("/sse/dashboard", dashboard)
	http.ListenAndServe(":8080", nil)
}

func dashboard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	counter++
	fmt.Fprintf(w, "data: %v\n\n", counter)
	fmt.Printf("data: %v\n", counter)
}
