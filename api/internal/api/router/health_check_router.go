package router

import (
	"github.com/frozen599/s3-assignment/api/internal/controller"
	"github.com/go-chi/chi/v5"
)

func HealthCheckRouter() chi.Router {
	r := chi.NewRouter()
	r.Get("/", controller.HealthCheck)
	return r
}
