package structs

// Device structure
type Device struct {
	Serial    string `json:"serial"`
	MAC       string `json:"mac"`
	Model     string `json:"model"`
	NetworkID string `json:"networkId"`
}

// Devices array
type Devices []Device
