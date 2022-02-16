package repository

import (
	"context"

	"rest-ddd/internal/db"
)

type (
	pgUserRepository struct {
		client *db.PostgresqlClient
		table  string
	}
)

func NewPGUserRepository(client *db.PostgresqlClient) UserRepository {
	return newPGUserRepository(client)
}

func newPGUserRepository(client *db.PostgresqlClient) *pgUserRepository {
	return &pgUserRepository{
		client: client,
		table:  "user",
	}
}

func (r *pgUserRepository) GetAllUsers(context.Context) ([]User, error) {
	fixtureData := []User{
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
