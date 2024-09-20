package user

import (
	"botx/domain/field"
	"botx/pkg/validate"
)

type HashID struct {
	field.HashID
}

func (hashID *HashID) BuildByID(id int64) {
	hashID.BuildByIDAndCategory(id, hashID.Category())
}

func (hashID *HashID) Build(value string) (err error) {
	return hashID.BuildByValueAndCategory(value, hashID.Category())
}

func (hashID *HashID) BuildOmit(value string) error {
	if validate.IsEmpty(value) {
		return nil
	}
	return hashID.Build(value)
}

func (hashID *HashID) Category() int64 {
	return field.HashIDCategoryUser
}
