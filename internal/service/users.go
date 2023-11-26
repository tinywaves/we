package service

import (
	"context"
	"wehub/internal/domain"
	"wehub/internal/repository"
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
	return svc.repo.CreateNewUser(ctx, u)
}
