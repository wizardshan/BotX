package user

import (
	"botx/pkg/validate"
	"fmt"
)

type Age struct {
	Value int
}

func (age *Age) Build(value int) error {
	age.Value = value
	return age.valid()
}

func (age *Age) BuildOmit(value int) error {
	if validate.IsEmpty(value) {
		return nil
	}
	return age.Build(value)
}

func (age *Age) valid() error {
	if !validate.Min(age.Value, age.Min()) {
		return fmt.Errorf(validate.MinMsg, age.Name(), age.Min())
	}

	if !validate.Max(age.Value, age.Max()) {
		return fmt.Errorf(validate.MaxMsg, age.Name(), age.Max())
	}
	return nil
}

func (age *Age) Min() int {
	return 1
}

func (age *Age) Max() int {
	return 150
}

func (age *Age) Name() string {
	return "年龄"
}
