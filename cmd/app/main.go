package main

import (
	"net/http"
	"toDoApp/pkg/config"
	"toDoApp/pkg/db"
	"toDoApp/pkg/repository/postgres"
	"toDoApp/pkg/server"
	"toDoApp/pkg/usecases"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {

	config, err := config.Init()
	if err != nil {
		logrus.Fatal(err)
	}

	db, err := db.Connect(config)
	if err != nil {
		logrus.Fatal(err)
	}
	defer db.Close()

	userRepo := postgres.NewUserRepositoty(db)
	taskRepo := postgres.NewTaskRepository(db)
	eventRepo := postgres.NewEventRepository(db)
	myDayRepo := postgres.NewMyDayRepository(db)

	userUseCase := usecases.NewUserUseCase(userRepo)
	taskUseCase := usecases.NewTaskUseCase(taskRepo)
	eventUseCase := usecases.NewEventUseCase(eventRepo)
	myDayUseCase := usecases.NewMyDayUseCase(myDayRepo)

	userHandler := server.NewUserHandler(userUseCase)
	taskHandler := server.NewTaskHandler(taskUseCase)
	eventHandler := server.NewEventHandler(eventUseCase)
	myDayHandler := server.NewMyDayHandler(myDayUseCase)

	handler := server.InitHandlers(userHandler, taskHandler, eventHandler, myDayHandler)

	server := server.Start(handler, config)
	if err := http.ListenAndServe(config.Http.Port, server); err != nil {
		logrus.Fatal(err)
	}
}
