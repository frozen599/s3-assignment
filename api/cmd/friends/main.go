package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/frozen599/s3-assignment/api/internal/api/router"
	"github.com/frozen599/s3-assignment/api/internal/config"
	"github.com/frozen599/s3-assignment/api/internal/repo"
)

func main() {
	cfg := config.NewConfig("./../../..")
	log.Println(cfg)
	db := config.InitDB(cfg)
	log.Println(db == nil)
	if db == nil {
		log.Fatal("cannot establish connection to db")
	}
	defer db.Close()
	fmt.Println(cfg)

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
