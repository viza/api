package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"town"`
	Zipcode string `json:"zip_code" xml:"zip"`
}

func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Kharkivman", City: "Kharkiv", Zipcode: "61029"},
		{Name: "Kyivman", City: "Kyiv", Zipcode: "32000"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)

	}

}

// func GetCustomer(w http.ResponseWriter, r *http.Request) {

// 	customers := []Customer{
// 		{Name: "Kharkivman", City: "Kharkiv", Zipcode: "61029"},
// 		{Name: "Kyivman", City: "Kyiv", Zipcode: "32000"},
// 	}

// 	if r.Header.Get("Content-Type") == "application/xml" {
// 		w.Header().Add("Content-Type", "application/xml")
// 		xml.NewEncoder(w).Encode(customers)
// 	} else {
// 		w.Header().Add("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(customers)

// 	}

// }
