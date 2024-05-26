package repository

import (
	"database/sql"
	"log"
	"restapi/internal/model"
	"strings"
)

// ProjectRepository provides methods to interact with the project database
type ProjectRepository struct {
	db *sql.DB
}

// NewProjectRepository creates a new instance of ProjectRepository
func NewProjectRepository(db *sql.DB) *ProjectRepository {
	return &ProjectRepository{db}
}

// CreateProjectsTable creates the projects table if it doesn't exist
func CreateProjectsTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS projects (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        description TEXT,
        url TEXT,
        tags TEXT
    )`)
	return err
}

// GetAll retrieves all projects from the database
func (r *ProjectRepository) GetAll() ([]*model.Project, error) {
	var projects []*model.Project

	rows, err := r.db.Query("SELECT id, title, description, url, tags FROM projects")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var project model.Project
		err := rows.Scan(&project.ID, &project.Title, &project.Description, &project.URL, &project.Tags)
		if err != nil {
			return nil, err
		}
		projects = append(projects, &project)
	}

	return projects, nil
}

// Create inserts a new project into the database
func (r *ProjectRepository) Create(project *model.Project) error {
	// Serialize tags slice into a comma-separated string
	tagsString := strings.Join(project.Tags, ",")

	// Insert project into the database
	_, err := r.db.Exec("INSERT INTO projects (title, description, url, tags) VALUES (?, ?, ?, ?)", project.Title, project.Description, project.URL, tagsString)
	if err != nil {
		log.Println("Error inserting project:", err)
		return err
	}
	return nil
}
