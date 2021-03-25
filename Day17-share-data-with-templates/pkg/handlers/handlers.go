package handlers

import (
	"github/niuniuanran/Day17/pkg/config"
	"github/niuniuanran/Day17/pkg/render"
	"net/http"
)

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
}

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
	render.RenderTemplate(w, "home")
}

// ServeAbout serves about page
func (m *Repository) ServeAbout(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about")
}
