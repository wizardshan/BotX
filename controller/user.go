package controller

import (
	"botx/controller/request"
	"botx/controller/response"
	"botx/repository"
	"fmt"
	"github.com/gin-gonic/gin"
)

type User struct {
	repo *repository.User
}

func NewUser(repo *repository.User) *User {
	ctr := new(User)
	ctr.repo = repo
	return ctr
}

func (ctr *User) Show(c *gin.Context) (response.Data, error) {
	req := request.User{HashID: c.Param("hashID")}
	domUser, err := req.Mapper()
	if err != nil {
		return nil, err
	}

	domUser = ctr.repo.FetchByID(c, domUser.HashID.ID)
	return domUser.Mapper(), nil
}

func (ctr *User) One(c *gin.Context) (response.Data, error) {
	req := new(request.User)
	if err := c.ShouldBind(req); err != nil {
		return nil, err
	}

	domUser, err := req.Mapper()
	if err != nil {
		return nil, err
	}

	domUser = ctr.repo.FetchByID(c, domUser.HashID.ID)
	return domUser.Mapper(), nil
}

func (ctr *User) Login(c *gin.Context) (response.Data, error) {
	req := new(request.UserLogin)
	if err := c.ShouldBind(req); err != nil {
		return nil, err
	}

	domUser, domCaptcha, err := req.Mapper()
	if err != nil {
		return nil, err
	}
	fmt.Println(domUser, domCaptcha)
	return nil, nil
}

func (ctr *User) Register(c *gin.Context) (response.Data, error) {
	req := new(request.UserRegister)
	if err := c.ShouldBind(req); err != nil {
		return nil, err
	}

	domUser, err := req.Mapper()
	if err != nil {
		return nil, err
	}

	domUser = ctr.repo.FetchByID(c, domUser.HashID.ID)
	return domUser.Mapper(), nil
}
