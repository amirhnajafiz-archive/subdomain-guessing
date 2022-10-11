package pkg

import (
	"fmt"

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
		message dns.Msg
		ips     []string
	)

	message.SetQuestion(dns.Fqdn(a.fqdn), dns.TypeA)

	in, err := dns.Exchange(&message, a.serverAddr)
	if err != nil {
		return ips, fmt.Errorf("dns exchange error: %w", err)
	}

	if len(in.Answer) < 1 {
		return ips, fmt.Errorf("no answer for this address")
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
