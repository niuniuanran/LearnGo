package main

import (
	"fmt"
	"github/niuniuanran/Day17/pkg/config"
	"github/niuniuanran/Day17/pkg/handlers"
	"github/niuniuanran/Day17/pkg/render"
	"log"
	"net/http"
)

const addr = ":8080"

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
	srv := &http.Server{
		Addr: addr,
		Handler: routes(&app),
	}
	fmt.Printf("Starting application on port %s\n", addr)
	log.Fatal(srv.ListenAndServe())
}
