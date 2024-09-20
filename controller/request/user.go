package request

import (
	"botx/domain"
)

type User struct {
	HashID string
}

func (reqUser *User) Mapper() (*domain.User, error) {
	domUser := new(domain.User)
	domUser.HashID.Value = reqUser.HashID
	var err error
	if err = domUser.HashID.Decode(); err != nil {
		return nil, err
	}
	domUser.ID.Value = domUser.HashID.ID
	return domUser, nil
}
