// main.go
package main

import (
	"log"
	"net/http"
	"study-case/handler"
	"study-case/repository"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize the database connection
	repository.InitDB()

	// Initialize the router
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/transactions", handler.CreateTransactionHandler).Methods("POST")

	// Start the server
	log.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
