package service

import "time"

// Device represents a physical piece of hardware in the inventory.
type Device struct {
	ID                string                 `json:"id"`
	Name              string                 `json:"name"`
	Hostname          *string                `json:"hostname,omitempty"`
	ComponentType     string                 `json:"componentType"`
	Manufacturer      string                 `json:"manufacturer"`
	PartNumber        string                 `json:"partNumber"`
	SerialNumber      string                 `json:"serialNumber"`
	CurrentLocationID *string                `json:"currentLocationId,omitempty"`
	Status            string                 `json:"status"`
	Properties        map[string]interface{} `json:"properties,omitempty"`
	ParentDeviceID    *string                `json:"parentDeviceId,omitempty"`
	ChildrenDeviceIDs []string               `json:"childrenDeviceIds,omitempty"`
	CreatedAt         time.Time              `json:"createdAt"`
	UpdatedAt         *time.Time             `json:"updatedAt,omitempty"`
	DeletedAt         *time.Time             `json:"deletedAt,omitempty"`
}

// Location represents a physical slot or bay where hardware can be installed.
type Location struct {
	ID                  string                 `json:"id"`
	Name                string                 `json:"name"`
	LocationType        string                 `json:"locationType"`
	ParentLocationID    *string                `json:"parentLocationId,omitempty"`
	ChildrenLocationIDs []string               `json:"childrenLocationIds,omitempty"`
	CurrentDeviceID     *string                `json:"currentDeviceId,omitempty"`
	Status              string                 `json:"status"`
	Properties          map[string]interface{} `json:"properties,omitempty"`
	CreatedAt           time.Time              `json:"createdAt"`
	UpdatedAt           *time.Time             `json:"updatedAt,omitempty"`
	DeletedAt           *time.Time             `json:"deletedAt,omitempty"`
}

// Event represents a historical record, conforming to the CloudEvents v1.0 spec.
type Event struct {
	ID              string    `json:"id"`
	Source          string    `json:"source"`
	SpecVersion     string    `json:"specversion"`
	Type            string    `json:"type"`
	DataContentType *string   `json:"datacontenttype,omitempty"`
	Subject         *string   `json:"subject,omitempty"`
	Time            time.Time `json:"time"`
	Data            EventData `json:"data"`
}

// EventData contains the inventory-specific payload of a CloudEvent.
type EventData struct {
	DeviceID    *string                `json:"deviceId,omitempty"`
	LocationID  *string                `json:"locationId,omitempty"`
	Actor       *string                `json:"actor,omitempty"`
	Comment     *string                `json:"comment,omitempty"`
	Duration    *int32                 `json:"duration,omitempty"`
	StateBefore map[string]interface{} `json:"stateBefore,omitempty"`
	StateAfter  map[string]interface{} `json:"stateAfter,omitempty"`
}
