package user

import (
	"botx/pkg/crypt"
	"botx/pkg/validate"
	"errors"
	"fmt"
)

type Password struct {
	Value     string
	ReValue   string
	HashValue string
}

func (password *Password) Build(value string, reValue string) error {
	password.Value, password.ReValue = value, reValue
	if err := password.valid(); err != nil {
		return err
	}
	password.HashValue = crypt.EncodePassword(password.Value)
	return nil
}

func (password *Password) valid() error {
	if !validate.IsEqual(password.Value, password.ReValue) {
		return errors.New("两次密码不一致")
	}

	if !validate.MinLen(password.Value, password.MinLen()) {
		return fmt.Errorf(validate.MinLenMsg, password.Name(), password.MinLen())
	}

	if !validate.MaxLen(password.Value, password.MaxLen()) {
		return fmt.Errorf(validate.MaxLenMsg, password.Name(), password.MaxLen())
	}
	return nil
}

func (password *Password) MinLen() int {
	return 6
}

func (password *Password) MaxLen() int {
	return 20
}

func (password *Password) Name() string {
	return "密码"
}
