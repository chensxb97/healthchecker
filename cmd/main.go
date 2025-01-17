package main

import (
	"healthchecker/config"
	"healthchecker/internal/health"
	"healthchecker/web"
	"log"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize health checker
	healthChecker := health.NewHealthChecker()
	go healthChecker.CheckHealth(getURLs(cfg.Endpoints))

	// Start web server
	server := web.NewServer(healthChecker)
	server.Start()
}

func getURLs(endpoints []config.Endpoint) []string {
	urls := make([]string, len(endpoints))
	for i, endpoint := range endpoints {
		urls[i] = endpoint.URL
	}
	return urls
}
