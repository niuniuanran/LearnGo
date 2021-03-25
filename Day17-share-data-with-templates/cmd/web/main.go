package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github/niuniuanran/Day17/pkg/config"
	"github/niuniuanran/Day17/pkg/handlers"
	"github/niuniuanran/Day17/pkg/render"
	"log"
	"net/http"
	"time"
)

const addr = ":8080"
var app config.AppConfig
var session *scs.SessionManager

func main() {
	app.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

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
