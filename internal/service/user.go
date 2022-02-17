package service

import (
	"context"
	"rest-ddd/internal/repository"
)

type (
	UserService interface {
		GetAllUsers(ctx context.Context) ([]repository.User, error)
	}

	userService struct {
		repo repository.UserRepository
	}
)

func NewUserService(repo repository.UserRepository) UserService {
	return newUserService(repo)
}

func newUserService(repo repository.UserRepository) *userService {
	return &userService{
		repo: repo,
	}
}

func (h *userService) GetAllUsers(ctx context.Context) ([]repository.User, error) {
	users, err := h.repo.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}
