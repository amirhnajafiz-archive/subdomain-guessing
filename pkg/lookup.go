package pkg

// Lookup
// manages 3 methods for lookup types.
type Lookup interface {
	// TypeA ns lookup.
	TypeA(fqdn, serverAddr string) ([]string, error)
	// TypeCNAME ns lookup.
	TypeCNAME(fqdn, serverAddr string) ([]string, error)
	// Navigate ns lookup.
	Navigate(fqdn, serverAddr string) []Result
}

type application struct{}

func NewApp() Lookup {
	return &application{}
}

func (a *application) TypeA(fqdn, serverAddr string) ([]string, error) {
	return nil, nil
}

func (a *application) TypeCNAME(fqdn, serverAddr string) ([]string, error) {
	return nil, nil
}

func (a *application) Navigate(fqdn, serverAddr string) []Result {
	return nil
}
