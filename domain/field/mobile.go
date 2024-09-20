package field

import (
	"botx/pkg/validate"
	"fmt"
	"regexp"
)

type Mobile struct {
	Value string
}

func (mobile *Mobile) Valid() error {
	if matched, _ := regexp.MatchString(`^(1[3-9][0-9]\d{8})$`, mobile.Value); !matched {
		return fmt.Errorf(validate.FormatMsg, mobile.Name())
	}
	return nil
}

func (mobile *Mobile) ValidOmit() error {
	if validate.IsEmpty(mobile.Value) {
		return nil
	}
	return mobile.Valid()
}

func (mobile *Mobile) Name() string {
	return "手机号码"
}
