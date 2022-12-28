package handler

import (
	"net/http"
)

type healthCheckHandler struct {
}

func NewHealthCheckHandler() healthCheckHandler {
	return healthCheckHandler{}
}

func (h healthCheckHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
