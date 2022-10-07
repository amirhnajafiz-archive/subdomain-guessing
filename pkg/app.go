package pkg

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
