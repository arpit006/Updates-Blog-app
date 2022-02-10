package services

import (
	c "arpit006/web_app_with_go/constants"
	"arpit006/web_app_with_go/datastore"
	"fmt"
	"log"
)

func GetUserNameFromUserId(id int64) (string, error) {
	hashKey := fmt.Sprintf(c.USER_BY_ID, id)
	username, err := datastore.HGetStrStr(hashKey, c.USERNAME)
	if err != nil {
		log.Printf("Error in getting username for id :[%d]. Error is %s", id, err)
		return "", err
	}
	return username, nil
}
