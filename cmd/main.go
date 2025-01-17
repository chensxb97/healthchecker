package main

import (
	"healthchecker/config"
	"healthchecker/internal"
	"healthchecker/web"
	"log"
	"os"

	"github.com/alecthomas/kingpin"
)

func main() {
	app := kingpin.New("healthchecker", "A healthchecker app made using golang.")
	configFile := app.Arg("config", "Path to the config file").Required().String()

	kingpin.MustParse(app.Parse(os.Args[1:]))

	err := config.LoadConfig(*configFile)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize health checker
	healthChecker := internal.NewHealthChecker()
	go healthChecker.CheckHealth(config.GetEndpoints())

	// Start web server
	server := web.NewServer(healthChecker)
	server.Start()
}
