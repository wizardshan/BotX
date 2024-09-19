package domain

import (
	"botx/domain/field"
	"botx/domain/user"
)

type Users []*User

type User struct {
	field.IDField
	user.HashIDField
}
