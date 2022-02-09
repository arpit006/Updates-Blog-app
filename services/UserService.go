package services

import (
	"arpit006/web_app_with_go/datastore"
	"arpit006/web_app_with_go/models"
	"context"
	"fmt"
	"log"
)

var ctx context.Context = datastore.ContextFactory()

func NewUser(username string, hash []byte) (*models.User, error) {
	// gets next integer key available in redis
	id, err := datastore.Incr()
	if err != nil {
		log.Printf("Error getting Incr")
		return nil, err
	}
	key := fmt.Sprintf("user:%d", id)
	// Bulk Redis update
	redisPipe := datastore.StartPipeline()

	redisPipe.ExecInPipeStrInt(key, "id", id)
	redisPipe.ExecInPipeStrStr(key, "username", username)
	redisPipe.ExecInPipeStrBytesArr(key, "hash", hash)
	// user by username
	redisPipe.ExecInPipeStrInt("user:by-username", username, id)
	// Execute Redis Pipeline
	err = redisPipe.ExecPipeline()
	if err != nil {
		return nil, err
	}
	return &models.User{Key: key, UserName: username}, nil
}


