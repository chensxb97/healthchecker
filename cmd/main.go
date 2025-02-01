package main

import (
	"healthchecker/config"
	"healthchecker/internal"
	"healthchecker/web"
	"log"
	"os"
	"time"

	"github.com/alecthomas/kingpin"
)

func main() {
	// Handle input argument
	app := kingpin.New("healthchecker", "A healthchecker app made using golang.")
	configFile := app.Arg("config", "Path to the config file").Required().String()
	kingpin.MustParse(app.Parse(os.Args[1:]))

	// Load config
	err := config.LoadConfig(*configFile)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize health checker
	healthChecker := internal.NewHealthChecker()
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()
		for range ticker.C {
			go healthChecker.CheckHealth(config.GetEndpoints())
		}
	}()

	// Start web server
	server := web.NewServer(healthChecker)
	server.Start()
}
