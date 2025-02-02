package main

import (
	"bank-service/config"
	"bank-service/internal/handler"
	"bank-service/internal/repository"
	"bank-service/internal/routes"
	"log"
	"net/http"
)

func main() {
	// Initialize DB
	config.ConnectDB()
	config.InitalBalance()

	// Initialize repository & handlers
	repo := repository.NewAccountRepository(config.DB)
	handler := handlers.NewAccountHandler(repo)

	// Setup routes
	router := routes.SetupRoutes(handler)

	// Start server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
