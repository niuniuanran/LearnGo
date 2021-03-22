package handlers

import (
	"github/niuniuanran/Day15/pkg/render"
	"net/http"
)

// ServeHome serves home page
func ServeHome(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home")
}

// ServeAbout serves about page
func ServeAbout(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about")
}
