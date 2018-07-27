package structs

// SSID structure
type SSID struct {
	Name    string `json:"name"`
	Enabled bool   `json:"enabled"`
	Count   int
}

// SSIDs array
type SSIDs []SSID
