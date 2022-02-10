package sessions

import (
	c "arpit006/web_app_with_go/constants"
	"arpit006/web_app_with_go/datastore"
	"arpit006/web_app_with_go/error_handler"
	"arpit006/web_app_with_go/models"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func ValidateSession(w http.ResponseWriter, r *http.Request) {
	_, err := GetUserIdFromSession(w, r)
	if err != nil {
		log.Printf("No Session created. Internal Error!. Please try again.")
		//http.Error(w, "No Session created. Internal Error!. Please Login again.", 401)
		http.Redirect(w, r, "/login", 302)
	}
}

func GetUsernameFromSession(w http.ResponseWriter, r *http.Request) (string, error) {
	userID, err := GetUserIdFromSession(w, r)
	if err != nil {
		log.Printf("No userID exists in the session. Please login again")
		return "", err
	}
	hashKey := fmt.Sprintf(c.USER_BY_ID, userID)
	username, err := datastore.HGetStrStr(hashKey, c.USERNAME)
	if err != nil {
		log.Printf("Error getting username for userId: [%d]. Error is %s", userID, username)
		return "", err
	}
	return username, nil
}

func GetUserIdFromSession(w http.ResponseWriter, r *http.Request) (int64, error) {
	sessionStore := GetSessionStoreFactory()
	session, err := sessionStore.Get(r, "session")
	if err != nil {
		log.Println("No Session created. Internal Error!. Please try again. Error is: ", err)
		return -1, err
	}
	untyped, ok := session.Values[c.USER_ID]
	if !ok {
		log.Println("Error in retrieving userId from session. Error is ",  ok)
		return -1, errors.New("could not retrieve userId from session")
	}
	userId := untyped.(int64)
	log.Printf("UserID in session is [%d]", userId)
	return userId, nil
}

func SaveSession(w http.ResponseWriter, r *http.Request, user *models.User) {
	sessionStore := GetSessionStoreFactory()
	session, err := sessionStore.Get(r, c.SESSION)
	userId, err := user.GetUserId()
	if err != nil {
		log.Printf("Error in saving user [%s:%d] to session. Error is [%s]", user.GetUsername(), userId, err)
		error_handler.HandleAuthError(w, r)
	}
	session.Values[c.USER_ID] = userId
	session.Save(r, w)
	log.Printf("Session created sucessfully for user : [%s:%d]", user.GetUsername(), userId)
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
