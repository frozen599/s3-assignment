package handler

import (
	"fmt"
	"net/http"
)

type healthCheckHandler struct {
}

func NewHealthCheckHandler() healthCheckHandler {
	return healthCheckHandler{}
}

func (h healthCheckHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Println("cccccc")
	w.WriteHeader(200)
}
