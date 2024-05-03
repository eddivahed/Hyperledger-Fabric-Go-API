package main

import (
    "log"
    "net/http"
    
    "go-api/routes"
)

func main() {
    // Initialize and configure dependencies
    routes.RegisterRoutes()

    // Start the server
    log.Fatal(http.ListenAndServe(":8082", nil))
}