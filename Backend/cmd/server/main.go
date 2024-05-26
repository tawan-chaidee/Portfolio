package main

import (
	"log"
	"net/http"

	"restapi/internal/handlers"
	"restapi/internal/repository"
)

func main() {
	// Initialize the repository
	repo, err := repository.NewProjectRepository("database.db")
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Initialize the project handler
	projectHandler := handlers.NewProjectHandler(repo)

	// Set up HTTP routes
	http.HandleFunc("/projects", projectHandler.GetAllProjects)

	// Start the server
	log.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
