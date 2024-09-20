package request

import (
	"botx/domain"
	"botx/domain/field"
)

type User struct {
	HashID string
}

func (req *User) Mapper() (*domain.User, error) {
	domUser := new(domain.User)
	if err := domUser.HashID.Build(req.HashID); err != nil {
		return nil, err
	}
	return domUser, nil
}

type UserLogin struct {
	Mobile  string
	Captcha string
}

func (req *UserLogin) Mapper() (*domain.User, *field.Captcha, error) {
	domUser := new(domain.User)
	if err := domUser.Mobile.Build(req.Mobile); err != nil {
		return nil, nil, err
	}

	domCaptcha := new(field.Captcha)
	if err := domCaptcha.Build(req.Captcha); err != nil {
		return nil, nil, err
	}

	return domUser, domCaptcha, nil
}

type UserRegister struct {
	Mobile     string
	Captcha    string
	Nickname   string
	Password   string
	RePassword string
}

func (req *UserRegister) Mapper() (*domain.User, error) {
	domUser := new(domain.User)
	if err := domUser.Mobile.Build(req.Mobile); err != nil {
		return nil, err
	}
	return domUser, nil
}
