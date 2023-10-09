package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kodekage/banking/domain"
	"github.com/kodekage/banking/service"
)

func StartServer() {
	router := mux.NewRouter()

	// wiring
	//customerHandler := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	customerHandler := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// define routes
	router.HandleFunc("/customers", customerHandler.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", customerHandler.getCustomer).Methods(http.MethodGet)

	// Start server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
