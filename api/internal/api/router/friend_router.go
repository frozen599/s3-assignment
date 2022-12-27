package router

import (
	"github.com/frozen599/s3-assignment/api/internal/controller"
	"github.com/go-chi/chi/v5"
)

func FriendRouter() chi.Router {
	r := chi.NewRouter()

	r.Post("/", controller.CreateFriendConnection)
	r.Post("/list", controller.GetFriendList)

	return r
}
