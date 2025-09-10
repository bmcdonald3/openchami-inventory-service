package service

import "github.com/bmcdonald3/openchami-inventory-service/internal/datastore"

// Server is the main application struct that holds dependencies.
type Server struct {
	DB datastore.Datastore
}

// NewServer creates a new server with its dependencies.
func NewServer(db datastore.Datastore) *Server {
	return &Server{DB: db}
}
