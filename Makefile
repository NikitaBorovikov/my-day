.PHONY:

docker-build:
	docker build -t my-day-app .
run:
	docker-compose up 
swag:
	swag init -g cmd/app/main.go