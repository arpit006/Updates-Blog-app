package services

import (
	"arpit006/web_app_with_go/datastore"
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
	// match with the encrypted password used during signup
	hash, err := datastore.GetBytesFromRedis("user:" + name)
	if err != nil {
		log.Println("Could not retrieve Saved password!")
		//http.Error(w, "Username and Password do not match. Unauthorized", 401)
		http.Redirect(w, r, "/auth-error", 302)
		return
	}
	err = ValidatePassword(hash, password)
	if err != nil {
		log.Printf("Wrong Password for username: %s", name)
		//http.Error(w, "Incorrect Password!", 401)
		http.Redirect(w, r, "/auth-error", 302)
		return
	}
	sessions2.SaveSession(w, r, name)
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
