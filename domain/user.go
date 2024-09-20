package domain

import (
	"botx/controller/response"
	"botx/domain/field"
	"botx/domain/user"
)

type Users []*User

type User struct {
	ID       field.ID
	HashID   user.HashID
	Mobile   field.Mobile
	Password user.Password
	Age      user.Age
	Level    user.Level
}

func (domUser *User) Mapper() *response.User {
	respUser := new(response.User)
	respUser.HashID = domUser.HashID.Value
	respUser.Mobile = domUser.Mobile.Value
	respUser.Age = domUser.Age.Value
	respUser.Level = domUser.Level.Value
	respUser.LevelDesc = domUser.Level.Desc
	return respUser
}
