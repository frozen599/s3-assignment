package forms

type BlockingRequest struct {
	Requestor string `json:"requestor"`
	Target    string `json:"target"`
}
