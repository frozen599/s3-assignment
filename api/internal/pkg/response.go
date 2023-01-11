package pkg

import (
	"encoding/json"
	"fmt"
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
		fmt.Println(err)
	}
	w.Write(respData)
}

func ResponseError(w http.ResponseWriter, code int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	resp := make(map[string]interface{})
	resp["success"] = false
	resp["message"] = err.Error()
	respData, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(respData)
}
