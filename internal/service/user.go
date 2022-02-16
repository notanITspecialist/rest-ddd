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

	// There will be call a repository layer

	fixtureData := []repository.User{
		{
			Id:        "1",
			FirstName: "John",
			LastName:  "Doe",
		},
		{
			Id:        "2",
			FirstName: "John",
			LastName:  "Doe",
		},
		{
			Id:        "3",
			FirstName: "John",
			LastName:  "Doe",
		},
	}

	return fixtureData, nil
}
