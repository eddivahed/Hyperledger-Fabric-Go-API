package routes

import (
	"net/http"

	"go-api/handlers"
)

func RegisterRoutes() {
	// http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/mint", handlers.Minter)
	http.HandleFunc("/balance", handlers.Balancer)
	http.HandleFunc("/transfer", handlers.Transferer)
	http.HandleFunc("/accountid", handlers.ClientAccountIDer)
	http.HandleFunc("/initializer", handlers.Initializer)
}