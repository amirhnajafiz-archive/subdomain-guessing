package pkg

type Lookup interface {
	TypeA(fqdn, serverAddr string) ([]string, error)
}
