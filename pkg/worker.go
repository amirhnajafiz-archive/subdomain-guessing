package pkg

func Worker(tracker chan Empty, fqdns chan string, gather chan []Result, serverAddr string) {
	for fqdn := range fqdns {
		app := NewApp(fqdn, serverAddr)
		results := app.Navigate()
		if len(results) > 0 {
			gather <- results
		}
	}

	var e Empty
	tracker <- e
}
