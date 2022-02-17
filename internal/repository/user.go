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
