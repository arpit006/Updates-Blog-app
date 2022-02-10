package services

import (
	c "arpit006/web_app_with_go/constants"
	"arpit006/web_app_with_go/datastore"
	"arpit006/web_app_with_go/error_handler"
	"arpit006/web_app_with_go/models"
	"arpit006/web_app_with_go/models/vo"
	"arpit006/web_app_with_go/sessions"
	"arpit006/web_app_with_go/templ"
	"arpit006/web_app_with_go/util"
	"fmt"
	"log"
	"net/http"
	"time"
)




func UpdatesGetHandler(w http.ResponseWriter, r *http.Request) {
	//sessions2.ValidateSession(w, r)
	templates := templ.GetTemplateFactory()
	updatesIds, err := datastore.GetRangeFromRedis(c.UPDATES, 0, 100)
	if err != nil {
		panic(err)
	}
	updatesVos := make([]*vo.UpdatesVo, len(updatesIds))
	for i, id := range updatesIds {
		updateVo := CreateUpdatesVos(id)
		updatesVos[i] = updateVo
	}
	templates.ExecuteTemplate(w, "index.html", updatesVos)
}

func UpdatesPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	body := r.PostForm.Get(c.UPDATES)
	userId, err := sessions.GetUserIdFromSession(w, r)
	if err != nil {
		log.Printf("No userId exists in the session")
		error_handler.HandleAuthError(w, r)
	}
	err = PostUpdate(userId, body)
	http.Redirect(w, r, "/", 302)
}

func PostUpdate(userId int64, body string) error {
	update, err := CreateUpdates(userId, body)
	if err != nil {
		log.Printf("Error in posting update for user: [%d]! Error is %s", userId, err)
		return err
	}
	log.Printf("Update [%+v\n] posted successfully! ", update)
	return nil
}

func CreateUpdates(userId int64, body string) (*models.Updates, error){
	id, err := datastore.Incr(c.UPDATES_NEXT_ID)
	if err != nil {
		log.Printf("Error in getting next INCR key for Redis. Error is %s", err)
		return nil, err
	}
	time := time.Now().Format("2006-01-02 15:04:05.000")
	hashKey := fmt.Sprintf(c.UPDATE_BY_ID, id)
	redisPipe := datastore.StartPipeline()
	redisPipe.ExecInPipeStrInt(hashKey, c.ID, id)
	redisPipe.ExecInPipeStrInt(hashKey, c.USER_ID, userId)
	redisPipe.ExecInPipeStrStr(hashKey, c.BODY, body)
	redisPipe.ExecInPipeStrStr(hashKey, c.TIME, time)
	redisPipe.LPushInPipeStrInt(c.UPDATES, id)
	err = redisPipe.ExecPipeline()
	if err != nil {
		log.Printf("Error in running Redis Pipeline for user-updates Id: [%d:%d]. Error is %s", userId, id, err)
		return nil, err
	}
	return models.NewUpdates(hashKey, userId, id), nil
}

func GetUserIdForPost(hashKey string, postId int64) (int64, error) {
	userId, err := datastore.HGetStrInt(hashKey, c.USER_ID)
	if err != nil {
		log.Printf("No userId exists for this post: [%d]. Error is %s", postId, err)
		return -1, err
	}
	return userId, nil
}

func GetPostBodyForPost(hashKey string, postId int64) (string, error) {
	postBody, err := datastore.HGetStrStr(hashKey, c.BODY)
	if err != nil {
		log.Printf("No post exists against postId :[%d] . Error is %s", postId, err)
		return "", nil
	}
	return postBody, nil
}

func GetPostCreatedTime(hashKey string, postId int64) (time.Time, error) {
	postTimeStr, err := datastore.HGetStrStr(hashKey, c.TIME)
	if err != nil {
		log.Printf("No time for post: [%d] exists. Error is %s", postId, err)
		return time.Time{}, err
	}
	postTime, err := util.ParseStringToDateTime(postTimeStr)
	if err != nil {
		return time.Time{}, err
	}
	return postTime, nil
}

