package app

import (
	"fmt"
	"log"
	"net/http"
)

func Start() {
	mux := http.NewServeMux()
		// define routes
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getAllCustomers)

	// starting server
	fmt.Println("🟢 Listen port 8000")
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}