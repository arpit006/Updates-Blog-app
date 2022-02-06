package services

import (
	"arpit006/web_app_with_go/datastore"
	"arpit006/web_app_with_go/templ"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func RegisterGetHandler(w http.ResponseWriter, r *http.Request) {
	templates := templ.GetTemplateFactory()
	templates.ExecuteTemplate(w, "register.html", nil)
}

func RegisterPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	cost := bcrypt.DefaultCost
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Println("Error in hashing the user password for username : ", name)
		http.Error(w, "Could not hash the password. Please retry", 500)
		return
	}
	datastore.SaveBytesToRedis("user:" + name, hash)
	http.Redirect(w, r, "/login", 301)
}
