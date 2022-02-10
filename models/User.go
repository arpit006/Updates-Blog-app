package models

import (
	c "arpit006/web_app_with_go/constants"
	"arpit006/web_app_with_go/datastore"
	"golang.org/x/crypto/bcrypt"
	"log"
)

// User object
type User struct {
	HashKey string
	Username string
}
// NewUser User constructor
func NewUser(hashKey, username string) *User {
	return &User{HashKey: hashKey, Username: username}
}

func (user *User) GetUserHashKey() string {
	return user.HashKey
}

func (user *User) GetUsername() string {
	return user.Username
}

func (user *User) GetPasswordHash() ([]byte, error) {
	hashKey := user.GetUserHashKey()
	hash, err := datastore.HGetStrBytesArr(hashKey, c.PASSWORD_HASH)
	if err != nil {
		log.Printf("No password found for user: [%s]. Please login/signup again. Error is %s", user.GetUsername(), err)
		return nil, err
	}
	return hash, nil
}

func (user *User) GetUserId() (int64, error) {
	// user:by-username <username : id>
	id, err := datastore.HGetStrInt(c.USER_BY_USERNAME, user.GetUsername())
	if err != nil {
		log.Printf("No UserId found for user: [%s]. Error is %s", user.GetUsername(), err)
		return -1, err
	}
	return id, nil
}

func (user *User) Authenticate(password string) bool {
	passwordHash, err := user.GetPasswordHash()
	if err != nil {
		log.Printf("Error getting password for the user: [%s]. Error is %s", user.GetUsername(), err)
		return false
	}
	err = bcrypt.CompareHashAndPassword(passwordHash, []byte(password))
	if err != nil {
		log.Printf("Username and Password did not match for user: [%s]. Error is %s", user.GetUsername(), err)
		return false
	}
	return true
}
