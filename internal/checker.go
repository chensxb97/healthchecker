package internal

import (
	"log"
	"net/http"
	"sync"
)

var inspectionCount = 0

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
	var wg sync.WaitGroup
	hc.statuses = []EndpointStatus{}

	for _, endpoint := range endpoints {
		wg.Add(1)
		go func(endpoint string) {
			defer wg.Done()
			resp, err := http.Get(endpoint)
			status := "down"
			if err == nil && resp.StatusCode == 200 {
				status = "live"
			}
			hc.mu.Lock()
			hc.statuses = append(hc.statuses, EndpointStatus{
				URL:    endpoint,
				Status: status,
			})
			log.Printf("Checked %s: %s\n", endpoint, status) // Logging
			hc.mu.Unlock()
		}(endpoint)
	}
	wg.Wait()
	inspectionCount += 1
	log.Printf("Checked all endpoints!")
	log.Printf("Inspection Count: %d", inspectionCount)
}

func (hc *HealthChecker) GetStatuses() []EndpointStatus {
	hc.mu.Lock()
	defer hc.mu.Unlock()
	return hc.statuses
}
