package ping

import (
	"fmt"
	"sync"
	"testing"
)

func runPings(endpoints []string) []string {
	var wg sync.WaitGroup
	results := make(chan Result, len(endpoints))

	for _, url := range endpoints {
		wg.Add(1)
		go checkEndpoint(url, &wg, results)
	}
	wg.Wait()
	close(results)

	var ping_results []string

	for r := range results {
		if r.Err != nil {
			ping_results = append(ping_results, fmt.Sprintf("[%s] %s - %dms - ERROR: %v\n", r.Status, r.Endpoint, r.LatencyMs, r.Err))
		} else {
			ping_results = append(ping_results, fmt.Sprintf("[%s] %s - %dms\n", r.Status, r.Endpoint, r.LatencyMs))
		}
	}

	return ping_results
}

func TestPing(t *testing.T) {
	endpoints := []string{
		"https://www.google.com",
		"https://www.caida.org",
		"https://www.daniellov.com",
		"https://fakeurl.com",
	}

	fmt.Println(runPings(endpoints))
}
