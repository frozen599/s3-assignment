package handler

import (
	"encoding/json"
	"net/http"

	"github.com/frozen599/s3-assignment/api/internal/controller"
	"github.com/frozen599/s3-assignment/api/internal/forms"
	"github.com/frozen599/s3-assignment/api/internal/pkg"
)

type friendHandler struct {
	friendController controller.FriendController
}

func NewFriendHanlder(friendController controller.FriendController) friendHandler {
	return friendHandler{friendController: friendController}
}

func (h friendHandler) CreateFriendConnection(w http.ResponseWriter, r *http.Request) {
	var req forms.CreateFriendRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		pkg.ResponseError(w, http.StatusBadRequest, pkg.ErrRequestBodyMalformed)
		return
	}

	isValidInput := pkg.ValidateEmailInput([]string{req.Friends[0], req.Friends[1]})
	if !isValidInput {
		pkg.ResponseError(w, http.StatusBadRequest, pkg.ErrInvalidEmailFormat)
		return
	}

	err = h.friendController.CreateFriendConnection(req.Friends[0], req.Friends[1])
	if err != nil {
		pkg.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	pkg.ResponseOk(w)
}

func (h friendHandler) GetFriendList(w http.ResponseWriter, r *http.Request) {
	var req forms.FriendListRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		pkg.ResponseError(w, http.StatusBadRequest, pkg.ErrRequestBodyMalformed)
		return
	}

	isValidInput := pkg.ValidateEmailInput([]string{req.Email})
	if !isValidInput {
		pkg.ResponseError(w, http.StatusBadRequest, pkg.ErrInvalidEmailFormat)
		return
	}

	users, err := h.friendController.GetFriendList(req.Email)
	if err != nil {
		pkg.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	var emails []string
	for _, user := range users {
		emails = append(emails, user.Email)
	}

	resp := forms.FriendListResponse{
		Response: forms.Response{
			Success: true,
		},
		Friends: emails,
		Count:   len(emails),
	}
	respData, err := json.Marshal(&resp)
	if err != nil {
		pkg.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(respData)
}

func (h friendHandler) GetMutualFriendList(w http.ResponseWriter, r *http.Request) {
	var req forms.MutualFriendListRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		pkg.ResponseError(w, http.StatusBadRequest, pkg.ErrRequestBodyMalformed)
		return
	}

	isValidInput := pkg.ValidateEmailInput([]string{req.Friends[0], req.Friends[1]})
	if !isValidInput {
		pkg.ResponseError(w, http.StatusBadRequest, pkg.ErrInvalidEmailFormat)
		return
	}

	users, err := h.friendController.GetMutualFriendList(req.Friends[0], req.Friends[1])
	if err != nil {
		pkg.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	var emails []string
	for _, user := range users {
		emails = append(emails, user.Email)
	}

	resp := forms.FriendListResponse{
		Response: forms.Response{
			Success: true,
		},
		Friends: emails,
		Count:   len(emails),
	}
	respData, err := json.Marshal(&resp)
	if err != nil {
		pkg.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(respData)
}
