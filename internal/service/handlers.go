package service

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func getDevicesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Endpoint hit: GET /inventory/v1/devices")
}

func createDeviceHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Endpoint hit: POST /inventory/v1/devices")
}

func getDeviceByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	fmt.Fprintf(w, "Endpoint hit: GET /inventory/v1/devices/%s\n", id)
}
