package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
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
	query, _, err := sq.
		Select("id, first_name, last_name").
		From(r.table).
		ToSql()
	if err != nil {
		r.log.Error(fmt.Sprintf("error while format sql squirrel"), zap.Error(err))
	}
	users := []User{}
	err = r.client.DB.SelectContext(ctx, &users, query)
	if err != nil {
		r.log.Error(fmt.Sprintf("error while select from %v table", r.table), zap.Error(err))
	}

	return users, err
}

func (r *pgUserRepository) CreateUser(ctx context.Context, data User) error {
	// TODO: it's better to do it through a transaction

	query, args, err := sq.
		Insert(r.table).
		Columns("first_name", "last_name").
		Values(data.FirstName, data.LastName).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.client.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
