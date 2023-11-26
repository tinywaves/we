package service

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"wehub/internal/domain"
	"wehub/internal/repository"
)

var (
	ErrorUserDuplicateEmail = repository.ErrorUserDuplicateEmail
)

type UsersService struct {
	repo *repository.UserRepository
}

func InitUsersService(repo *repository.UserRepository) *UsersService {
	return &UsersService{
		repo: repo,
	}
}

func (svc *UsersService) SignUp(ctx context.Context, u domain.User) error {
	password, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(password)
	return svc.repo.CreateNewUser(ctx, u)
}
