package main

import (
	"log"

	"github.com/walter2310/desarrollo-backend/internal/config"
	"github.com/walter2310/desarrollo-backend/internal/repository"
)

func main() {
	envs := config.LoadEnvs()
	store := repository.NewStorage(nil)

	app := &Application{
		Config: Config{
			Address: envs.Port,
			Store:   store,
		},
	}

	mux := app.mount()
	log.Fatal(app.start(mux))
}
