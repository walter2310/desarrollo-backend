package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/walter2310/desarrollo-backend/internal/config"
	repository "github.com/walter2310/desarrollo-backend/internal/db/repository/models"
)

func main() {
	envs := config.LoadEnvs()
	ctx := context.Background()

	conn, err := pgxpool.New(ctx, envs.DbUrl)
	if err != nil {
		os.Exit(1)
	}

	defer conn.Close()

	app := &Application{
		Config: Config{
			Address: envs.Port,
		},
	}

	repo := repository.New(conn)
	users, _ := repo.GetAllUsers(ctx)

	fmt.Println(users)

	mux := app.mount()
	log.Fatal(app.start(mux))
}
