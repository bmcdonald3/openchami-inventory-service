package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bmcdonald3/openchami-inventory-service/internal/datastore"
	"github.com/bmcdonald3/openchami-inventory-service/pkg/models"
	"github.com/go-chi/chi/v5"
)

func setupTestServer() *chi.Mux {
	db := datastore.NewMemoryStore()
	server := NewServer(db)
	router := NewRouter(server)
	return router
}

func TestDeviceLifecycle(t *testing.T) {
	router := setupTestServer()
	var createdDeviceID string

	t.Run("CreateDevice", func(t *testing.T) {
		devicePayload := `{"name":"Lifecycle Test Device","componentType":"Node","manufacturer":"HPE","partNumber":"P456","serialNumber":"SN456","status":"active"}`
		req := httptest.NewRequest("POST", "/inventory/v1/devices", bytes.NewBufferString(devicePayload))
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusCreated {
			t.Fatalf("CreateDevice failed: got status %v want %v", status, http.StatusCreated)
		}
		var device models.Device
		json.NewDecoder(rr.Body).Decode(&device)
		createdDeviceID = device.ID
	})

	t.Run("GetDeviceByID", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/inventory/v1/devices/"+createdDeviceID, nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusOK {
			t.Fatalf("GetDeviceByID failed: got status %v want %v", status, http.StatusOK)
		}
	})

	t.Run("DeleteDevice", func(t *testing.T) {
		req := httptest.NewRequest("DELETE", "/inventory/v1/devices/"+createdDeviceID, nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusNoContent {
			t.Fatalf("DeleteDevice failed: got status %v want %v", status, http.StatusNoContent)
		}
	})

	t.Run("GetAfterDelete", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/inventory/v1/devices/"+createdDeviceID, nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusNotFound {
			t.Fatalf("GetAfterDelete failed: got status %v want %v", status, http.StatusNotFound)
		}
	})
}

func TestInstallAndRemoveDevice(t *testing.T) {
	router := setupTestServer()
	var testDevice models.Device
	var testLocation models.Location

	// Setup: Create a device and a location
	devicePayload := `{"name":"Install Test Device","componentType":"Node","manufacturer":"Test","partNumber":"T1","serialNumber":"SN-T1","status":"active"}`
	req := httptest.NewRequest("POST", "/inventory/v1/devices", bytes.NewBufferString(devicePayload))
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	json.NewDecoder(rr.Body).Decode(&testDevice)

	locationPayload := `{"id":"test-slot-1","name":"Test Slot","locationType":"node_slot","status":"empty"}`
	req = httptest.NewRequest("POST", "/inventory/v1/locations", bytes.NewBufferString(locationPayload))
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	json.NewDecoder(rr.Body).Decode(&testLocation)

	// Step 1: Install the device
	t.Run("InstallDevice", func(t *testing.T) {
		installPayload := `{"deviceId":"` + testDevice.ID + `"}`
		req := httptest.NewRequest("PUT", "/inventory/v1/locations/"+testLocation.ID+"/device", bytes.NewBufferString(installPayload))
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Fatalf("InstallDevice failed: got status %v want %v", status, http.StatusOK)
		}

		var response struct{ Location models.Location }
		json.NewDecoder(rr.Body).Decode(&response)
		if response.Location.CurrentDeviceID == nil || *response.Location.CurrentDeviceID != testDevice.ID {
			t.Errorf("InstallDevice did not update the location's CurrentDeviceID")
		}
	})

	// Step 2: Verify the device's location is updated
	t.Run("VerifyDeviceLocation", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/inventory/v1/devices/"+testDevice.ID, nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		var device models.Device
		json.NewDecoder(rr.Body).Decode(&device)
		if device.CurrentLocationID == nil || *device.CurrentLocationID != testLocation.ID {
			t.Errorf("Device currentLocationId was not updated after install")
		}
	})

	// Step 3: Remove the device
	t.Run("RemoveDevice", func(t *testing.T) {
		req := httptest.NewRequest("DELETE", "/inventory/v1/locations/"+testLocation.ID+"/device", nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Fatalf("RemoveDevice failed: got status %v want %v", status, http.StatusOK)
		}

		var response struct{ Location models.Location }
		json.NewDecoder(rr.Body).Decode(&response)
		if response.Location.CurrentDeviceID != nil {
			t.Errorf("RemoveDevice did not clear the location's CurrentDeviceID")
		}
	})

	// Step 4: Verify the location's history
	t.Run("VerifyHistory", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/inventory/v1/locations/"+testLocation.ID+"/history", nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		var response struct{ Items []models.Event }
		json.NewDecoder(rr.Body).Decode(&response)
		if len(response.Items) != 2 {
			t.Fatalf("Expected 2 events in history, got %d", len(response.Items))
		}
		if !strings.Contains(response.Items[0].Type, "installed") {
			t.Errorf("First event was not 'installed'")
		}
		if !strings.Contains(response.Items[1].Type, "removed") {
			t.Errorf("Second event was not 'removed'")
		}
	})
}
