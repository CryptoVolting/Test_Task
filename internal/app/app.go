package app

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"testProject/configs"
	"testProject/internal/controller/http/v1"
	"testProject/internal/usecase"
	"testProject/internal/usecase/repository"
	"testProject/pkg"
)

func Run(cfg *configs.Config) {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	db, err := pkg.NewPostgresDB(pkg.Config{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		Username: cfg.DB.UserName,
		DBName:   cfg.DB.DbName,
		SSLMode:  cfg.DB.SslMode,
		Password: cfg.DB.Password,
	})
	if err != nil {
		logrus.Fatalln("failed to initialize db:", err.Error())
	}
	redisClient := redis.NewClient(&redis.Options{
		Addr: cfg.Redis.Port,
	})

	repositoryAtWork := repository.NewRepository(db)
	cacheService := repository.NewRedisRepository(db)
	usecases := usecase.NewUsecase(repositoryAtWork, cacheService, redisClient)
	handlers := v1.NewHandler(usecases)

	srv := new(pkg.Server)

	go func() {
		if err := srv.Run(cfg.HTTP.Port, handlers.InitRoutes()); err != nil && err != http.ErrServerClosed {
			logrus.Fatalln("error occurred while running http server:", err)
		}
	}()

	logrus.Println("TestTask Started")

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Println("TestTask Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}

	if err := redisClient.Close(); err != nil {
		logrus.Errorf("error occured on redisClient connection close: %s", err.Error())
	}
}
