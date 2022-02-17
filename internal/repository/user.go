package repository

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"rest-ddd/internal/db"
)

type (
	UserRepository interface {
		GetAllUsers(ctx context.Context) ([]User, error)
		CreateUser(ctx context.Context, data User) error
	}

	pgUserRepository struct {
		log *zap.Logger

		client *db.PostgresqlClient
		table  string
	}
)

func NewPGUserRepository(log *zap.Logger, client *db.PostgresqlClient) UserRepository {
	return newPGUserRepository(log, client)
}

func newPGUserRepository(log *zap.Logger, client *db.PostgresqlClient) *pgUserRepository {
	return &pgUserRepository{
		log:    log,
		client: client,
		table:  "users",
	}
}

func (r *pgUserRepository) GetAllUsers(ctx context.Context) ([]User, error) {
	var users []User
	query := fmt.Sprintf(`SELECT * FROM %v`, r.table)
	err := r.client.DB.SelectContext(ctx, &users, query)
	if err != nil {
		r.log.Error(fmt.Sprintf("error while select from %v table", r.table), zap.Error(err))
	}

	return users, err
}

func (r *pgUserRepository) CreateUser(ctx context.Context, data User) error {
	query := fmt.Sprintf(`INSERT INTO %v(first_name, last_name) VALUES (%v, %v)`, r.table, data.FirstName, data.LastName)
	_, err := r.client.DB.QueryxContext(ctx, query)
	if err != nil {
		r.log.Error(fmt.Sprintf("error while insert into %v table", r.table), zap.Error(err))
	}

	return err
}
