package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Application struct {
	Config Config
}

type Config struct {
	Address string
}

func (app *Application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.HealthCheckHandler)
	})

	return r
}

func (app *Application) start(mux http.Handler) error {
	server := &http.Server{
		Addr:         app.Config.Address,
		Handler:      mux,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 8,
		IdleTimeout:  time.Minute,
	}

	log.Printf("server running on port: %s", server.Addr)

	return server.ListenAndServe()
}
