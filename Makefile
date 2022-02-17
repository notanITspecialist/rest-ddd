env:
	cp .env.example .env

export_env:
	export $(cat .env)

run:
	go run cmd/rest-ddd/main.go

create_migrations:
	migrate create -ext sql -dir migrations/postgres -seq $(name)

migrate_up:
	migrate -database "${POSTGRESQL_DSN_MIGRATION}" -path migrations/postgres up $(count)

migrate_down:
	migrate -database "${POSTGRESQL_DSN_MIGRATION}" -path migrations/postgres down $(count)

update_mocks:
	mockgen -source internal/endpoints/user.go -destination internal/mocks/endpoints/user.go
	mockgen -source internal/service/user.go -destination internal/mocks/service/user.go
	mockgen -source internal/repository/user.go -destination internal/mocks/repository/user.go