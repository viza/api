package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	//mux := http.NewServeMux()
	router := mux.NewRouter()

	// routes
	router.HandleFunc("/greet", Greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", CreateCustomer).Methods(http.MethodPost)

	router.HandleFunc("/customers/{customer_id:[0-9]+}", GetCustomer).Methods(http.MethodGet)

	//starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))

}

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])
}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post request received")
}
