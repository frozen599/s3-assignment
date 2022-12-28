package hanlder

import (
	"fmt"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Println("cccccc")
	w.WriteHeader(200)
}
