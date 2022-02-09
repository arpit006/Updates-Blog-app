package services

import (
	c "arpit006/web_app_with_go/constants"
	"arpit006/web_app_with_go/datastore"
	"arpit006/web_app_with_go/handlers"
	"arpit006/web_app_with_go/templ"
	"fmt"
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
		handlers.HandleAuthError(w, r)
		return
	}
	_, err = registerUser(name, hash)
	if err != nil {
		log.Printf("Error occurred while Registering the user. Please login/signup again!. Error is %s", err)
		handlers.HandleAuthError(w, r)
		return
	}
	http.Redirect(w, r, "/login", 301)
}

func registerUser(username string, hash []byte) (*User, error) {
	// get next id in redis
	id, err := datastore.Incr(c.USER_NEXT_ID)
	if err != nil {
		log.Printf("Error in getting next INCR key from Redis. Error is %s", err)
		return nil, err
	}
	hashKey := fmt.Sprintf(c.USER_BY_ID, id)
	// store data in redis in a pipeline
	pipe := datastore.StartPipeline()
	pipe.ExecInPipeStrInt(hashKey, c.ID, id)
	pipe.ExecInPipeStrStr(hashKey, c.USERNAME, username)
	pipe.ExecInPipeStrBytesArr(hashKey, c.PASSWORD_HASH, hash)
	// save username mapped to userid
	pipe.ExecInPipeStrInt(c.USER_BY_USERNAME, username, id)
	// Execute pipeline
	err = pipe.ExecPipeline()
	if err != nil {
		log.Printf("Error in running Redis Pipeline for user: [%s]. Error is %s", username, err)
		return nil, err
	}
	log.Printf("Registration request for user: [%s]", username)
	return NewUser(hashKey, username), nil
}
