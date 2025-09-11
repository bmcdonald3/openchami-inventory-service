package datastore

import (
	"fmt"
	"sync"
	"time"

	"github.com/bmcdonald3/openchami-inventory-service/pkg/models"
	"github.com/google/uuid"
)

// MemoryStore is an in-memory implementation of the Datastore interface.
type MemoryStore struct {
	mu        sync.RWMutex
	devices   map[string]*models.Device
	locations map[string]*models.Location
	events    map[string]*models.Event
}

// NewMemoryStore creates and returns a new MemoryStore.
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		devices:   make(map[string]*models.Device),
		locations: make(map[string]*models.Location),
		events:    make(map[string]*models.Event),
	}
}

// --- Device Methods ---

func (s *MemoryStore) CreateDevice(device *models.Device) (*models.Device, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	device.ID = uuid.NewString()
	device.CreatedAt = time.Now()
	s.devices[device.ID] = device
	return device, nil
}

func (s *MemoryStore) GetDeviceByID(id string) (*models.Device, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	device, exists := s.devices[id]
	if !exists {
		return nil, fmt.Errorf("device with ID %s not found", id)
	}
	return device, nil
}

func (s *MemoryStore) GetDeviceByName(name string) (*models.Device, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, device := range s.devices {
		if device.Name == name {
			return device, nil
		}
	}
	return nil, fmt.Errorf("device with name '%s' not found", name)
}

func (s *MemoryStore) ListDevices() ([]models.Device, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	allDevices := make([]models.Device, 0, len(s.devices))
	for _, device := range s.devices {
		allDevices = append(allDevices, *device)
	}
	return allDevices, nil
}

func (s *MemoryStore) UpdateDevice(id string, device *models.Device) (*models.Device, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	existingDevice, exists := s.devices[id]
	if !exists {
		return nil, fmt.Errorf("device with ID %s not found", id)
	}
	// Preserve original creation time and ID
	device.CreatedAt = existingDevice.CreatedAt
	device.ID = id
	now := time.Now()
	device.UpdatedAt = &now
	s.devices[id] = device
	return device, nil
}

func (s *MemoryStore) DeleteDevice(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.devices[id]; !exists {
		return fmt.Errorf("device with ID %s not found", id)
	}
	delete(s.devices, id)
	return nil
}

// --- Location Methods ---

func (s *MemoryStore) CreateLocation(location *models.Location) (*models.Location, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	location.CreatedAt = time.Now()
	if _, exists := s.locations[location.ID]; exists {
		return nil, fmt.Errorf("location with ID %s already exists", location.ID)
	}
	s.locations[location.ID] = location
	return location, nil
}

func (s *MemoryStore) GetLocationByID(id string) (*models.Location, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	location, exists := s.locations[id]
	if !exists {
		return nil, fmt.Errorf("location with ID %s not found", id)
	}
	return location, nil
}

func (s *MemoryStore) GetLocationByName(name string) (*models.Location, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, location := range s.locations {
		if location.Name == name {
			return location, nil
		}
	}
	return nil, fmt.Errorf("location with name '%s' not found", name)
}

func (s *MemoryStore) ListLocations() ([]models.Location, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	allLocations := make([]models.Location, 0, len(s.locations))
	for _, location := range s.locations {
		allLocations = append(allLocations, *location)
	}
	return allLocations, nil
}

func (s *MemoryStore) UpdateLocation(id string, location *models.Location) (*models.Location, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	existingLocation, exists := s.locations[id]
	if !exists {
		return nil, fmt.Errorf("location with ID %s not found", id)
	}
	// Preserve original creation time and ID
	location.CreatedAt = existingLocation.CreatedAt
	location.ID = id
	now := time.Now()
	location.UpdatedAt = &now
	s.locations[id] = location
	return location, nil
}

func (s *MemoryStore) DeleteLocation(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.locations[id]; !exists {
		return fmt.Errorf("location with ID %s not found", id)
	}
	delete(s.locations, id)
	return nil
}

// --- Event Methods ---

func (s *MemoryStore) CreateEvent(event *models.Event) (*models.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	event.ID = uuid.NewString()
	event.Time = time.Now()
	s.events[event.ID] = event
	return event, nil
}

func (s *MemoryStore) GetEventByID(id string) (*models.Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	event, exists := s.events[id]
	if !exists {
		return nil, fmt.Errorf("event with ID %s not found", id)
	}
	return event, nil
}

func (s *MemoryStore) ListEvents() ([]models.Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	allEvents := make([]models.Event, 0, len(s.events))
	for _, event := range s.events {
		allEvents = append(allEvents, *event)
	}
	return allEvents, nil
}

func (s *MemoryStore) ListEventsByDeviceID(deviceID string) ([]models.Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var deviceEvents []models.Event
	for _, event := range s.events {
		if event.Data.DeviceID != nil && *event.Data.DeviceID == deviceID {
			deviceEvents = append(deviceEvents, *event)
		}
	}
	return deviceEvents, nil
}

func (s *MemoryStore) ListEventsByLocationID(locationID string) ([]models.Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var locationEvents []models.Event
	for _, event := range s.events {
		if event.Data.LocationID != nil && *event.Data.LocationID == locationID {
			locationEvents = append(locationEvents, *event)
		}
	}
	return locationEvents, nil
}
