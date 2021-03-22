package main

import (
	"github/niuniuanran/Day15/pkg/config"
	"github/niuniuanran/Day15/pkg/handlers"
	"github/niuniuanran/Day15/pkg/render"
	"log"
	"net/http"
)

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = tc
	render.NewTemplates(&app)
	http.HandleFunc("/", handlers.ServeHome)
	http.HandleFunc("/about", handlers.ServeAbout)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
