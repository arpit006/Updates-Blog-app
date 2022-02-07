package services

import (
	"arpit006/web_app_with_go/datastore"
	"arpit006/web_app_with_go/templ"
	"net/http"
)

func IndexGetHandler(w http.ResponseWriter, r *http.Request) {
	//sessions2.ValidateSession(w, r)
	templates := templ.GetTemplateFactory()
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
