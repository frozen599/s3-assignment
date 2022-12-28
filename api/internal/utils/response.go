package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/frozen599/s3-assignment/api/internal/forms"
)

func ResponseOk(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	successResp := forms.Response{
		Success: true,
	}
	respData, err := json.Marshal(&successResp)
	if err != nil {
		log.Fatal(err)
		w.Write(nil)
		return
	}
	w.Write(respData)
}

func ResponseOkWithData(w http.ResponseWriter, data []byte) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func ResponseError(w http.ResponseWriter, code int, err error) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]interface{})
	resp["success"] = true
	resp["message"] = err.Error()
	if code == http.StatusInternalServerError {
		resp["message"] = "INTERNAL SERVER ERROR"
	}
	respData, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
		w.Write(nil)
	}
	w.Write(respData)
}
