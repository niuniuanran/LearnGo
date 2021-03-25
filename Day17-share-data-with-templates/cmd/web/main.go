package main

import (
	"github/niuniuanran/Day17/pkg/config"
	"github/niuniuanran/Day17/pkg/handlers"
	"github/niuniuanran/Day17/pkg/render"
	"log"
	"net/http"
)

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	app.TemplateCache = tc
	app.UseCache = false
	render.SetAppConfig(&app)
	http.HandleFunc("/", handlers.Repo.ServeHome)
	http.HandleFunc("/about", handlers.Repo.ServeAbout)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
