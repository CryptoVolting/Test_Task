package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"testProject"
	"testProject/pkg/handler"
	"testProject/pkg/repository"
	"testProject/pkg/service"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalln("error initializing configs:", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalln("error loading env variables:", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalln("failed to initialize db:", err.Error())
	}
	repos := repository.NewSRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(testProject.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalln("error occurred while running http server:", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
