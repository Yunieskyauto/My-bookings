package config

import "html/template"

// AppConfig hols the application config
type AppConfig struct {
	TemplateCache map[string]*template.Template
}
