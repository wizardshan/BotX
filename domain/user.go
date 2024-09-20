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

func (user *User) Mapper() *response.User {
	if user == nil {
		return nil
	}
	resp := new(response.User)
	resp.HashID = user.HashID.Value
	resp.Mobile = user.Mobile.Value
	resp.Age = user.Age.Value
	resp.Level = user.Level.Value
	resp.LevelDesc = user.Level.Desc
	return resp
}
