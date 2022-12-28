package router

import (
	"github.com/frozen599/s3-assignment/api/internal/handler"
	"github.com/go-chi/chi/v5"
)

func FriendRouter() chi.Router {
	r := chi.NewRouter()

	r.Post("/", hanlder.CreateFriendConnection)
	r.Post("/list", hanlder.GetFriendList)

	return r
}
