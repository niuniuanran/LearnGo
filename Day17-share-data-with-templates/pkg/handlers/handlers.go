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
	remoteIP := r.RemoteAddr
	// every time someone hits the homepage, store the ip address into their session
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home", &models.TemplateData{})
}

// ServeAbout serves about page
func (m *Repository) ServeAbout(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, world"

	remoteIP := m.App.Session.GetString(r.Context(),"remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about", &models.TemplateData{StringMap: stringMap})
}
