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

// NewRouter now accepts the Server struct.
func NewRouter(s *Server) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	// Pass the server to generate the routes.
	routes := generateRoutes(s)

	for _, route := range routes {
		router.MethodFunc(route.Method, route.Pattern, route.HandlerFunc)
	}

	return router
}

func generateRoutes(s *Server) Routes {
	return Routes{
		// --- Device Routes ---
		{"ListDevices", "GET", "/inventory/v1/devices", s.listDevicesHandler},
		{"CreateDevice", "POST", "/inventory/v1/devices", s.createDeviceHandler},
		{"GetDeviceByID", "GET", "/inventory/v1/devices/{id}", s.getDeviceByIDHandler},
		{"GetDeviceByName", "GET", "/inventory/v1/devices/by-name/{name}", s.getDeviceByNameHandler},
		{"UpdateDevice", "PUT", "/inventory/v1/devices/{id}", s.updateDeviceHandler},
		{"DeleteDevice", "DELETE", "/inventory/v1/devices/{id}", s.deleteDeviceHandler},
		{"GetDeviceHistory", "GET", "/inventory/v1/devices/{id}/history", s.getDeviceHistoryHandler},

		// --- Location Routes ---
		{"ListLocations", "GET", "/inventory/v1/locations", s.listLocationsHandler},
		{"CreateLocation", "POST", "/inventory/v1/locations", s.createLocationHandler},
		{"GetLocationByID", "GET", "/inventory/v1/locations/{id}", s.getLocationByIDHandler},
		{"GetLocationByName", "GET", "/inventory/v1/locations/by-name/{name}", s.getLocationByNameHandler},
		{"UpdateLocation", "PUT", "/inventory/v1/locations/{id}", s.updateLocationHandler},
		{"DeleteLocation", "DELETE", "/inventory/v1/locations/{id}", s.deleteLocationHandler},
		{"GetLocationHistory", "GET", "/inventory/v1/locations/{id}/history", s.getLocationHistoryHandler},
		{"GetDeviceAtLocation", "GET", "/inventory/v1/locations/{id}/device", s.getDeviceAtLocationHandler},
		{"InstallDevice", "PUT", "/inventory/v1/locations/{id}/device", s.installDeviceHandler},
		{"RemoveDevice", "DELETE", "/inventory/v1/locations/{id}/device", s.removeDeviceHandler},

		// --- Event Routes ---
		{"ListEvents", "GET", "/inventory/v1/events", s.listEventsHandler},
		{"GetEventByID", "GET", "/inventory/v1/events/{id}", s.getEventByIDHandler},
	}
}
