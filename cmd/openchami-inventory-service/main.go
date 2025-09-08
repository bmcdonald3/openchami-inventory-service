package main

import (
	"log"
	"net/http"

	"github.com/bmcdonald3/openchami-inventory-service/internal/service"
)

func main() {
	router := service.NewRouter()

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", router)
	log.Fatal(err)
}
