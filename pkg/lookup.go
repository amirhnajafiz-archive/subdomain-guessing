package pkg

// Lookup
// manages 3 methods for lookup types.
type Lookup interface {
	// TypeA ns lookup.
	TypeA() ([]string, error)
	// TypeCNAME ns lookup.
	TypeCNAME() ([]string, error)
	// Navigate ns lookup.
	Navigate() []Result
}
