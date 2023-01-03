package router

import (
	"github.com/frozen599/s3-assignment/api/internal/controller"
	"github.com/frozen599/s3-assignment/api/internal/handler"
	"github.com/frozen599/s3-assignment/api/internal/repo"
	"github.com/go-chi/chi/v5"
)

func FriendRouter(userRepo repo.UserRepo, relaRepo repo.RelationshipRepo) chi.Router {
	r := chi.NewRouter()
	friendController := controller.NewFriendController(userRepo, relaRepo)
	friendHandler := handler.NewFriendHanlder(friendController)
	r.Post("/", friendHandler.CreateFriendConnection)
	r.Post("/list", friendHandler.GetFriendList)

	return r
}
