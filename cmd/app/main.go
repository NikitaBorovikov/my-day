package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"toDoApp/pkg/config"
	"toDoApp/pkg/db"
	"toDoApp/pkg/handlers"
	"toDoApp/pkg/repository"
	"toDoApp/pkg/server"
	"toDoApp/pkg/usecases"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

// @title MyDay App API
// @version 1.0
// @description REST API for planning your day

// @host localhost:2323
// @BasePath /

// @securityDefinitions.apikey sessionKey
// @in header
// @name Authorization
func main() {

	config, err := config.Init()
	if err != nil {
		logrus.Fatal(err)
	}

	db, err := db.Connect(config)
	if err != nil {
		logrus.Fatal(err)
	}

	repository := repository.InitRepository(db)
	usecases := usecases.InitUseCases(repository)
	handlers := handlers.InitHandlers(usecases)

	srv := server.Server{}

	go func() {
		if err := srv.Run(handlers, config); err != nil {
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
