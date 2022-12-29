package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/frozen599/s3-assignment/api/internal/controller"
	"github.com/frozen599/s3-assignment/api/internal/forms"
	"github.com/frozen599/s3-assignment/api/internal/utils"
)

type subscriberHandler struct {
	subscriberController controller.SubscriberController
}

func NewSubscriberHandler() subscriberHandler {
	return subscriberHandler{subscriberController: controller.NewSubscriberController()}
}

func (h subscriberHandler) CreateSubscription(w http.ResponseWriter, r *http.Request) {
	var req forms.SubscribeToEmailRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	isValidInput := utils.ValidateEmailInput([]string{req.Requestor, req.Target})
	if !isValidInput {
		utils.ResponseError(w, 103, utils.ErrInvalidEmailFormat)
		return
	}

	err = h.subscriberController.CreateSubScription(req.Requestor, req.Target)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.ResponseOk(w)
}

func (h subscriberHandler) CanReceiveUpdate(w http.ResponseWriter, r *http.Request) {
	var req forms.CanReceiveUpdateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err)
		return
	}
	isValidInput := utils.ValidateEmailInput([]string{req.Sender})
	if !isValidInput {
		utils.ResponseError(w, 103, utils.ErrInvalidEmailFormat)
		return
	}

	emails, err := h.subscriberController.CanReceiveUpdate(req.Sender, req.Text)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	resp := forms.CanReceiveUpdateResponse{
		Response: forms.Response{
			Success: true,
		},
		Recipients: emails,
	}
	respData, err := json.Marshal(resp)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, respData)
}
