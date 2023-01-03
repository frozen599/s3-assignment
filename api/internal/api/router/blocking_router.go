package router

import (
	"github.com/frozen599/s3-assignment/api/internal/controller"
	"github.com/frozen599/s3-assignment/api/internal/handler"
	"github.com/frozen599/s3-assignment/api/internal/repo"
	"github.com/go-chi/chi/v5"
)

func BlockingRouter(userRepo repo.UserRepo, relaRepo repo.RelationshipRepo) chi.Router {
	r := chi.NewRouter()
	blockingController := controller.NewBlockingController(userRepo, relaRepo)
	blockingHandler := handler.NewBlockingHandler(blockingController)
	r.Post("/", blockingHandler.BlockUpdate)
	return r
}
