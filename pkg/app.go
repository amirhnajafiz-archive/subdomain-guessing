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
	var (
		message dns.Msg
		fqdns   []string
	)

	message.SetQuestion(dns.Fqdn(a.fqdn), dns.TypeCNAME)

	in, err := dns.Exchange(&message, a.serverAddr)
	if err != nil {
		return fqdns, fmt.Errorf("dns exchange error: %w", err)
	}

	if len(in.Answer) < 1 {
		return fqdns, fmt.Errorf("no answer for this address")
	}

	for _, answer := range in.Answer {
		if c, ok := answer.(*dns.CNAME); ok {
			fqdns = append(fqdns, c.Target)
		}
	}

	return fqdns, nil
}

func (a *application) Navigate() []Result {
	var (
		results []Result
		cfqdn   = a.fqdn
	)

	for {
		cnames, err := a.TypeCNAME()
		if err == nil && len(cnames) > 0 {
			cfqdn = cnames[0]

			continue
		}

		ips, err := a.TypeA()
		if err != nil {
			break
		}

		for _, ip := range ips {
			results = append(results, Result{IPAddress: ip, Hostname: cfqdn})
		}

		break
	}

	return results
}
