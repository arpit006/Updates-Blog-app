package services

import (
	"arpit006/web_app_with_go/models"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func Authenticate(username string, password string) error {
	user, err := models.GetUserByUsername(username)
	if err != nil {
		log.Printf("No User exists for this username [%s]. Error is %s", username, err)
		// TODO: throw user not exists exception
		return err
	}
	hash, err := user.GetHash()
	if err != nil {
		log.Printf("No password saved for the user: [%s]. Error is %s", username, err)
	}
	return ValidatePassword(hash, password)
}

func ValidatePassword(hash []byte, password string) error {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err != nil {
		log.Println("Password Invalid!")
		return err
	}
	return nil
}
