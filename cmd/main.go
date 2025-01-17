package main

import (
	"healthchecker/config"
	"healthchecker/internal/health"
	"healthchecker/web"
	"log"
	"os"

	"github.com/alecthomas/kingpin"
)

func main() {
	app := kingpin.New("healthchecker", "A healthchecker app made using golang.")
	configFile := app.Arg("config", "Path to the config file").Required().String()

	kingpin.MustParse(app.Parse(os.Args[1:]))

	cfg, err := config.LoadConfig(*configFile)
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
