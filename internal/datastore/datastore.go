package datastore

import (
	"fmt"
	"sync"

	"github.com/bmcdonald3/openchami-inventory-service/pkg/models"
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

func (s *MemoryStore) CreateDevice(device *models.Device) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.devices[device.ID]; exists {
		return fmt.Errorf("device with ID %s already exists", device.ID)
	}
	s.devices[device.ID] = device
	return nil
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
	return nil, fmt.Errorf("device with name %s not found", name)
}

func (s *MemoryStore) ListDevices() ([]*models.Device, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var allDevices []*models.Device
	for _, device := range s.devices {
		allDevices = append(allDevices, device)
	}
	return allDevices, nil
}

func (s *MemoryStore) UpdateDevice(device *models.Device) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.devices[device.ID]; !exists {
		return fmt.Errorf("device with ID %s not found", device.ID)
	}
	s.devices[device.ID] = device
	return nil
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

// --- Location Methods (To be implemented later) ---

func (s *MemoryStore) CreateLocation(location *models.Location) error {
	return fmt.Errorf("not implemented")
}
func (s *MemoryStore) GetLocationByID(id string) (*models.Location, error) {
	return nil, fmt.Errorf("not implemented")
}

// ... and so on for other location methods

// --- Event Methods (To be implemented later) ---

func (s *MemoryStore) CreateEvent(event *models.Event) error {
	return fmt.Errorf("not implemented")
}

// ... and so on for other event methods
