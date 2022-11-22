package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Customer struct {
	Name 			string 	`json:"name"`
	City 			string	`json:"city"`
	ZipCode 	string	`json:"zipcode"`
}

func greet(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hello world!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request){
	customers := []Customer {
		{Name: "Tiago Vaz", City: "São Paulo", ZipCode: "0123"},
		{Name: "João Silva", City: "Rio de Janeiro", ZipCode: "3210"},
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}