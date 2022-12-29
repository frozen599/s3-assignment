package utils

import "fmt"

type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewApiError(code int, message string) ApiError {
	return ApiError{
		Code:    code,
		Message: message,
	}
}

func (e ApiError) Error() string {
	return fmt.Sprintf("code: %d - message: %s", e.Code, e.Message)
}

var (
	ErrFriendshipAlreadyExists     = NewApiError(100, "friendship already exists")
	ErrCurrentUserIsBlockingTarget = NewApiError(101, "requestor is blocking target")
)
