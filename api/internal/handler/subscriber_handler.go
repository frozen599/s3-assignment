package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/frozen599/s3-assignment/api/internal/controller"
	"github.com/frozen599/s3-assignment/api/internal/forms"
	"github.com/frozen599/s3-assignment/api/internal/pkg"
)

type subscriberHandler struct {
	subscriberController controller.SubscriberController
}

func NewSubscriberHandler(subscriberController controller.SubscriberController) subscriberHandler {
	return subscriberHandler{subscriberController: subscriberController}
}

func (h subscriberHandler) CreateSubscription(w http.ResponseWriter, r *http.Request) {
	var req forms.SubscribeToEmailRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	isValidInput := pkg.ValidateEmailInput([]string{req.Requestor, req.Target})
	if !isValidInput {
		pkg.ResponseError(w, 103, pkg.ErrInvalidEmailFormat)
		return
	}

	err = h.subscriberController.CreateSubScription(req.Requestor, req.Target)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	pkg.ResponseOk(w)
}

func (h subscriberHandler) CanReceiveUpdate(w http.ResponseWriter, r *http.Request) {
	var req forms.CanReceiveUpdateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		pkg.ResponseError(w, http.StatusBadRequest, err)
		return
	}
	isValidInput := pkg.ValidateEmailInput([]string{req.Sender})
	if !isValidInput {
		pkg.ResponseError(w, 103, pkg.ErrInvalidEmailFormat)
		return
	}

	emails, err := h.subscriberController.CanReceiveUpdate(req.Sender, req.Text)
	if err != nil {
		pkg.ResponseError(w, http.StatusInternalServerError, err)
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
		pkg.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, respData)
}
