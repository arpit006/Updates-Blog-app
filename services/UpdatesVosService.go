package services

import (
	c "arpit006/web_app_with_go/constants"
	"arpit006/web_app_with_go/models/vo"
	"fmt"
	"log"
	"strconv"
)

func CreateUpdatesVos(id string) *vo.UpdatesVo {
	postId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Printf("PostID: is not integer, Error is %s", err)
		log.Print("POST-ID :: =>  " +  id)
		return nil
	}
	updatesHashKey := fmt.Sprintf(c.UPDATE_BY_ID, postId)
	userId, err := GetUserIdForPost(updatesHashKey, postId)
	postBody, err := GetPostBodyForPost(updatesHashKey, postId)
	username, err := GetUserNameFromUserId(userId)
	postTime, err := GetPostCreatedTime(updatesHashKey, postId)

	return vo.NewUpdatesVo(username, postBody, postTime)
}
