package sessions

import (
	"errors"
	"log"
	"net/http"
)

func ValidateSession(w http.ResponseWriter, r *http.Request) {
	_, err := GetUsernameFromSession(w, r)
	if err != nil {
		log.Printf("No Session created. Internal Error!. Please try again.")
		//http.Error(w, "No Session created. Internal Error!. Please Login again.", 401)
		http.Redirect(w, r, "/login", 302)
	}
}

func GetUsernameFromSession(w http.ResponseWriter, r *http.Request) (string, error) {
	sessionStore := GetSessionStoreFactory()
	session, err := sessionStore.Get(r, "session")
	if err != nil {
		log.Println("No Session created. Internal Error!. Please try again. Error is: ", err)
		return "", err
	}
	untyped, ok := session.Values["username"]
	if !ok {
		log.Println("Error in retrieving username from session. Error is ",  ok)
		return "", errors.New("could not retrieve Username from session")
	}
	username := untyped.(string)
	log.Printf("Username in session is [%s]", username)
	return username, nil
}

func SaveSession(w http.ResponseWriter, r *http.Request, username string) {
	sessionStore := GetSessionStoreFactory()
	session, err := sessionStore.Get(r, "session")
	if err != nil {
		log.Printf("Error in saving username [%s] to session. Error is [%s]", username, err)
		http.Redirect(w, r, "/login", 302)
	}
	session.Values["username"] = username
	session.Save(r, w)
	log.Printf("Session created sucessfully for user : [%s]", username)
}

func ClearSession(w http.ResponseWriter, r *http.Request, username string) {
	sessionStore := GetSessionStoreFactory()
	session, err := sessionStore.Get(r, "session")
	session.Options.MaxAge = -1
	session.Values["username"] = username
	if err != nil {
		log.Printf("Error in clearing Session for username [%s]. Error is [%s]", username, err)
		http.Redirect(w, r, "/login", 302)
	}
	session.Save(r, w)
}
