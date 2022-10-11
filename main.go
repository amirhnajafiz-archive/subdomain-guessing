package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/amirhnajafiz/subdomain-guessing/pkg"
)

func main() {
	var (
		flDomain      = flag.String("domain", "", "This domain to perform guessing against.")
		flWordlist    = flag.String("wordlist", "", "The wordlist to use for guessing.")
		flWorkerCount = flag.Int("c", 100, "The amount of workers to use.")
		flServerAddr  = flag.String("server", "8.8.8.8:53", "The DNS server to use.")
	)

	flag.Parse()

	if *flDomain == "" || *flWordlist == "" {
		fmt.Println("-domain and -wordlist are required")
		os.Exit(1)
	}

	var results []pkg.Result

	fqdns := make(chan string, *flWorkerCount)
	gather := make(chan []pkg.Result)
	tracker := make(chan pkg.Empty)

	fh, err := os.Open(*flWordlist)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fh)

	for i := 0; i < *flWorkerCount; i++ {
		go pkg.Worker(tracker, fqdns, gather, *flServerAddr)
	}

	for scanner.Scan() {
		fqdns <- fmt.Sprintf("%s.%s", scanner.Text(), *flDomain)
	}

	go func() {
		for r := range gather {
			results = append(results, r[0])
		}

		var e pkg.Empty

		tracker <- e
	}()

	close(fqdns)

	for i := 0; i < *flWorkerCount; i++ {
		<-tracker
	}

	close(gather)

	<-tracker

	w := tabwriter.NewWriter(os.Stdout, 0, 8, 0, 0, 0)
	for _, r := range results {
		fmt.Fprintf(w, "%s \"%s\"\n", r.Hostname, r.IPAddress)
	}
	w.Flush()
}
