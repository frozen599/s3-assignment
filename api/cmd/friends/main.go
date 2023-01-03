package main

import (
	"log"
	"net/http"

	"github.com/frozen599/s3-assignment/api/internal/api/router"
	"github.com/frozen599/s3-assignment/api/internal/config"
	"github.com/frozen599/s3-assignment/api/internal/repo"
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
	defer db.Close()

	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(cfg.ReadTimeout))
	r.Use(middleware.Recoverer)

	userRepo := repo.NewUserRepo(db)
	relaRepo := repo.NewRelationshipRepo(db)

	r.Mount("/health_check", router.HealthCheckRouter())
	r.Mount("/api/v1/friends", router.FriendRouter(userRepo, relaRepo))
	r.Mount("/api/v1/blockings", router.BlockingRouter(userRepo, relaRepo))
	r.Mount("/api/v1/subscribers", router.SubscriberRouter(userRepo, relaRepo))

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
