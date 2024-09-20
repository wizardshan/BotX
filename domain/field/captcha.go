package field

import (
	"botx/pkg/validate"
	"fmt"
)

type Captcha struct {
	Value string
}

func (captcha *Captcha) Build(value string) error {
	captcha.Value = value
	return captcha.valid()
}

func (captcha *Captcha) BuildOmit(value string) error {
	if validate.IsEmpty(value) {
		return nil
	}
	return captcha.Build(value)
}

func (captcha *Captcha) valid() error {
	if !validate.Len(captcha.Value, captcha.Len()) {
		return fmt.Errorf(validate.LenMsg, captcha.Name(), captcha.Len())
	}
	return nil
}

func (captcha *Captcha) Len() int {
	return 4
}

func (captcha *Captcha) Name() string {
	return "验证码"
}
