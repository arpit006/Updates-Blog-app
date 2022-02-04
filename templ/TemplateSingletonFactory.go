package templ

import (
	"html/template"
	"log"
)

var templates *template.Template

func GetTemplateFactory() *template.Template {
	if templates != nil {
		return templates
	}
	templates = template.Must(template.ParseGlob("templates/*.html"))
	log.Println("HTTP Template initialized!.")
	return templates
}
