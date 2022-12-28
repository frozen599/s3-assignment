package handler

import (
	"encoding/json"
	"net/http"

	"github.com/frozen599/s3-assignment/api/internal/controller"
	"github.com/frozen599/s3-assignment/api/internal/forms"
	"github.com/frozen599/s3-assignment/api/internal/utils"
)

type blockingHandler struct {
	blockingController controller.BlockingController
}

func NewBlockingHandler() blockingHandler {
	return blockingHandler{blockingController: controller.NewBlockingController()}
}

func (h blockingHandler) BlockUpdate(w http.ResponseWriter, r *http.Request) {
	var req forms.BlockingRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	err = h.blockingController.BlockUpdate(req.Requestor, req.Target)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.ResponseOk(w)
}
