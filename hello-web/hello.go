package main

import (
	"fmt"
	"log"
	"net/http"
)

// Set function handler and listen on port 80
func main() {
	http.HandleFunc("/", HelloServer)
	log.Fatal(http.ListenAndServe(":80", nil))
}

// Write response
func HelloServer(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "Hello, world!")
	if err != nil {
		panic(err)
	}
}
