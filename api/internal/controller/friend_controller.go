package controller

import (
	"encoding/json"
	"net/http"

	"github.com/frozen599/s3-assignment/api/internal/forms"
)

func CreateFriendConnection(w http.ResponseWriter, r *http.Request) {
	var req forms.CreateFriendRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
	w.Write(respData)
}

func GetFriendList(w http.ResponseWriter, r *http.Request) {
	var req forms.FriendListRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := forms.FriendListResponse{
		Response: forms.Response{
			Success: true,
		},
		Friends: []string{},
		Count:   0,
	}
	respData, err := json.Marshal(&resp)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(respData)
}
