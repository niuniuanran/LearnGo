package handlers

import (
	"github/niuniuanran/Day17/pkg/config"
	"github/niuniuanran/Day17/pkg/models"
	"github/niuniuanran/Day17/pkg/render"
	"net/http"
)

var Repo *Repository

// Repository pattern
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// ServeHome serves home page
func (m *Repository) ServeHome(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home", &models.TemplateData{})
}

// ServeAbout serves about page
func (m *Repository) ServeAbout(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, world"
	render.RenderTemplate(w, "about", &models.TemplateData{StringMap: stringMap})
}
