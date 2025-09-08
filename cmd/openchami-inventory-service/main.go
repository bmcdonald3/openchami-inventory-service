package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func newRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	routes := generateRoutes()

	for _, route := range routes {
		// more complex stuff like auth, etc., as seen in SMD to come later...
		router.MethodFunc(route.Method, route.Pattern, route.HandlerFunc)
	}

	return router
}

func generateRoutes() Routes {
	return Routes{
		// --- Device Routes ---
		{
			"GetDevices",
			"GET",
			"/inventory/v1/devices",
			getDevicesHandler,
		},
		{
			"CreateDevice",
			"POST",
			"/inventory/v1/devices",
			createDeviceHandler,
		},
		{
			"GetDeviceByID",
			"GET",
			"/inventory/v1/devices/{id}",
			getDeviceByIdHandler,
		},
		// ... add all other device, location, and event routes here ...
	}
}

func main() {
	router := newRouter()

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", router)
	log.Fatal(err)
}

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
