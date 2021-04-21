package data

import "time"

type Session struct {
	Id       int
	Uuid     string
	Email    string
	UserId   int
	CreateAt time.Time
}

func (session *Session) Check() (result bool, err error) {
	return
}
