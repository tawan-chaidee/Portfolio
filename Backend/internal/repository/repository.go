package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"restapi/internal/model"

	_ "github.com/mattn/go-sqlite3"
)

type ProjectRepository struct {
	db *sql.DB
}

func NewProjectRepository(dbFile string) (*ProjectRepository, error) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	// Create projects table if it doesn't exist
	query := `CREATE TABLE IF NOT EXISTS projects (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		description TEXT,
		url TEXT,
		tags TEXT
	);`
	_, err = db.Exec(query)
	if err != nil {
		return nil, fmt.Errorf("failed to create projects table: %v", err)
	}

	return &ProjectRepository{db: db}, nil
}

func (r *ProjectRepository) CreateProject(project *model.Project) error {
	tagsJSON, err := json.Marshal(project.Tags)
	if err != nil {
		return fmt.Errorf("failed to marshal tags: %v", err)
	}

	query := `INSERT INTO projects (title, description, url, tags) VALUES (?, ?, ?, ?)`
	result, err := r.db.Exec(query, project.Title, project.Description, project.URL, tagsJSON)
	if err != nil {
		return fmt.Errorf("failed to insert project: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert id: %v", err)
	}

	project.ID = int(id)
	return nil
}

func (r *ProjectRepository) GetProjectByID(id int) (*model.Project, error) {
	query := `SELECT id, title, description, url, tags FROM projects WHERE id = ?`
	row := r.db.QueryRow(query, id)

	var project model.Project
	var tagsJSON string
	if err := row.Scan(&project.ID, &project.Title, &project.Description, &project.URL, &tagsJSON); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("project not found")
		}
		return nil, fmt.Errorf("failed to scan project: %v", err)
	}

	if err := json.Unmarshal([]byte(tagsJSON), &project.Tags); err != nil {
		return nil, fmt.Errorf("failed to unmarshal tags: %v", err)
	}

	return &project, nil
}

func (r *ProjectRepository) UpdateProject(project *model.Project) error {
	tagsJSON, err := json.Marshal(project.Tags)
	if err != nil {
		return fmt.Errorf("failed to marshal tags: %v", err)
	}

	query := `UPDATE projects SET title = ?, description = ?, url = ?, tags = ? WHERE id = ?`
	_, err = r.db.Exec(query, project.Title, project.Description, project.URL, tagsJSON, project.ID)
	if err != nil {
		return fmt.Errorf("failed to update project: %v", err)
	}
	return nil
}

func (r *ProjectRepository) DeleteProject(id int) error {
	query := `DELETE FROM projects WHERE id = ?`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete project: %v", err)
	}
	return nil
}

func (r *ProjectRepository) GetAll() ([]*model.Project, error) {
	query := `SELECT id, title, description, url, tags FROM projects`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query projects: %v", err)
	}
	defer rows.Close()

	var projects []*model.Project
	for rows.Next() {
		var project model.Project
		var tagsJSON string
		if err := rows.Scan(&project.ID, &project.Title, &project.Description, &project.URL, &tagsJSON); err != nil {
			return nil, fmt.Errorf("failed to scan project: %v", err)
		}

		if err := json.Unmarshal([]byte(tagsJSON), &project.Tags); err != nil {
			return nil, fmt.Errorf("failed to unmarshal tags: %v", err)
		}

		projects = append(projects, &project)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	return projects, nil
}
