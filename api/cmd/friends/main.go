package main

import (
	"net/http"

	"github.com/frozen599/s3-assignment/api/internal/config"
	"github.com/frozen599/s3-assignment/api/internal/controller"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	cfg := config.NewConfig()

	r.Use(middleware.Logger)

	r.Use(middleware.Timeout(cfg.ReadTimeout))

	r.Get("/", controller.HealthCheck)

	http.ListenAndServe(cfg.Addr, r)
}
