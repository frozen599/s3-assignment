package main

import (
	"log"
	"net/http"

	"github.com/frozen599/s3-assignment/api/internal/api/router"
	"github.com/frozen599/s3-assignment/api/internal/config"
	"github.com/frozen599/s3-assignment/api/internal/repo"
)

func main() {
	cfg := config.NewConfig()
	db := config.InitDB(cfg)
	if db == nil {
		panic("cannot establish connection to db")
	}
	defer db.Close()

	userRepo := repo.NewUserRepo(db)
	relaRepo := repo.NewRelationshipRepo(db)
	r := router.NewApiRouter(userRepo, relaRepo, *cfg)

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
