package domain

import (
	"botx/domain/field"
	"botx/domain/user"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var validate = binding.Validator.Engine().(*validator.Validate)

func init() {
	validate.RegisterValidation("mobile", func(fl validator.FieldLevel) bool {
		mobile := new(field.MobileField)
		mobile.Mobile = fl.Field().String()
		return mobile.ValidateMobile()
	})
	validate.RegisterValidation("userHashID", func(fl validator.FieldLevel) bool {
		hashID := new(user.HashIDField)
		hashID.HashID = fl.Field().String()
		if _, err := hashID.DecodeID(); err != nil {
			return false
		}
		return true
	})
}
