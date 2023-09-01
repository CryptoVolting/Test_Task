package main

import (
	"log"
	"testProject/configs"
	"testProject/internal/app"
)

func main() {
	cfg, err := configs.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
