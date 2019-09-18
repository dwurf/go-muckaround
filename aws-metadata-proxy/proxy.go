package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
)

// Set function handler and listen on port 80
func main() {
	http.HandleFunc("/", InstanceIdentity)
	log.Fatal(http.ListenAndServe(":80", nil))
}

// Get instance metadata and write to client
func InstanceIdentity(w http.ResponseWriter, r *http.Request) {
	// Fetch instance metadata - restricted to just this endpoint for security reasons
	// https://ejj.io/blog/capital-one
	// Even this document leaks more info about your instance than you really should so
	// don't run this code for anything other than a quick test.
	resp, err := http.Get("http://169.254.169.254/latest/dynamic/instance-identity/document")
	if err != nil {
		_, err := fmt.Fprintf(w, "Error: %+v", err)
		if err != nil {
			panic(err)
		}
	}
	defer resp.Body.Close()

	// Read instance metadata
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		_, err := fmt.Fprintf(w, "Error: %+v", err)
		if err != nil {
			panic(err)
		}
	}

	// Return instance metadata
	_, err = fmt.Fprintf(w, "Hello! Here is some info about me\n%s\n", body)
	if err != nil {
		panic(err)
	}
}
