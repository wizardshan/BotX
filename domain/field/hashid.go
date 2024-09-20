package field

import (
	"botx/pkg/idx"
	"botx/pkg/validate"
	"fmt"
)

type HashID struct {
	Value    string
	ID       int64
	category int64
}

func (hashID *HashID) BuildByValueAndCategory(value string, category int64) error {
	hashID.Value = value
	hashID.category = category
	if err := hashID.valid(); err != nil {
		return err
	}
	return hashID.decode()
}

func (hashID *HashID) BuildByIDAndCategory(id int64, category int64) {
	hashID.ID = id
	hashID.category = category
	hashID.encode()
}

func (hashID *HashID) encode() {
	hashID.Value = idx.Encode(hashID.ID, hashID.category)
}

func (hashID *HashID) decode() (err error) {
	hashID.ID, err = idx.Decode(hashID.Value, hashID.category)
	return
}

func (hashID *HashID) valid() error {
	if !validate.IsAlphanumeric(hashID.Value) {
		return fmt.Errorf(validate.IsAlphanumericMsg, hashID.Name())
	}
	return nil
}

func (hashID *HashID) Name() string {
	return "HashID"
}
