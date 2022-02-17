package repository

import "context"

type (
	UserRepository interface {
		GetAllUsers(ctx context.Context) ([]User, error)
	}

	User struct {
		Id        string `json:"id" db:"id"`
		FirstName string `json:"first_name" db:"first_name"`
		LastName  string `json:"last_name"  db:"last_name"`
	}
)
