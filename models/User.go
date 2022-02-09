package models

import (
	"arpit006/web_app_with_go/datastore"
	"arpit006/web_app_with_go/services"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

// User user model
type User struct {
	Key string
	UserName string
}

func (user *User) GetUsername() (string, error) {
	return datastore.HGetStrStr(user.Key, "username")
}

func (user *User) GetHash() ([]byte, error) {
	return datastore.HGetStrBytesArr(user.Key, "hash")
}

func GetUserByUsername(username string) (*User, error) {
	id, err := datastore.HGetStrInt("user:by-username", username)
	if err == redis.Nil {
		// user not found
	} else if err != nil {
		return nil, err
	}
	key := fmt.Sprintf("user:%d", id)
	return &User{Key: key, UserName: username}, err
}

func (user *User) AuthenticateUser(password string) error {
	//hash, err := user.GetHash()
	_, err := user.GetHash()
	if err != nil {
		log.Printf("No password found for user [%s:%s]. Error is %s", user.Key, user.UserName, err)
		return err
	}
	//TODO
	err = services.Authenticate("", password)
	return err
}

