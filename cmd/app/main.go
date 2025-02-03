package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
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

	srv := server.Server{}

	go func() {
		if err := srv.Run(handler, config); err != nil {
			logrus.Fatal(err)
		}
	}()

	logrus.Info("server is running")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Info("server is shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}

}
