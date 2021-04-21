package data

import "time"

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

func UserByEmail(email string) (user User, err error) {
	return
}

func (user *User) CreateSession() (session Session) {
	return
}
