package models

type Updates struct {
	HashKey string
	UserId int64
	PostId int64
}

func NewUpdates(hashKey string, userId, postId int64) *Updates {
	return &Updates{hashKey, userId, postId}
}
