package pkg

// Empty type.
type Empty struct{}

// Result type of each lookup.
type Result struct {
	IPAddress string `json:"ip_address"`
	Hostname  string `json:"hostname"`
}
