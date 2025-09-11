package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bmcdonald3/openchami-inventory-service/internal/datastore"
	"github.com/bmcdonald3/openchami-inventory-service/pkg/models"
)

func TestCreateDeviceHandler(t *testing.T) {
	// 1. Setup
	db := datastore.NewMemoryStore()
	server := NewServer(db)
	router := NewRouter(server)

	// 2. Create a test payload
	devicePayload := `{"name":"Test Node 1","componentType":"Node","manufacturer":"HPE","partNumber":"P123","serialNumber":"SN123","status":"active"}`
	reqBody := bytes.NewBuffer([]byte(devicePayload))

	// 3. Create a mock request and a response recorder
	req := httptest.NewRequest("POST", "/inventory/v1/devices", reqBody)
	rr := httptest.NewRecorder()

	// 4. Execute the request
	router.ServeHTTP(rr, req)

	// 5. Check the results
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// 6. Decode the response and check the content
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
