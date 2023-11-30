package main

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"

	"gitlab.com/sunrise-dev/utest/utest-auth-app/config"
	"gitlab.com/sunrise-dev/utest/utest-auth-app/internal/app"
)

func main() {
	// Configuration
	var cfg config.Config

	err := cleanenv.ReadConfig("./config/config.yml", &cfg)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(&cfg)
}
