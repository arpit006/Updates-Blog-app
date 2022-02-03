package services

import (
	"arpit006/web_app_with_go/datastore"
	"html/template"
	"net/http"
)

var templates *template.Template

func IndexGetHandler(w http.ResponseWriter, r *http.Request) {
	templates = template.Must(template.ParseGlob("templates/*.html"))
	comments, err := datastore.GetRangeFromRedis("comments", 0, 10)
	if err != nil {
		panic(err)
	}
	templates.ExecuteTemplate(w, "index.html", comments)
}

func IndexPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	comment := r.PostForm.Get("comment")
	datastore.LPushToRedis("comments", comment)
	http.Redirect(w, r, "/", 302)
}
