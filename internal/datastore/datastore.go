package datastore

import "github.com/bmcdonald3/openchami-inventory-service/pkg/models"

type Datastore interface {
	// --- Device Methods ---
	CreateDevice(device *models.Device) error
	GetDeviceByID(id string) (*models.Device, error)
	GetDeviceByName(name string) (*models.Device, error)
	ListDevices() ([]*models.Device, error)
	UpdateDevice(device *models.Device) error
	DeleteDevice(id string) error

	// --- Location Methods ---
	CreateLocation(location *models.Location) error
	GetLocationByID(id string) (*models.Location, error)
	GetLocationByName(name string) (*models.Location, error)
	ListLocations() ([]*models.Location, error)
	UpdateLocation(location *models.Location) error
	DeleteLocation(id string) error

	// --- Event Methods ---
	CreateEvent(event *models.Event) error
	GetEventByID(id string) (*models.Event, error)
	ListEvents() ([]*models.Event, error)
	ListEventsByDeviceID(deviceID string) ([]*models.Event, error)
	ListEventsByLocationID(locationID string) ([]*models.Event, error)
}
