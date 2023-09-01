package app

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"testProject/configs"
	"testProject/internal/handler"
	repository2 "testProject/internal/repository"
	"testProject/internal/usecase"
	"testProject/pkg"
)

func Run(cfg *configs.Config) {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := godotenv.Load(); err != nil {
		logrus.Fatalln("error loading env variables:", err.Error())
	}

	db, err := pkg.NewPostgresDB(pkg.Config{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		Username: cfg.DB.UserName,
		DBName:   cfg.DB.DbName,
		SSLMode:  cfg.DB.SslMode,
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalln("failed to initialize db:", err.Error())
	}
	repos := repository2.NewSRepository(db)
	usecases := usecase.NewUsecase(repos)
	handlers := handler.NewHandler(usecases)

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
}
