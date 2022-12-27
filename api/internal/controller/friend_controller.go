package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/frozen599/s3-assignment/api/internal/forms"
	"github.com/frozen599/s3-assignment/api/internal/pkg"
	"github.com/frozen599/s3-assignment/api/internal/repository"
)

func CreateFriendConnection(w http.ResponseWriter, r *http.Request) {
	var req forms.CreateFriendRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = repository.CreateFriend(req.Friends[0], req.Friends[1])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	successResp := forms.Response{
		Success: true,
	}
	respData, err := json.Marshal(&successResp)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respData)
}

func GetFriendList(w http.ResponseWriter, r *http.Request) {
	var req forms.FriendListRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	relationships, err := repository.GetFriendList(req.Email)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var friendIDs []int
	for _, relationship := range relationships {
		friendIDs = append(friendIDs, relationship.ID)
	}
	friends, err := repository.GetUserByIds(friendIDs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var friendEmails []string
	for _, friend := range friends {
		friendEmails = append(friendEmails, friend.Email)
	}

	resp := forms.FriendListResponse{
		Response: forms.Response{
			Success: true,
		},
		Friends: friendEmails,
		Count:   len(friendEmails),
	}
	respData, err := json.Marshal(&resp)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respData)
}

func GetMutualFriendList(w http.ResponseWriter, r *http.Request) {
	var req forms.MutualFriendListRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user1Relationships, err := repository.GetFriendList(req.Friends[0])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user2Relationships, err := repository.GetFriendList(req.Friends[1])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	commonRelations := pkg.GetMutualFriendList(user1Relationships, user2Relationships)
	var friendIDs []int
	for _, relationship := range commonRelations {
		friendIDs = append(friendIDs, relationship.ID)
	}
	friends, err := repository.GetUserByIds(friendIDs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var friendEmails []string
	for _, friend := range friends {
		friendEmails = append(friendEmails, friend.Email)
	}

	resp := forms.FriendListResponse{
		Response: forms.Response{
			Success: true,
		},
		Friends: friendEmails,
		Count:   len(friendEmails),
	}
	respData, err := json.Marshal(&resp)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respData)
}
