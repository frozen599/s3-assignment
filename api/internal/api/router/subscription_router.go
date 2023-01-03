package router

import (
	"github.com/frozen599/s3-assignment/api/internal/controller"
	"github.com/frozen599/s3-assignment/api/internal/handler"
	"github.com/frozen599/s3-assignment/api/internal/repo"
	"github.com/go-chi/chi/v5"
)

func SubscriberRouter(userRepo repo.UserRepo, relaRepo repo.RelationshipRepo) chi.Router {
	r := chi.NewRouter()
	subscriberController := controller.NewSubscriberController(userRepo, relaRepo)
	subscriberHandler := handler.NewSubscriberHandler(subscriberController)
	r.Post("/", subscriberHandler.CreateSubscription)
	r.Post("/can-receive-update", subscriberHandler.CanReceiveUpdate)
	return r
}
