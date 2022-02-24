package main

import (
	"fmt"
	"net/http"
)

func main() {
	// define routes
	http.HandleFunc("/greeting", greet)

	http.ListenAndServe("localhost:8000", nil)
}

func greet(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "hello world")
}
