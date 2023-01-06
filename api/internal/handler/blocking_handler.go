package handler

import (
	"encoding/json"
	"net/http"

	"github.com/frozen599/s3-assignment/api/internal/controller"
	"github.com/frozen599/s3-assignment/api/internal/forms"
	"github.com/frozen599/s3-assignment/api/internal/pkg"
)

type blockingHandler struct {
	blockingController controller.BlockingController
}

func NewBlockingHandler(blockingController controller.BlockingController) blockingHandler {
	return blockingHandler{blockingController: blockingController}
}

func (h blockingHandler) Block(w http.ResponseWriter, r *http.Request) {
	var req forms.BlockingRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		pkg.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	isValidInput := pkg.ValidateEmailInput([]string{req.Requestor, req.Target})
	if !isValidInput {
		pkg.ResponseError(w, 103, pkg.ErrInvalidEmailFormat)
		return
	}

	err = h.blockingController.BlockUpdate(req.Requestor, req.Target)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	pkg.ResponseOk(w)
}
