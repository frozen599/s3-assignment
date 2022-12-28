package handler

import (
	"encoding/json"
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

	err = h.subscriberController.CreateSubScription(req.Requestor, req.Target)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.ResponseOk(w)
}
