package structs

// Organization structure
type Organization struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Organizations array
type Organizations []Organization
