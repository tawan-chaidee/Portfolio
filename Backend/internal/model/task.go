package model

// Project represents a project in the portfolio
type Project struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	URL         string   `json:"url"`
	Tags        []string `json:"tags"`
}

// NewProject creates a new Project instance
func NewProject(id int, title, description, url string, tags []string) *Project {
	return &Project{
		ID:          id,
		Title:       title,
		Description: description,
		URL:         url,
		Tags:        tags,
	}
}
