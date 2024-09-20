package field

import (
	"botx/pkg/validate"
	"fmt"
)

type ID struct {
	Value int64
}

func (id *ID) Build(value int64) error {
	id.Value = value
	return id.valid()
}

func (id *ID) BuildOmit(value int64) error {
	if validate.IsEmpty(value) {
		return nil
	}
	return id.Build(value)
}

func (id *ID) valid() error {
	if !validate.IsPositive(id.Value) {
		return fmt.Errorf(validate.IsPositiveMsg, id.Name())
	}
	return nil
}

func (id *ID) Name() string {
	return "ID"
}
