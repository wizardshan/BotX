package request

import (
	"botx/domain/field"
)

type SmsCaptcha struct {
	Mobile string
}

func (req *SmsCaptcha) Mapper() (*field.Mobile, error) {
	domMobile := new(field.Mobile)
	if err := domMobile.Build(req.Mobile); err != nil {
		return nil, err
	}
	return domMobile, nil
}
