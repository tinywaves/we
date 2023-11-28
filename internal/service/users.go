package service

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"wehub/internal/domain"
	"wehub/internal/repository"
)

var (
	ErrorUserDuplicateEmail     = repository.ErrorUserDuplicateEmail
	ErrorInvalidEmailOrPassword = repository.ErrorInvalidEmailOrPassword
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

func (svc *UsersService) SignIn(ctx context.Context, u domain.User) error {
	user, err := svc.repo.FindUserByEmail(ctx, u.Email)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
	if err != nil {
		return ErrorInvalidEmailOrPassword
	}
	return nil
}
