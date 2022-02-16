package service

import (
	"context"
	"rest-ddd/pkg/repository"
)

type (
	UserService interface {
		GetAllUsers(ctx context.Context) ([]repository.User, error)
	}

	userService struct{}
)

func NewUserService() UserService {
	return newUserService()
}

func newUserService() *userService {
	return &userService{}
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
