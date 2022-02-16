package repository

import "context"

type (
	UserRepository interface {
		GetAllUsers(context.Context) ([]User, error)
	}

	User struct {
		Id        string `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}
)
