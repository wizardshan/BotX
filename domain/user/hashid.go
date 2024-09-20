package user

import (
	"botx/domain/field"
	"botx/pkg/idx"
	"botx/pkg/validate"
)

var HashIDCategory = field.HashIDCategoryUser

type HashID struct {
	Value string
	ID    int64
}

func (hashID *HashID) New(value string) (err error) {
	hashID.Value = value
	return hashID.Decode()
}

func (hashID *HashID) NewByID(id int64) {
	hashID.ID = id
	hashID.Encode()
}

func (hashID *HashID) Valid() error {
	return nil
}

func (hashID *HashID) ValidOmit() error {
	if validate.IsEmpty(hashID.Value) {
		return nil
	}
	return hashID.Valid()
}

func (hashID *HashID) Encode() {
	hashID.Value = idx.Encode(hashID.ID, HashIDCategory)
}

func (hashID *HashID) Decode() (err error) {
	hashID.ID, err = idx.Decode(hashID.Value, HashIDCategory)
	return
}
