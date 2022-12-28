package router

import (
	"github.com/frozen599/s3-assignment/api/internal/handler"
	"github.com/go-chi/chi/v5"
)

func SubscriberRouter() chi.Router {
	r := chi.NewRouter()
	subscriberHandler := handler.NewSubscriberHandler()
	r.Get("/", subscriberHandler.CreateSubscription)
	return r
}
