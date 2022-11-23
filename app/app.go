package app

import (
	"fmt"
	"github.com/vaz-tiago/go-banking/domain"
	"github.com/vaz-tiago/go-banking/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	// wiring
	//ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDB())}

	// define routes
	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)

	// starting server
	fmt.Println("🟢 Listen port 8000")
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
