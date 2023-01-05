package pkg

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/frozen599/s3-assignment/api/internal/forms"
)

func ResponseOk(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	successResp := forms.Response{
		Success: true,
	}
	respData, err := json.Marshal(&successResp)
	if err != nil {
		log.Fatal(err)
		fmt.Fprintln(w, nil)
		return
	}
	fmt.Fprintln(w, respData)
}

func ResponseError(w http.ResponseWriter, code int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	resp := make(map[string]interface{})
	resp["success"] = false
	resp["message"] = err.Error()
	if code == http.StatusInternalServerError {
		resp["message"] = "INTERNAL SERVER ERROR"
	}
	respData, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
		fmt.Fprintln(w, nil)
	}
	fmt.Fprintln(w, respData)
}
