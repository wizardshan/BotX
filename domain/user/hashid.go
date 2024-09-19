package user

import (
	"botx/domain/field"
	"botx/pkg/idx"
)

var HashIDCategory = field.HashIDUser

type HashIDField struct {
	HashID string `binding:"required,alphanum,userHashID"`
}

func (f *HashIDField) EncodeID(id int64) string {
	return idx.Encode(id, HashIDCategory)
}

func (f *HashIDField) DecodeID() (int64, error) {
	return idx.Decode(f.HashID, HashIDCategory)
}
