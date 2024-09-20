package field

import (
	"botx/pkg/validate"
	"fmt"
	"regexp"
)

type Mobile struct {
	Value string
}

func (mobile *Mobile) Build(value string) error {
	mobile.Value = value
	return mobile.valid()
}

func (mobile *Mobile) BuildOmit(value string) error {
	if validate.IsEmpty(value) {
		return nil
	}
	return mobile.Build(value)
}

func (mobile *Mobile) valid() error {
	if matched := regexp.MustCompile(`^(1[3-9][0-9]\d{8})$`).MatchString(mobile.Value); !matched {
		return fmt.Errorf(validate.FormatMsg, mobile.Name())
	}
	return nil
}

func (mobile *Mobile) Name() string {
	return "手机号码"
}
