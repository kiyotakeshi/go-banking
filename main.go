package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"name"`
	City    string `json:"city"`
	Zipcode string `json:"zip_code"`
}

func main() {

	// define routes
	http.HandleFunc("/greeting", greet)
	http.HandleFunc("/customers", greetCustomers)

	http.ListenAndServe("localhost:8000", nil)
}

func greet(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func greetCustomers(w http.ResponseWriter, request *http.Request) {
	customers := []Customer{
		{"mike", "tokyo", "103-0027"},
		{"popcorn", "tokyo", "104-0061"},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
