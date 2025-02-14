.PHONY:

docker-build:
	docker build -t my-day-app .
run:
	docker-compose up 
migrate:
	migrate -path migrations -database 'postgres://postgres:23112005@0.0.0.0:5432/myDayDB?sslmode=disable' up
swag:
	swag init -g cmd/app/main.go