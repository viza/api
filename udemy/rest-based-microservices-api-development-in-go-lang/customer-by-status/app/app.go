package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/viza/banking/domain"
	"github.com/viza/banking/service"
)

func Start() {

	router := mux.NewRouter()

	//wiring
	//ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDB())}

	// routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	
	//starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
