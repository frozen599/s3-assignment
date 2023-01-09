package pkg

type ApiError string

func (e ApiError) Error() string {
	return string(e)
}

var (
	ErrFriendshipAlreadyExists               = ApiError("friendship already exists")
	ErrCurrentUserIsBlockingTarget           = ApiError("requestor is blocking target")
	ErrCurrentUserIsAlreadySubscribingTarget = ApiError("requestor is already subscribing target")
	ErrUserNotFound                          = ApiError("user not found")
	ErrInvalidEmailFormat                    = ApiError("invalid email format")
	ErrRequestBodyMalformed                  = ApiError("request body is malformed")
)
