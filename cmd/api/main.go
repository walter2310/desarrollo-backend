package main

import (
	"log"

	"github.com/walter2310/desarrollo-backend/internal/config"
)

func main() {
	envs := config.LoadEnvs()

	app := &Application{
		Config: Config{
			Address: envs.Port,
		},
	}

	mux := app.mount()
	log.Fatal(app.start(mux))
}
