package pkg

import (
	"errors"
	"github.com/miekg/dns"
)

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
	var (
		m   dns.Msg
		ips []string
	)

	m.SetQuestion(dns.Fqdn(a.fqdn), dns.TypeA)

	in, err := dns.Exchange(&m, a.serverAddr)
	if err != nil {
		return ips, err
	}

	if len(in.Answer) < 1 {
		return ips, errors.New("no answer")
	}

	for _, answer := range in.Answer {
		if a, ok := answer.(*dns.A); ok {
			ips = append(ips, a.A.String())
		}
	}

	return ips, nil
}

func (a *application) TypeCNAME() ([]string, error) {
	return nil, nil
}

func (a *application) Navigate() []Result {
	return nil
}
