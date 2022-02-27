package app

import (
	"banking/domain"
	"banking/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	router := mux.NewRouter()

	customerHandlers := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	// define routes
	router.HandleFunc("/greeting", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", customerHandlers.greetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", greetCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
