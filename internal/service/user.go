package service

import (
	"context"
	"rest-ddd/internal/repository"
)

type (
	UserService interface {
		GetAllUsers(ctx context.Context) ([]repository.User, error)
		CreateUser(ctx context.Context, body UserCreateData) error
	}

	UserCreateData struct {
		FirstName string `json:"first_name" validate:"min=2"`
		LastName  string `json:"last_name"  validate:"min=2"`
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

func (h *userService) CreateUser(ctx context.Context, body UserCreateData) error {
	userCreateData := repository.User{
		FirstName: body.FirstName,
		LastName:  body.LastName,
	}

	err := h.repo.CreateUser(ctx, userCreateData)
	if err != nil {
		return err
	}

	return nil
}
