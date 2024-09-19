package field

import "regexp"

type MobileField struct {
	Mobile string `binding:"required,number,mobile"`
}

func (f *MobileField) ValidateMobile() bool {
	matched, _ := regexp.MatchString(`^(1[3-9][0-9]\d{8})$`, f.Mobile)
	return matched
}
