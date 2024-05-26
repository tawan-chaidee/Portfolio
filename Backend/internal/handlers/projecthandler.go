package handlers

import (
	"encoding/json"
	"net/http"
	"restapi/internal/model"
	"restapi/internal/repository"
)

// ProjectHandler handles requests related to projects
type ProjectHandler struct {
	projectRepo *repository.ProjectRepository
}

// NewProjectHandler creates a new instance of ProjectHandler
func NewProjectHandler(projectRepo *repository.ProjectRepository) *ProjectHandler {
	return &ProjectHandler{projectRepo}
}

// GetAllProjects handles both GET and POST requests for projects
func (h *ProjectHandler) GetAllProjects(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Handle GET request to fetch all projects
		projects, err := h.projectRepo.GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Convert projects to JSON and send response
		jsonProjects, err := json.Marshal(projects)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonProjects)

	case http.MethodPost:
		// Handle POST request to create a new project
		var project model.Project
		if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		if err := h.projectRepo.Create(&project); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)

	default:
		// Handle unsupported methods (optional)
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}
