package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bmcdonald3/openchami-inventory-service/internal/datastore"
	"github.com/bmcdonald3/openchami-inventory-service/pkg/models"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func setupTestServer() *chi.Mux {
	db := datastore.NewMemoryStore()
	server := NewServer(db)

	router := chi.NewRouter()
	router.Post("/inventory/v1/devices", server.createDeviceHandler)
	router.Get("/inventory/v1/devices", server.listDevicesHandler)
	router.Get("/inventory/v1/devices/{id}", server.getDeviceByIDHandler)
	router.Delete("/inventory/v1/devices/{id}", server.deleteDeviceHandler)

	return router
}
func TestCreateDeviceHandler(t *testing.T) {
	router := setupTestServer()

	devicePayload := `{"name":"Test Node 1","componentType":"Node","manufacturer":"HPE","partNumber":"P123","serialNumber":"SN123","status":"active"}`
	reqBody := bytes.NewBuffer([]byte(devicePayload))

	req := httptest.NewRequest("POST", "/inventory/v1/devices", reqBody)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	var createdDevice models.Device
	if err := json.NewDecoder(rr.Body).Decode(&createdDevice); err != nil {
		t.Fatalf("could not decode response: %v", err)
	}

	if createdDevice.Name != "Test Node 1" {
		t.Errorf("handler returned unexpected name: got %v want %v", createdDevice.Name, "Test Node 1")
	}
	if createdDevice.ID == "" {
		t.Errorf("handler returned device with no ID")
	}
}

func TestDeviceLifecycle(t *testing.T) {
	router := setupTestServer()
	var createdDeviceID string

	// Step 1: Create a device
	t.Run("CreateDevice", func(t *testing.T) {
		devicePayload := `{"name":"Lifecycle Test Device","componentType":"Node","manufacturer":"HPE","partNumber":"P456","serialNumber":"SN456","status":"active"}`
		reqBody := bytes.NewBuffer([]byte(devicePayload))
		req := httptest.NewRequest("POST", "/inventory/v1/devices", reqBody)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusCreated {
			t.Fatalf("CreateDevice failed: got status %v want %v", status, http.StatusCreated)
		}
		var device models.Device
		json.NewDecoder(rr.Body).Decode(&device)
		createdDeviceID = device.ID // Save the ID for the next steps
	})

	// Step 2: Get the device by ID to verify creation
	t.Run("GetDeviceByID", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/inventory/v1/devices/"+createdDeviceID, nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Fatalf("GetDeviceByID failed: got status %v want %v", status, http.StatusOK)
		}
		var device models.Device
		json.NewDecoder(rr.Body).Decode(&device)
		if device.Name != "Lifecycle Test Device" {
			t.Errorf("GetDeviceByID returned wrong name: got %v", device.Name)
		}
	})

	// Step 3: List devices to ensure it appears in the collection
	t.Run("ListDevices", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/inventory/v1/devices", nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Fatalf("ListDevices failed: got status %v want %v", status, http.StatusOK)
		}
		var response struct{ Items []models.Device }
		json.NewDecoder(rr.Body).Decode(&response)
		if len(response.Items) != 1 || response.Items[0].ID != createdDeviceID {
			t.Errorf("ListDevices did not contain the created device")
		}
	})

	// Step 4: Delete the device
	t.Run("DeleteDevice", func(t *testing.T) {
		req := httptest.NewRequest("DELETE", "/inventory/v1/devices/"+createdDeviceID, nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusNoContent {
			t.Fatalf("DeleteDevice failed: got status %v want %v", status, http.StatusNoContent)
		}
	})

	// Step 5: Get the device again to ensure it's gone
	t.Run("GetAfterDelete", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/inventory/v1/devices/"+createdDeviceID, nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusNotFound {
			t.Fatalf("GetAfterDelete failed: got status %v want %v", status, http.StatusNotFound)
		}
	})
}

func TestGetDeviceNotFound(t *testing.T) {
	router := setupTestServer()
	nonExistentID := uuid.NewString()

	req := httptest.NewRequest("GET", "/inventory/v1/devices/"+nonExistentID, nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code for non-existent device: got %v want %v", status, http.StatusNotFound)
	}
}
