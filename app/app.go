package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	router := mux.NewRouter()
	// define routes
	router.HandleFunc("/greeting", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", greetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", greetCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "post request received")
}

func greetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])
}
