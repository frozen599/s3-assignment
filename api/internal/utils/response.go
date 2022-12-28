package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/frozen599/s3-assignment/api/internal/forms"
)

func ResponseOk(w http.ResponseWriter) {
	successResp := forms.Response{
		Success: true,
	}
	respData, err := json.Marshal(&successResp)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(respData)
}

func ResponseOkWithData(w http.ResponseWriter, data []byte) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
