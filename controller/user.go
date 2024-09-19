package controller

import (
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
	return ctr.repo.FetchByID(c, 1), nil
}
