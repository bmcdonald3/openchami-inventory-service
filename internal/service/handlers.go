package service

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/bmcdonald3/openchami-inventory-service/pkg/models"
	"github.com/go-chi/chi/v5"
)

// --- Helper Functions ---

func writeJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func strPtr(s string) *string        { return &s }
func timePtr(t time.Time) *time.Time { return &t }

// --- Device Handlers ---

func (s *Server) listDevicesHandler(w http.ResponseWriter, r *http.Request) {
	devices, err := s.DB.ListDevices()
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, models.ErrorResponse{Code: "internal_error", Message: err.Error()})
		return
	}
	response := struct {
		Items      []models.Device       `json:"items"`
		Pagination models.PaginationInfo `json:"pagination"`
	}{
		Items:      devices,
		Pagination: models.PaginationInfo{Count: len(devices), Total: len(devices), Offset: 0},
	}
	writeJSON(w, http.StatusOK, response)
}

func (s *Server) createDeviceHandler(w http.ResponseWriter, r *http.Request) {
	var device models.Device
	if err := json.NewDecoder(r.Body).Decode(&device); err != nil {
		writeJSON(w, http.StatusBadRequest, models.ErrorResponse{Code: "bad_request", Message: "Invalid JSON format"})
		return
	}
	createdDevice, err := s.DB.CreateDevice(&device)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, models.ErrorResponse{Code: "internal_error", Message: err.Error()})
		return
	}
	writeJSON(w, http.StatusCreated, createdDevice)
}

func (s *Server) getDeviceByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	device, err := s.DB.GetDeviceByID(id)
	if err != nil {
		writeJSON(w, http.StatusNotFound, models.ErrorResponse{Code: "not_found", Message: err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, device)
}

func (s *Server) getDeviceByNameHandler(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	device, err := s.DB.GetDeviceByName(name)
	if err != nil {
		writeJSON(w, http.StatusNotFound, models.ErrorResponse{Code: "not_found", Message: err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, device)
}

func (s *Server) updateDeviceHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var device models.Device
	if err := json.NewDecoder(r.Body).Decode(&device); err != nil {
		writeJSON(w, http.StatusBadRequest, models.ErrorResponse{Code: "bad_request", Message: "Invalid JSON format"})
		return
	}
	updatedDevice, err := s.DB.UpdateDevice(id, &device)
	if err != nil {
		writeJSON(w, http.StatusNotFound, models.ErrorResponse{Code: "not_found", Message: err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, updatedDevice)
}

func (s *Server) deleteDeviceHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := s.DB.DeleteDevice(id); err != nil {
		writeJSON(w, http.StatusNotFound, models.ErrorResponse{Code: "not_found", Message: err.Error()})
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// --- Location Handlers ---

func (s *Server) listLocationsHandler(w http.ResponseWriter, r *http.Request) {
	locations, err := s.DB.ListLocations()
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, models.ErrorResponse{Code: "internal_error", Message: err.Error()})
		return
	}
	response := struct {
		Items      []models.Location     `json:"items"`
		Pagination models.PaginationInfo `json:"pagination"`
	}{
		Items:      locations,
		Pagination: models.PaginationInfo{Count: len(locations), Total: len(locations), Offset: 0},
	}
	writeJSON(w, http.StatusOK, response)
}

func (s *Server) createLocationHandler(w http.ResponseWriter, r *http.Request) {
	var location models.Location
	if err := json.NewDecoder(r.Body).Decode(&location); err != nil {
		writeJSON(w, http.StatusBadRequest, models.ErrorResponse{Code: "bad_request", Message: "Invalid JSON format"})
		return
	}
	createdLocation, err := s.DB.CreateLocation(&location)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, models.ErrorResponse{Code: "internal_error", Message: err.Error()})
		return
	}
	writeJSON(w, http.StatusCreated, createdLocation)
}

func (s *Server) getLocationByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	location, err := s.DB.GetLocationByID(id)
	if err != nil {
		writeJSON(w, http.StatusNotFound, models.ErrorResponse{Code: "not_found", Message: err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, location)
}

func (s *Server) getLocationByNameHandler(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	location, err := s.DB.GetLocationByName(name)
	if err != nil {
		writeJSON(w, http.StatusNotFound, models.ErrorResponse{Code: "not_found", Message: err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, location)
}

func (s *Server) updateLocationHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var location models.Location
	if err := json.NewDecoder(r.Body).Decode(&location); err != nil {
		writeJSON(w, http.StatusBadRequest, models.ErrorResponse{Code: "bad_request", Message: "Invalid JSON format"})
		return
	}
	updatedLocation, err := s.DB.UpdateLocation(id, &location)
	if err != nil {
		writeJSON(w, http.StatusNotFound, models.ErrorResponse{Code: "not_found", Message: err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, updatedLocation)
}

func (s *Server) deleteLocationHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := s.DB.DeleteLocation(id); err != nil {
		writeJSON(w, http.StatusNotFound, models.ErrorResponse{Code: "not_found", Message: err.Error()})
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// --- Event and History Handlers ---

func (s *Server) listEventsHandler(w http.ResponseWriter, r *http.Request) {
	events, err := s.DB.ListEvents()
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, models.ErrorResponse{Code: "internal_error", Message: err.Error()})
		return
	}
	response := struct {
		Items      []models.Event        `json:"items"`
		Pagination models.PaginationInfo `json:"pagination"`
	}{
		Items:      events,
		Pagination: models.PaginationInfo{Count: len(events), Total: len(events), Offset: 0},
	}
	writeJSON(w, http.StatusOK, response)
}

func (s *Server) getEventByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	event, err := s.DB.GetEventByID(id)
	if err != nil {
		writeJSON(w, http.StatusNotFound, models.ErrorResponse{Code: "not_found", Message: err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, event)
}

func (s *Server) getDeviceHistoryHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	events, err := s.DB.ListEventsByDeviceID(id)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, models.ErrorResponse{Code: "internal_error", Message: err.Error()})
		return
	}
	response := struct {
		Items      []models.Event        `json:"items"`
		Pagination models.PaginationInfo `json:"pagination"`
	}{
		Items:      events,
		Pagination: models.PaginationInfo{Count: len(events), Total: len(events), Offset: 0},
	}
	writeJSON(w, http.StatusOK, response)
}

func (s *Server) getLocationHistoryHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	events, err := s.DB.ListEventsByLocationID(id)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, models.ErrorResponse{Code: "internal_error", Message: err.Error()})
		return
	}
	response := struct {
		Items      []models.Event        `json:"items"`
		Pagination models.PaginationInfo `json:"pagination"`
	}{
		Items:      events,
		Pagination: models.PaginationInfo{Count: len(events), Total: len(events), Offset: 0},
	}
	writeJSON(w, http.StatusOK, response)
}

// --- Composite Handlers (Install/Remove) ---

func (s *Server) getDeviceAtLocationHandler(w http.ResponseWriter, r *http.Request) {
	locationId := chi.URLParam(r, "id")
	location, err := s.DB.GetLocationByID(locationId)
	if err != nil {
		writeJSON(w, http.StatusNotFound, models.ErrorResponse{Code: "not_found", Message: err.Error()})
		return
	}
	if location.CurrentDeviceID == nil {
		writeJSON(w, http.StatusNotFound, models.ErrorResponse{Code: "not_found", Message: "No device at this location"})
		return
	}
	device, err := s.DB.GetDeviceByID(*location.CurrentDeviceID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, models.ErrorResponse{Code: "internal_error", Message: "Data inconsistency: device for this location not found"})
		return
	}
	writeJSON(w, http.StatusOK, device)
}

func (s *Server) installDeviceHandler(w http.ResponseWriter, r *http.Request) {
	locationId := chi.URLParam(r, "id")
	var body struct {
		DeviceID string `json:"deviceId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeJSON(w, http.StatusBadRequest, models.ErrorResponse{Code: "bad_request", Message: "Invalid JSON format"})
		return
	}
	location, err := s.DB.GetLocationByID(locationId)
	if err != nil {
		writeJSON(w, http.StatusNotFound, models.ErrorResponse{Code: "not_found", Message: "Location not found"})
		return
	}
	if location.CurrentDeviceID != nil {
		writeJSON(w, http.StatusBadRequest, models.ErrorResponse{Code: "bad_request", Message: "Location is already occupied"})
		return
	}
	device, err := s.DB.GetDeviceByID(body.DeviceID)
	if err != nil {
		writeJSON(w, http.StatusNotFound, models.ErrorResponse{Code: "not_found", Message: "Device not found"})
		return
	}
	location.CurrentDeviceID = &device.ID
	location.Status = "occupied"
	device.CurrentLocationID = &location.ID
	s.DB.UpdateLocation(location.ID, location)
	s.DB.UpdateDevice(device.ID, device)
	installEvent := &models.Event{
		Source:      "/inventory/v1/api",
		SpecVersion: "1.0",
		Type:        "com.openchami.inventory.device.installed",
		Data: models.EventData{
			DeviceID:   &device.ID,
			LocationID: &location.ID,
			Actor:      strPtr("api-user"),
		},
	}
	createdEvent, _ := s.DB.CreateEvent(installEvent)
	response := struct {
		Location models.Location `json:"location"`
		Event    models.Event    `json:"event"`
	}{
		Location: *location,
		Event:    *createdEvent,
	}
	writeJSON(w, http.StatusOK, response)
}

func (s *Server) removeDeviceHandler(w http.ResponseWriter, r *http.Request) {
	locationId := chi.URLParam(r, "id")
	location, err := s.DB.GetLocationByID(locationId)
	if err != nil {
		writeJSON(w, http.StatusNotFound, models.ErrorResponse{Code: "not_found", Message: "Location not found"})
		return
	}
	if location.CurrentDeviceID == nil {
		writeJSON(w, http.StatusBadRequest, models.ErrorResponse{Code: "bad_request", Message: "Location is already empty"})
		return
	}
	device, err := s.DB.GetDeviceByID(*location.CurrentDeviceID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, models.ErrorResponse{Code: "internal_error", Message: "Could not find device associated with this location"})
		return
	}
	location.CurrentDeviceID = nil
	location.Status = "empty"
	device.CurrentLocationID = nil
	s.DB.UpdateLocation(location.ID, location)
	s.DB.UpdateDevice(device.ID, device)
	removeEvent := &models.Event{
		Source:      "/inventory/v1/api",
		SpecVersion: "1.0",
		Type:        "com.openchami.inventory.device.removed",
		Data: models.EventData{
			DeviceID:   &device.ID,
			LocationID: &location.ID,
			Actor:      strPtr("api-user"),
		},
	}
	createdEvent, _ := s.DB.CreateEvent(removeEvent)
	response := struct {
		Location models.Location `json:"location"`
		Event    models.Event    `json:"event"`
	}{
		Location: *location,
		Event:    *createdEvent,
	}
	writeJSON(w, http.StatusOK, response)
}
