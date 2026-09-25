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
