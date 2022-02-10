package vo

import "time"

type UpdatesVo struct {
	Username string
	Body string
	Time time.Time
}

func NewUpdatesVo(username, body string, createdAt time.Time) *UpdatesVo {
	return &UpdatesVo{username, body, createdAt}
}

func (u *UpdatesVo) GetBody() string {
	return u.Body
}

func (u *UpdatesVo) GetCreatedAtTime() time.Time {
	return u.Time
}

func (u *UpdatesVo) GetUsername() string {
	return u.Username
}




