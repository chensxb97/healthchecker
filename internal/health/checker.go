package health

import (
	"net/http"
	"sync"
	"time"
)

type EndpointStatus struct {
	URL    string `json:"url"`
	Status string `json:"status"`
}

type HealthChecker struct {
	statuses []EndpointStatus
	mu       sync.Mutex
}

func NewHealthChecker() *HealthChecker {
	return &HealthChecker{
		statuses: []EndpointStatus{},
	}
}

func (hc *HealthChecker) CheckHealth(endpoints []string) {
	for {
		hc.mu.Lock()
		hc.statuses = []EndpointStatus{}

		for _, endpoint := range endpoints {
			resp, err := http.Get(endpoint)
			status := "down"
			if err == nil && resp.StatusCode == 200 {
				status = "live"
			}
			hc.statuses = append(hc.statuses, EndpointStatus{
				URL:    endpoint,
				Status: status,
			})
		}

		hc.mu.Unlock()
		time.Sleep(30 * time.Second)
	}
}

func (hc *HealthChecker) GetStatuses() []EndpointStatus {
	hc.mu.Lock()
	defer hc.mu.Unlock()
	return hc.statuses
}
