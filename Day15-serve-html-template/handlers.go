package main

import (
	"net/http"
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home")
}

func serveAbout(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about")
}
