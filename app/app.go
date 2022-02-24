package app

import (
	"log"
	"net/http"
)

func Start() {
	mux := http.NewServeMux()
	// define routes
	mux.HandleFunc("/greeting", greet)
	mux.HandleFunc("/customers", greetCustomers)

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
