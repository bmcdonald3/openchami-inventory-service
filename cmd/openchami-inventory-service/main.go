package main

import (
	"log"
	"net/http"

	"github.com/bmcdonald3/openchami-inventory-service/internal/datastore"
	"github.com/bmcdonald3/openchami-inventory-service/internal/service"
)

func main() {
	// Create the in-memory datastore.
	db := datastore.NewMemoryStore()

	// Create the server, injecting the datastore.
	server := service.NewServer(db)

	// Create the router, passing the server to it.
	router := service.NewRouter(server)

	// Start the server.
	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", router)
	log.Fatal(err)
}
