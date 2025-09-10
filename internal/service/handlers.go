package service

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/bmcdonald3/openchami-inventory-service/pkg/models"
)

// --- Reusable Mock Data ---

var mockDevice1 = models.Device{
	ID:                "c3d4e5f6-a1b2-4c1d-8e9f-0c1d2e3f4a5b",
	Name:              "Compute Node 01",
	Hostname:          strPtr("c0-0c0s1n0.local"),
	ComponentType:     "Node",
	Manufacturer:      "HPE",
	PartNumber:        "NODE-EX235A",
	SerialNumber:      "NODESN001",
	CurrentLocationID: strPtr("x3000c7s1b0n0"),
	Status:            "active",
	Properties: map[string]interface{}{
		"nid": 1001,
	},
	CreatedAt: time.Now().Add(-24 * time.Hour),
	UpdatedAt: timePtr(time.Now()),
}

var mockLocation1 = Location{
	ID:               "x3000c7s1b0n0",
	Name:             "Node Slot 1 in Chassis 7",
	LocationType:     "node_slot",
	ParentLocationID: strPtr("x3000c7"),
	CurrentDeviceID:  strPtr("c3d4e5f6-a1b2-4c1d-8e9f-0c1d2e3f4a5b"),
	Status:           "occupied",
	CreatedAt:        time.Now().Add(-48 * time.Hour),
}

var mockEvent1 = Event{
	ID:          "e1a2b3c4-d5e6-4f1a-8b9c-0a1b2c3d4e5f",
	Source:      "/inventory/v1/api",
	SpecVersion: "1.0",
	Type:        "com.openchami.inventory.device.installed",
	Subject:     strPtr(mockDevice1.ID),
	Time:        time.Now(),
	Data: EventData{
		DeviceID:   strPtr(mockDevice1.ID),
		LocationID: strPtr(mockLocation1.ID),
		Actor:      strPtr("system"),
		Comment:    strPtr("Device installed during initial discovery."),
	},
}

// --- Handler Functions ---

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

func getDevicesHandler(w http.ResponseWriter, r *http.Request) {
	response := struct {
		Items      []Device       `json:"items"`
		Pagination PaginationInfo `json:"pagination"`
	}{
		Items: []Device{mockDevice1},
		Pagination: PaginationInfo{
			Count:  1,
			Total:  1,
			Offset: 0,
		},
	}
	writeJSON(w, http.StatusOK, response)
}

func createDeviceHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusCreated, mockDevice1)
}

func getDeviceByIdHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, mockDevice1)
}

func getDeviceByNameHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, mockDevice1)
}

func updateDeviceHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, mockDevice1)
}

func deleteDeviceHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func getDeviceHistoryHandler(w http.ResponseWriter, r *http.Request) {
	response := struct {
		Items      []Event        `json:"items"`
		Pagination PaginationInfo `json:"pagination"`
	}{
		Items: []Event{mockEvent1},
		Pagination: PaginationInfo{
			Count:  1,
			Total:  1,
			Offset: 0,
		},
	}
	writeJSON(w, http.StatusOK, response)
}

// --- Location Handlers ---

func getLocationsHandler(w http.ResponseWriter, r *http.Request) {
	response := struct {
		Items      []Location     `json:"items"`
		Pagination PaginationInfo `json:"pagination"`
	}{
		Items: []Location{mockLocation1},
		Pagination: PaginationInfo{
			Count:  1,
			Total:  1,
			Offset: 0,
		},
	}
	writeJSON(w, http.StatusOK, response)
}

func createLocationHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusCreated, mockLocation1)
}

func getLocationByIdHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, mockLocation1)
}

func getLocationByNameHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, mockLocation1)
}

func updateLocationHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, mockLocation1)
}

func deleteLocationHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func getLocationHistoryHandler(w http.ResponseWriter, r *http.Request) {
	getDeviceHistoryHandler(w, r)
}

func getDeviceAtLocationHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, mockDevice1)
}

func installDeviceHandler(w http.ResponseWriter, r *http.Request) {
	response := struct {
		Location Location `json:"location"`
		Event    Event    `json:"event"`
	}{
		Location: mockLocation1,
		Event:    mockEvent1,
	}
	writeJSON(w, http.StatusOK, response)
}

func removeDeviceHandler(w http.ResponseWriter, r *http.Request) {
	installDeviceHandler(w, r)
}

// --- Event Handlers ---

func getEventsHandler(w http.ResponseWriter, r *http.Request) {
	getDeviceHistoryHandler(w, r)
}

func getEventByIdHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, mockEvent1)
}
