package db

import (
	"fmt"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"

	"rest-ddd/internal/config"
)

type PostgresqlClient struct {
	DB     *sqlx.DB
	Config config.Postgresql
}

func NewPostgresClient(cfg config.Postgresql) (*PostgresqlClient, error) {
	client, err := sqlx.Connect("pgx", cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("error when try to connect to postgresql: %w", err)
	}
	err = client.Ping()
	if err != nil {
		return nil, fmt.Errorf("error when try to ping to postgresql: %w", err)
	}
	return &PostgresqlClient{
		DB:     client,
		Config: cfg,
	}, nil
}

func (client *PostgresqlClient) Close() error {
	return client.DB.Close()
}
