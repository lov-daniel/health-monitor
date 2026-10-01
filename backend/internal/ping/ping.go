package ping

import (
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

// result val: result, error
func CheckEndpoint(endpoint string) (Result, bool) {
	var wg sync.WaitGroup
	results := make(chan Result, len(endpoint))

	wg.Add(1)
	go runEndpoint(endpoint, &wg, results)

	wg.Wait()
	close(results)

	var return_val Result

	for result := range results {
		return_val = result
	}

	return return_val, false
}

// result val: result, error
func CheckAll(endpoints []string) ([]Result, bool) {

	var wg sync.WaitGroup
	results := make(chan Result, len(endpoints))

	for _, url := range endpoints {
		wg.Add(1)
		go runEndpoint(url, &wg, results)
	}

	wg.Wait()
	close(results)

	var return_val []Result

	for result := range results {
		return_val = append(return_val, result)
	}

	return return_val, false
}

func runEndpoint(url string, wg *sync.WaitGroup, results chan<- Result) {
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
