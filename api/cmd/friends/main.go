package main

import (
	"log"
	"net/http"

	"github.com/frozen599/s3-assignment/api/internal/config"
	"github.com/frozen599/s3-assignment/api/internal/controller"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	cfg := config.NewConfig()
	db := config.InitDB(cfg)
	if db == nil {
		panic("cannot establish connection to db")
	}

	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(cfg.ReadTimeout))

	r.Get("/", controller.HealthCheck)

	server := &http.Server{
		Handler:      r,
		Addr:         cfg.Addr,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
