package request

import (
	"botx/domain"
	"botx/domain/user"
)

type User struct {
	user.HashIDField
}

func (req *User) Mapper() *domain.User {
	domUser := new(domain.User)
	domUser.HashID = req.HashID
	domUser.ID, _ = domUser.DecodeID()
	return domUser
}
