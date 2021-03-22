package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// RenderTemplate renders template
func RenderTemplate(w http.ResponseWriter, pageName string) {
	templatePath := pageName + ".page.gohtml"
	templateCache, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	t, ok := templateCache[templatePath]
	if !ok {
		log.Fatal("No template found for ", pageName)
	}

	buf := new(bytes.Buffer)
	_ = t.Execute(buf, nil)
	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}
}

var functions = template.FuncMap{
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
		fmt.Println(name)
		templatesCache[name] = templateSet
	}
	return templatesCache, nil
}
