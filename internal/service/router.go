package service

import (
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

func NewRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	routes := generateRoutes()

	for _, route := range routes {
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
