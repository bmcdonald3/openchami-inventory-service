package service

import (
	"encoding/json"
	"net/http"

	"github.com/bmcdonald3/openchami-inventory-service/pkg/models"
	"github.com/go-chi/chi/v5"
)

// --- Handlers ---

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
		Items: devices,
		Pagination: models.PaginationInfo{
			Count:  len(devices),
			Total:  len(devices), // In-memory doesn't support total count easily
			Offset: 0,
		},
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
		writeJSON(w, http.StatusInternalServerError, models.ErrorResponse{Code: "internal_error", Message: err.Error()})
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

// --- Mock/Placeholder Handlers for other resources ---

func (s *Server) getDeviceHistoryHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusNotImplemented, "Not Implemented")
}
func (s *Server) listLocationsHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusNotImplemented, "Not Implemented")
}
func (s *Server) createLocationHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusNotImplemented, "Not Implemented")
}
func (s *Server) getLocationByIDHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusNotImplemented, "Not Implemented")
}
func (s *Server) getLocationByNameHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusNotImplemented, "Not Implemented")
}
func (s *Server) updateLocationHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusNotImplemented, "Not Implemented")
}
func (s *Server) deleteLocationHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusNotImplemented, "Not Implemented")
}
func (s *Server) getLocationHistoryHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusNotImplemented, "Not Implemented")
}
func (s *Server) getDeviceAtLocationHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusNotImplemented, "Not Implemented")
}
func (s *Server) installDeviceHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusNotImplemented, "Not Implemented")
}
func (s *Server) removeDeviceHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusNotImplemented, "Not Implemented")
}
func (s *Server) listEventsHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusNotImplemented, "Not Implemented")
}
func (s *Server) getEventByIDHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusNotImplemented, "Not Implemented")
}

// --- Helper ---
func writeJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
