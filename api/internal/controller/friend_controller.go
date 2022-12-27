package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/frozen599/s3-assignment/api/internal/forms"
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

	fmt.Println("EMMAIL", req.Email)

	friends, err := repository.GetFriendList(req.Email)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var friendListEmails []string
	for _, friend := range friends {
		friendListEmails = append(friendListEmails, fmt.Sprint(friend.ID))
	}

	resp := forms.FriendListResponse{
		Response: forms.Response{
			Success: true,
		},
		Friends: friendListEmails,
		Count:   len(friendListEmails),
	}
	respData, err := json.Marshal(&resp)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respData)
}
