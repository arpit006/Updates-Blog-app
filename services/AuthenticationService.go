package services

import (
	c "arpit006/web_app_with_go/constants"
	"arpit006/web_app_with_go/datastore"
	"arpit006/web_app_with_go/error_handler"
	"arpit006/web_app_with_go/models"
	"fmt"
	"log"
)

func AuthenticateUser(username, password string) (*models.User, error) {
	// 1. Get user-id from the user-name
	id, err := datastore.HGetStrInt(c.USER_BY_USERNAME, username)
	if err != nil {
		log.Printf("No user-id exists for the user: [%s]. Error is %s", username, err)
		return nil, error_handler.NewCustomError(err.Error(), c.AUTH_ERROR_PAGE, c.LOGIN_AGAIN, c.HTTP_UNAUTHORIZED)
	}
	// 2. prepare hash-key (user:id)
	hashKey := fmt.Sprintf(c.USER_BY_ID, id)
	// 3. prepare user object
	user := models.NewUser(hashKey, username)
	// authenticate the user
	authenticated := user.Authenticate(password)
	if authenticated {
		return user, nil
	}
	return nil, error_handler.NewCustomError(c.INVALID_PASSWORD, c.AUTH_ERROR_PAGE, c.LOGIN_AGAIN, c.HTTP_UNAUTHORIZED)
}
