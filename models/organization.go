package models

type OrganizationStatus string

const (
	OrganizationStatusActive  OrganizationStatus = "active"
	OrganizationStatusPending OrganizationStatus = "pending"
	OrganizationStatusDeleted OrganizationStatus = "deleted"
)

type Organization struct {
	ID          string             `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description" `
	Status      OrganizationStatus `json:"status"`
}
