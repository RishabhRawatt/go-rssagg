.PHONY: help build run migrate-up migrate-down sqlc

help:
	@echo "Available commands:"
	@echo "  make build       - Build the application"
	@echo "  make run         - Run the application"
	@echo "  make migrate-up  - Run database migrations"
	@echo "  make migrate-down- Rollback database migrations"
	@echo "  make sqlc        - Generate database code with sqlc"

build:
	go build -o go-rssagg

run:
	go run .

migrate-up:
	goose -dir sql/schema postgres "${DB_URL}" up

migrate-down:
	goose -dir sql/schema postgres "${DB_URL}" down

sqlc:
	sqlc generate
