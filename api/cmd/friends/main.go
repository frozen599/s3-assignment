package main

import (
	"log"
	"net/http"

	"github.com/frozen599/s3-assignment/api/internal/api/router"
	"github.com/frozen599/s3-assignment/api/internal/config"
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
	r.Use(middleware.Recoverer)

	r.Mount("/", router.HealthCheckRouter())
	r.Mount("/api/v1/friends", router.FriendRouter())

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
