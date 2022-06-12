package main

import (
	"context"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/martynove/gophermart/internal/config"
	"github.com/martynove/gophermart/internal/delivery/rest"
	"github.com/martynove/gophermart/internal/delivery/rest/handler"
	"github.com/martynove/gophermart/internal/repository"
	"github.com/martynove/gophermart/internal/service"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// configure main logger
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	gin.SetMode(gin.ReleaseMode)
	cfg, err := config.NewConfig()

	if err != nil {
		log.Fatalf("error load new configuration: %s", err.Error())
	}
	// switch on debug mode
	if cfg.DebugMode {
		logger.SetLevel(logrus.DebugLevel)
		gin.SetMode(gin.DebugMode)
	}
	db, err := repository.NewPostgresDB(cfg.DatabaseURI)
	if err != nil {
		logger.Fatalf("error PostgresDB repository creation: %s", err.Error())
	}
	apiRepository := repository.NewRepository(logger, db)
	apiService := service.NewService(apiRepository, logger)
	apiHandler := handler.NewHandler(logger, apiService)
	apiServer := new(rest.Server)

	go func() {
		if err := apiServer.Run(cfg, apiHandler.InitRoutes()); err != nil {
			switch err.Error() {
			case "http: Server closed":
				logger.Infof("http server on the address: [%s] was down", cfg.ServerAddress)
			default:
				logger.Fatalf("error occured while running http server [%s] (%s)", cfg.ServerAddress, err.Error())
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := apiServer.Shutdown(context.Background()); err != nil {
		logger.Fatalf("error occured on server shutting down: %s", err.Error())
	}
}
