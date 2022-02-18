# rest-ddd
[![Go version](https://img.shields.io/static/v1?label=Go&message=1.17.1&color=blue)](https://golang.org/)

### Environment variables

Name        | Default
----------- | ---------
SERVER_PORT | 8000         

### Migrations
Before start migrations, you should start postgres container `docker-compose up -d`

For migrate database use 
```bigquery
make env
make export_env
make migrate_up
```

### Start app
For start up use `make run`