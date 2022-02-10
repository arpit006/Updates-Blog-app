package services

import (
	"arpit006/web_app_with_go/error_handler"
	sessions2 "arpit006/web_app_with_go/sessions"
	"arpit006/web_app_with_go/templ"
	"log"
	"net/http"
)


func LoginGetHandler(w http.ResponseWriter, r *http.Request) {
	templates := templ.GetTemplateFactory()
	templates.ExecuteTemplate(w, "login.html", nil)
}

func LoginPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	user, err := AuthenticateUser(name, password)
	if err != nil {
		error_handler.HandleAuthError(w, r)
		return
	}
	sessions2.SaveSession(w, r, user)
	http.Redirect(w, r, "/", 302)
}

// LoginTestHandler only to test if the username is getting saved in the session
func LoginTestHandler(w http.ResponseWriter, r *http.Request) {
	sessions2.ValidateSession(w, r)
	username, err := sessions2.GetUsernameFromSession(w, r)
	if err != nil {
		log.Println("Could not retrieve Username from Session. Error is ", err)
		http.Redirect(w, r, "/login", 302)
	}
	w.Write([]byte("Username is ::" +  username))
}
