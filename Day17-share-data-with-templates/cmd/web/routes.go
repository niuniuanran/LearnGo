package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github/niuniuanran/Day17/pkg/config"
	"github/niuniuanran/Day17/pkg/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler{
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Get("/", handlers.Repo.ServeHome)
	mux.Get("/about", handlers.Repo.ServeAbout)
	return mux
}
