package monitor

import (
	"encoding/json"
	"fmt"

	"daniellov.com/health-monitor/internal/ping"
)

type Monitor struct {
	id       int64
	endpoint string
	history  []ping.Result
}

var id int64 = 0

func SetMonitorConfig(encoder *json.Encoder) {
	encoder.SetIndent("", "    ")
}

func New(endpoint string) (*Monitor, bool) {
	id += 1

	return &Monitor{
		id:       id,
		endpoint: endpoint,
	}, false
}

// return value: error
func (m *Monitor) Run() (ping.Result, bool) {
	result, err := ping.CheckEndpoint(m.endpoint)

	fmt.Printf("monitor (%d) @ (%v): [%s] %s in %d ms, err: %v \n", m.id, result.Timestamp, result.Status, result.Endpoint, result.LatencyMs, result.Err)
	if err {
		return ping.Result{}, true
	}

	m.history = append(m.history, result)
	return result, false
}

func (m *Monitor) History() []ping.Result {
	return m.history
}
