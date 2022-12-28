package router

import (
	"github.com/frozen599/s3-assignment/api/internal/handler"
	"github.com/go-chi/chi/v5"
)

func HealthCheckRouter() chi.Router {
	r := chi.NewRouter()
	r.Get("/", hanlder.HealthCheck)
	return r
}
