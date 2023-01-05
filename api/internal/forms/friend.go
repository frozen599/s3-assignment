package forms

type CreateFriendRequest struct {
	Friends []string `json:"friends"`
}

type FriendListRequest struct {
	Email string `json:"email"`
}

type FriendListResponse struct {
	Response
	Friends []string `json:"friends"`
	Count   int      `json:"count"`
}

type MutualFriendListRequest struct {
	Friends []string `json:"friends"`
}

type MutualFriendListResponse struct {
	Response
	Friends []string `json:"friends"`
	Count   int      `json:"count"`
}
