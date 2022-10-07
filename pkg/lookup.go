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

type application struct {
	fqdn       string
	serverAddr string
}

func NewApp(fullyQualifiedDomainName string, serverAddr string) Lookup {
	return &application{
		fqdn:       fullyQualifiedDomainName,
		serverAddr: serverAddr,
	}
}

func (a *application) TypeA() ([]string, error) {
	return nil, nil
}

func (a *application) TypeCNAME() ([]string, error) {
	return nil, nil
}

func (a *application) Navigate() []Result {
	return nil
}
