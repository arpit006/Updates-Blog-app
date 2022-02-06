package services

import (
	"arpit006/web_app_with_go/sessions"
	"log"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	sessions.ValidateSession(w, r)
	username, err := sessions.GetUsernameFromSession(w, r)
	if err != nil {
		log.Printf("Error in getting the logged in user from the session. Error is %s", err)
		http.Redirect(w, r, "/login", 302)
	}
	log.Printf("Logging out user [%s] from our Go-Web-App", username)
	sessions.ClearSession(w, r, username)
	http.Redirect(w, r, "/login", 302)
}
