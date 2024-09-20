package field

import (
	"botx/pkg/validate"
	"fmt"
)

type ID struct {
	Value int64
}

func (id *ID) Valid() error {
	if !validate.IsPositive(id.Value) {
		return fmt.Errorf(validate.IsPositiveMsg, id.Name())
	}
	return nil
}

func (id *ID) ValidOmit() error {
	if validate.IsEmpty(id.Value) {
		return nil
	}
	return id.Valid()
}

func (id *ID) Name() string {
	return "ID"
}
