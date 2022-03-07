.PHONY:
.SILENT:
.DEFAULT_GOAL := run
include .env
export $(shell sed 's/=.*//' .env)


build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/app/main.go
run:
	go run ./cmd/app/main.go
swag:
	swag init --parseDependency --parseInternal --parseDepth 1 -g internal/app/app.go
migrate:
	migrate -path db/migrations -database "postgresql://$$DB_USER:$$DB_PASSWORD@localhost:5432/$$DB_NAME?sslmode=disable" -verbose up
migrate_force:
	migrate -path db/migrations -database "postgresql://$$DB_USER:$$DB_PASSWORD@localhost:5432/$$DB_NAME?sslmode=disable" force $(v)
gen:
	mockgen -source=internal/services/services.go -destination=internal/services/mocks/mock.go
	mockgen -source=internal/repositories/repositories.go -destination=internal/repositories/mocks/mock.go