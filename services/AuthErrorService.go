package services

import (
	"arpit006/web_app_with_go/templ"
	"net/http"
)

func AuthErrorHandler(w http.ResponseWriter, r *http.Request) {
	templates := templ.GetTemplateFactory()
	templates.ExecuteTemplate(w, "auth-error.html", nil)
}
