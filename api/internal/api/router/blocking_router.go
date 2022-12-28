package router

import (
	"github.com/frozen599/s3-assignment/api/internal/handler"
	"github.com/go-chi/chi/v5"
)

func NewBlockingRouter() chi.Router {
	r := chi.NewRouter()
	blockingHandler := handler.NewBlockingHandler()
	r.Post("/", blockingHandler.BlockUpdate)
	return r
}
