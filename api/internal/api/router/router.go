package router

import (
	"github.com/frozen599/s3-assignment/api/internal/config"
	"github.com/frozen599/s3-assignment/api/internal/controller"
	"github.com/frozen599/s3-assignment/api/internal/handler"
	"github.com/frozen599/s3-assignment/api/internal/repo"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func NewApiRouter(userRepo repo.UserRepo, relaRepo repo.RelationshipRepo, cfg config.Config) chi.Router {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Timeout(cfg.ReadTimeout))
	router.Use(middleware.Recoverer)

	healthCheckHandler := handler.NewHealthCheckHandler()
	router.Route("/", func(r chi.Router) {
		r.Get("/health_check", healthCheckHandler.HealthCheck)
	})

	friendController := controller.NewFriendController(userRepo, relaRepo)
	friendHandler := handler.NewFriendHanlder(friendController)
	router.Route("/api/v1/friends", func(r chi.Router) {
		r.Post("/", friendHandler.CreateFriendConnection)
		r.Post("/friend-list", friendHandler.GetFriendList)
		r.Post("/mutual-friend-list", friendHandler.GetMutualFriendList)
	})

	blockingController := controller.NewBlockingController(userRepo, relaRepo)
	blockingHandler := handler.NewBlockingHandler(blockingController)
	router.Route("/api/v1/blockings", func(r chi.Router) {
		r.Post("/", blockingHandler.BlockUpdate)
	})

	subscriberController := controller.NewSubscriberController(userRepo, relaRepo)
	subscriberHandler := handler.NewSubscriberHandler(subscriberController)
	router.Route("/api/v1/subscribers", func(r chi.Router) {
		r.Post("/", subscriberHandler.CreateSubscription)
		r.Post("/can-receive-update", subscriberHandler.CanReceiveUpdate)
	})

	return router
}
