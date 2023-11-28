package repository

import (
	"context"
	"wehub/internal/domain"
	"wehub/internal/repository/dao"
)

var (
	ErrorUserDuplicateEmail     = dao.ErrorUserDuplicateEmail
	ErrorInvalidEmailOrPassword = dao.ErrorInvalidEmailOrPassword
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

func (repo *UserRepository) FindUserByEmail(ctx context.Context, email string) (domain.User, error) {
	user, err := repo.dao.FindUserByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		Email:    user.Email,
		Password: user.Password,
	}, nil
}
