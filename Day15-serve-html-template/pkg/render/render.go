package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// RenderTemplate renders template
func RenderTemplate(w http.ResponseWriter, tmplPath string) {
	parsedTemplate, err := template.ParseFiles("./templates/" + tmplPath + ".page.tmpl")
	if err != nil {
		fmt.Println("Error parsing template", tmplPath)
		return
	}
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error executing template", tmplPath)
	}
}
