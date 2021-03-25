package render

import (
	"fmt"
	"github/niuniuanran/Day17/pkg/config"
	"github/niuniuanran/Day17/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// RenderTemplate renders template
func RenderTemplate(w http.ResponseWriter, pageName string, td *models.TemplateData) {
	templatePath := pageName + ".page.gohtml"
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[templatePath]
	if !ok {
		log.Fatal("No template found for ", pageName)
	}

	// err := t.ExecuteTemplate(w, templatePath,nil)
	err := t.Execute(w, td)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}
}

// SetAppConfig sets the config
func SetAppConfig(a *config.AppConfig) {
	app = a
}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	templatesCache := map[string]*template.Template{}
	// find all page template files
	pages, err := filepath.Glob("./templates/*.page.gohtml")
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		// The file name itself, for example, home.page.gohtml
		name := filepath.Base(page)
		// New allocates a new HTML template with the given name.
		// Funcs adds the elements of the argument map to the template's function map. It must be called before the template is parsed.
		templateSet, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		matches, err := filepath.Glob("./templates/*.layout.gohtml")
		if err != nil {
			return nil, err
		}

		if len(matches) > 0 {
			// Associate all layout templates with the parsed page template.
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.gohtml")
			if err != nil {
				return nil, err
			}
		}

		templatesCache[name] = templateSet
	}
	return templatesCache, nil
}
