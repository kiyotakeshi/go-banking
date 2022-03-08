package app

import (
	"banking/domain"
	"banking/service"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"os"
	"time"
)

func Start() {
	sanityCheck()

	router := mux.NewRouter()

	dbClient := getDbClient()

	// customerHandlers := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	customerHandlers := CustomerHandlers{service.NewCustomerService(customerRepositoryDb)}

	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)
	accountHandler := AccountHandler{service.NewAccountService(accountRepositoryDb)}

	// define routes
	router.HandleFunc("/customers", customerHandlers.getCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", customerHandlers.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", accountHandler.NewAccount).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", accountHandler.MakeTransaction).Methods(http.MethodPost)

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" {
		log.Fatal("environment variable not defined...")
	}
}

func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbAddress := os.Getenv("DB_ADDRESS")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbAddress, dbPort, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}
