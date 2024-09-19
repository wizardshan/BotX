package repository

import (
	"botx/domain"
	"botx/repository/ent"
	"botx/repository/ent/user"
	"context"
)

type User struct {
	db *ent.Client
}

func NewUser(db *ent.Client) *User {
	repo := new(User)
	repo.db = db
	return repo
}

func (repo *User) FetchByID(ctx context.Context, id int64) *domain.User {
	return repo.db.User.Query().Where(user.ID(id)).FirstX(ctx).Mapper()

}
