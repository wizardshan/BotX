package user

import (
	"botx/domain/field"
)

var HashIDCategory = field.HashIDUser

type HashIDField struct {
	field.HashIDField
}

func (f *HashIDField) Encode() string {
	return f.En(HashIDCategory)
}

func (f *HashIDField) Decode() (int64, error) {
	return f.De(HashIDCategory)
}
