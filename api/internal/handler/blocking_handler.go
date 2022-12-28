package handler

import (
	"encoding/json"
	"net/http"

	"github.com/frozen599/s3-assignment/api/internal/forms"
	"github.com/frozen599/s3-assignment/api/internal/usescase"
)

type blockingHandler struct {
	blockingUseCase usescase.BlockingUseCase
}

func NewBlockingHandler() blockingHandler {
	return blockingHandler{blockingUseCase: usescase.NewBlockingUseCase()}
}

func (h blockingHandler) BlockUpdate(w http.ResponseWriter, r *http.Request) {
	var req forms.BlockingRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.blockingUseCase.BlockUpdate(req.Requestor, req.Target)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
