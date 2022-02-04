package services

import (
	"arpit006/web_app_with_go/templ"
	"github.com/gorilla/sessions"
	"net/http"
)

// sessionStore using Cookie as our session store
var sessionStore = sessions.NewCookieStore([]byte("t0p-s3cr3t"))

func LoginGetHandler(w http.ResponseWriter, r *http.Request) {
	templates := templ.GetTemplateFactory()
	templates.ExecuteTemplate(w, "login.html", nil)
}

func LoginPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.PostForm.Get("username")
	session, _ := sessionStore.Get(r, "session")
	session.Values["username"] = name
	session.Save(r, w)
	http.Redirect(w, r, "/", 301)
}

// LoginTestHandler only to test if the username is getting saved in the session
func LoginTestHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessionStore.Get(r, "session")
	untyped, ok := session.Values["username"]
	if !ok {
		http.Error(w, "Username not saved to session", 500)
		return
	}
	username, ok := untyped.(string)
	if !ok {
		http.Error(w, "Username not parsed from session", 500)
		return
	}
	w.Write([]byte("Username is ::" +  username))
}
