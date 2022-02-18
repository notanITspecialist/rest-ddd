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
	users := []User{}
	query := fmt.Sprintf(`SELECT users.id, first_name, last_name, mobile FROM %v LEFT OUTER JOIN profiles ON (profiles.userId = users.id)`, r.table)
	err := r.client.DB.SelectContext(ctx, &users, query)
	if err != nil {
		r.log.Error(fmt.Sprintf("error while select from %v table", r.table), zap.Error(err))
	}

	return users, err
}

func (r *pgUserRepository) CreateUser(ctx context.Context, data User) error {
	// TODO: it's better to do it through a transaction

	query := `INSERT INTO users("first_name", "last_name") VALUES ($1, $2) RETURNING id`
	stmt, err := r.client.DB.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	var id int
	err = stmt.QueryRowContext(ctx, data.FirstName, data.LastName).Scan(&id)
	if err != nil {
		return err
	}

	query = fmt.Sprintf(`INSERT INTO profiles(userId, mobile) VALUES (%v, %v)`, id, data.Mobile)
	_, err = r.client.DB.QueryxContext(ctx, query)
	if err != nil {
		return err
	}

	return nil
}
