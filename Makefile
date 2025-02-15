include .env
export
.PHONY:

docker-build:
	docker build -t my-day-app .
run:
	docker-compose up 
migrate:
	migrate -path migrations -database 'postgres://${PG_USER}:${PG_PASSWORD}@0.0.0.0:${PG_PORT}/${PG_NAME}?sslmode=disable' up
swag:
	swag init -g cmd/app/main.go