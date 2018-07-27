package structs

// Client structure
type Client struct {
	ID   string `json:"id"`
	MAC  string `json:"mac"`
	SSID string `json:"ssid,omitempty"`
}

// Clients array
type Clients []Client
