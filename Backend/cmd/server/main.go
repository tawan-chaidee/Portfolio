package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"restapi/internal/handlers"
	"restapi/internal/repository"

	_ "github.com/mattn/go-sqlite3"
)

// Port we listen on.
const portNum string = ":8000"

func main() {
	log.Println("Starting our simple http server.")

	// Path to the SQLite database file
	dbPath := "path_to_your_sqlite_database.db"

	// Connect to your SQLite database
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close() // Defer closing the database connection until the function exits

	// Ensure the projects table exists in the database
	err = repository.CreateProjectsTable(db)
	if err != nil {
		log.Fatal("Error creating projects table:", err)
	}

	// Creating a project repository with the database connection
	projectRepo := repository.NewProjectRepository(db)

	// Creating a new instance of ProjectHandler with the project repository
	projectHandler := handlers.NewProjectHandler(projectRepo)

	// Registering handler functions for projects
	http.HandleFunc("/projects", projectHandler.GetAllProjects) // GET request for fetching all projects

	log.Println("Started on port", portNum)
	fmt.Println("To close connection CTRL+C :-)")

	// Spinning up the server.
	err = http.ListenAndServe(portNum, nil)
	if err != nil {
		log.Fatal(err)
	}
}
