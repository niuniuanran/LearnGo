package render

import (
	"bytes"
	"fmt"
	"github/niuniuanran/Day15/pkg/config"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{
}

var app *config.AppConfig

// NewTemplates sets the config for
func NewTemplates(a *config.AppConfig){
	app = a
}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	templatesCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.gohtml")
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		templateSet, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		matches, err := filepath.Glob("./templates/*.layout.gohtml")
		if err != nil {
			return nil, err
		}

		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.gohtml")
			if err != nil {
				return nil, err
			}
		}

		templatesCache[name] = templateSet
	}
	return templatesCache, nil
}


// RenderTemplate renders template
func RenderTemplate(w http.ResponseWriter, pageName string) {
	templatePath := pageName + ".page.gohtml"
	tc := app.TemplateCache
	t, ok := tc[templatePath]
	if !ok {
		log.Fatal("No template found for ", pageName)
	}

	buf := new(bytes.Buffer)
	_ = t.Execute(buf, nil)
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}
}