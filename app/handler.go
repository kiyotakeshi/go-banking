package app

import (
	"banking/service"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Customer struct {
	Name    string `json:"name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zip_code"`
}

func greet(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "hello world")
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (customerHandlers *CustomerHandlers) greetCustomers(w http.ResponseWriter, r *http.Request) {
	//customers := []Customer{
	//	{"mike", "tokyo", "103-0027"},
	//	{"popcorn", "tokyo", "104-0061"},
	//}
	customers, _ := customerHandlers.service.GetAllCustomer()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "post request received")
}

func greetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])
}
