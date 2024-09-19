package response

import (
	"botx/domain/field"
	"botx/domain/user"
)

type User struct {
	user.HashIDField
	field.MobileField
}
