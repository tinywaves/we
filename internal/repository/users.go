package repository

import (
	"context"
	"wehub/internal/domain"
	"wehub/internal/repository/dao"
)

var (
	ErrorUserDuplicateEmail = dao.ErrorUserDuplicateEmail
)

type UserRepository struct {
	dao *dao.UserDao
}

func InitUsersRepository(dao *dao.UserDao) *UserRepository {
	return &UserRepository{
		dao: dao,
	}
}

func (repo *UserRepository) CreateNewUser(ctx context.Context, u domain.User) error {
	return repo.dao.InsertNewUser(ctx, dao.User{
		Email:    u.Email,
		Password: u.Password,
	})
}
