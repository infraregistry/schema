package models

type ComponentStatus string

const (
	ComponentStatusActive   ComponentStatus = "active"
	ComponentStatusInactive ComponentStatus = "inactive"
)

type Component struct {
	ID          string          `json:"id,omitempty"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Status      ComponentStatus `json:"status"`
	Children    []Component     `json:"children,omitempty"`
}
