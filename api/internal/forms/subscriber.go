package forms

type SubscribeToEmailRequest struct {
	Requestor string `json:"requestor"`
	Target    string `json:"target"`
}
