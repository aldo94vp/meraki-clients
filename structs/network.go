package structs

// Network structure
type Network struct {
	ID             string `json:"id"`
	OrganizationID string `json:"organizationId"`
	Name           string `json:"name"`
	TimeZone       string `json:"timeZone"`
	Type           string `json:"type"`
}

// Networks array
type Networks []Network
