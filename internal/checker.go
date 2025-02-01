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

func checkEndpoint(endpoint string, results chan<- EndpointStatus, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(endpoint)
	status := "down"
	if err == nil && resp.StatusCode == 200 {
		status = "live"
	}

	results <- EndpointStatus{
		URL:    endpoint,
		Status: status,
	}
	log.Printf("Checked %s: %s\n", endpoint, status) // Logging
}

func (hc *HealthChecker) CheckHealth(endpoints []string) {
	var wg sync.WaitGroup
	hc.statuses = []EndpointStatus{}
	results := make(chan EndpointStatus)
	for _, endpoint := range endpoints {
		wg.Add(1)
		go func(endpoint string) {
			checkEndpoint(endpoint, results, &wg)
		}(endpoint)
	}

	// Make wg.Wait a goroutine to prevent unbuffered channel deadlock
	go func() {
		wg.Wait()      // wait for all go routines to finish by wg.Done()
		close(results) // close channel
	}()

	for status := range results {
		hc.statuses = append(hc.statuses, status)
	}

	inspectionCount += 1
	log.Printf("Inspection Count: %d", inspectionCount)
}

func (hc *HealthChecker) GetStatuses() []EndpointStatus {
	hc.mu.Lock()
	defer hc.mu.Unlock()
	return hc.statuses
}
