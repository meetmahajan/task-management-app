package main

import (
	"log"
	"net/http"
	"task-management-backend/config"
	"task-management-backend/routes"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize MySQL database
	config.Connect()

	// Initialize router
	r := mux.NewRouter()

	// Register routes
	routes.RegisterTaskRoutes(r)

	// Start the server
	log.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}