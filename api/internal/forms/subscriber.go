package forms

type SubscribeToEmailRequest struct {
	Requestor string `json:"requestor"`
	Target    string `json:"target"`
}

type CanReceiveUpdateRequest struct {
	Sender string `json:"email"`
	Text   string `json:"text"`
}

type CanReceiveUpdateResponse struct {
	Response
	Recipients []string `json:"recipients"`
}
