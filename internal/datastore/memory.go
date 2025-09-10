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

	// Set server-managed fields
	device.ID = uuid.NewString()
	device.CreatedAt = time.Now()

	if _, exists := s.devices[device.ID]; exists {
		return nil, fmt.Errorf("device with ID %s already exists", device.ID)
	}
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

	// Update fields
	existingDevice.Name = device.Name
	existingDevice.Hostname = device.Hostname
	// ... update other fields as needed
	now := time.Now()
	existingDevice.UpdatedAt = &now

	s.devices[id] = existingDevice
	return existingDevice, nil
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

// --- Location Methods (Not Implemented) ---
func (s *MemoryStore) CreateLocation(location *models.Location) (*models.Location, error) {
	return nil, fmt.Errorf("not implemented")
}
func (s *MemoryStore) GetLocationByID(id string) (*models.Location, error) {
	return nil, fmt.Errorf("not implemented")
}
func (s *MemoryStore) GetLocationByName(name string) (*models.Location, error) {
	return nil, fmt.Errorf("not implemented")
}
func (s *MemoryStore) ListLocations() ([]models.Location, error) {
	return nil, fmt.Errorf("not implemented")
}
func (s *MemoryStore) UpdateLocation(id string, location *models.Location) (*models.Location, error) {
	return nil, fmt.Errorf("not implemented")
}
func (s *MemoryStore) DeleteLocation(id string) error { return fmt.Errorf("not implemented") }

// --- Event Methods (Not Implemented) ---
func (s *MemoryStore) CreateEvent(event *models.Event) (*models.Event, error) {
	return nil, fmt.Errorf("not implemented")
}
func (s *MemoryStore) GetEventByID(id string) (*models.Event, error) {
	return nil, fmt.Errorf("not implemented")
}
func (s *MemoryStore) ListEvents() ([]models.Event, error) { return nil, fmt.Errorf("not implemented") }
func (s *MemoryStore) ListEventsByDeviceID(deviceID string) ([]models.Event, error) {
	return nil, fmt.Errorf("not implemented")
}
func (s *MemoryStore) ListEventsByLocationID(locationID string) ([]models.Event, error) {
	return nil, fmt.Errorf("not implemented")
}
