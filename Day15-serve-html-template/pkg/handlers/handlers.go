package handlers

import (
	"github/niuniuanran/Day15/pkg/config"
	"github/niuniuanran/Day15/pkg/render"
	"net/http"
)

// Repository pattern
type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

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
	render.RenderTemplate(w, "home")
}

// ServeAbout serves about page
func (m *Repository) ServeAbout(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about")
}
