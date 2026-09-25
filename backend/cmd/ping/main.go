package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	Endpoint  string
	Status    string
	LatencyMs int64
	Timestamp time.Time
	Err       error
}

func checkEndpoint(url string, wg *sync.WaitGroup, results chan<- Result) {
	defer wg.Done()

	client := http.Client{Timeout: 5 * time.Second}
	start := time.Now()

	resp, err := client.Get(url)
	latency := time.Since(start).Milliseconds()

	if err != nil {
		results <- Result{
			Endpoint:  url,
			Status:    "DOWN",
			LatencyMs: latency,
			Timestamp: start,
			Err:       err,
		}

		return
	}

	defer resp.Body.Close()

	status := "UP"
	if resp.StatusCode >= 400 {
		status = "DOWN"
	}

	results <- Result{
		Endpoint:  url,
		Status:    status,
		LatencyMs: latency,
		Timestamp: start,
	}
}

func main() {

	endpoints := []string{
		"https://www.google.com",
		"https://www.caida.org",
		"https://www.daniellov.com",
		"https://fakeurl.com",
	}

	var wg sync.WaitGroup
	results := make(chan Result, len(endpoints))

	for _, url := range endpoints {
		wg.Add(1)
		go checkEndpoint(url, &wg, results)
	}

	wg.Wait()
	close(results)

	for r := range results {
		if r.Err != nil {
			fmt.Printf("[%s] %s - %dms - ERROR: %v\n", r.Status, r.Endpoint, r.LatencyMs, r.Err)
		} else {
			fmt.Printf("[%s] %s - %dms\n", r.Status, r.Endpoint, r.LatencyMs)
		}
	}
}
