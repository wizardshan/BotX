package controller

import (
	"botx/controller/request"
	"botx/controller/response"
	"botx/repository"
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

func (ctr *User) One(c *gin.Context) (response.Data, error) {
	reqUser := new(request.User)
	if err := c.ShouldBind(reqUser); err != nil {
		return nil, err
	}

	domUser, err := reqUser.Mapper()
	if err != nil {
		return nil, err
	}
	domUser = ctr.repo.FetchByID(c, domUser.ID.Value)
	return domUser.Mapper(), nil
}
