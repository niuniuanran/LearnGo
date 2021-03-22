package config

import (
	"html/template"
	"log"
)

// AppConfig holds the site-wide application config
type AppConfig struct {
	TemplateCache map[string]*template.Template
	InfoLog log.Logger
}
