package routes

import (
	"net/http"

	"go-api/handlers"
)

func RegisterRoutes() {
	http.HandleFunc("/", handlers.Index)
	http.Handle("/register", handlers.AuthMiddleware(http.HandlerFunc(handlers.Register)))
	http.Handle("/mint", handlers.AuthMiddleware(http.HandlerFunc(handlers.Minter)))
	http.Handle("/balance", handlers.AuthMiddleware(http.HandlerFunc(handlers.Balancer)))
	http.Handle("/transfer", handlers.AuthMiddleware(http.HandlerFunc(handlers.Transferer)))
	http.Handle("/accountid", handlers.AuthMiddleware(http.HandlerFunc(handlers.ClientAccountIDer)))
	http.Handle("/initializer", handlers.AuthMiddleware(http.HandlerFunc(handlers.Initializer)))
	http.HandleFunc("/login", handlers.Login)
}