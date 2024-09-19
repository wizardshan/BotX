package ent

import (
	"botx/domain"
)

func (entUser *User) Mapper() *domain.User {
	if entUser == nil {
		return nil
	}
	domUser := new(domain.User)
	domUser.ID = entUser.ID
	domUser.HashID = entUser.HashID
	domUser.Mobile = entUser.Mobile
	return domUser
}
