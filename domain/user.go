package domain

import (
	"botx/controller/response"
	"botx/domain/field"
	"botx/domain/user"
)

type Users []*User

type User struct {
	field.IDField
	user.HashIDField
	field.MobileField
}

func (dom *User) Mapper() *response.User {
	resp := new(response.User)
	resp.HashID = dom.HashID
	resp.Mobile = dom.Mobile
	return resp
}
