package main

import (
	"github.com/joho/godotenv"
	"log"
	"testProject/configs"
	"testProject/internal/app"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading env variables: %s", err)
	}

	cfg, err := configs.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
