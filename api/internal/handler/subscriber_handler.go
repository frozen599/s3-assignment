package hanlder

import (
	"encoding/json"
	"net/http"

	"github.com/frozen599/s3-assignment/api/internal/forms"
)

func CreateSubscription(w http.ResponseWriter, r *http.Request) {
	var req forms.SubscribeToEmailRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return
	}

}
