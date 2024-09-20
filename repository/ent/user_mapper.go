package ent

import (
	"botx/domain"
)

func (entUser *User) Mapper() *domain.User {
	if entUser == nil {
		return nil
	}
	domUser := new(domain.User)
	domUser.ID.Value = entUser.ID
	domUser.HashID.Value = entUser.HashID
	domUser.Mobile.Value = entUser.Mobile
	domUser.Password.Value = entUser.Password
	domUser.Age.Value = entUser.Age
	domUser.Level.Build(entUser.Level)
	return domUser
}
