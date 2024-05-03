package main

import (
	"fmt"
	"log"
	"net/http"

	"go-api/config"
	"go-api/routes"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	routes.RegisterRoutes()

	log.Printf("Server running on port %d", cfg.Server.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Server.Port), nil))
}