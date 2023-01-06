package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/frozen599/s3-assignment/api/internal/controller"
	"github.com/frozen599/s3-assignment/api/internal/pkg"
	"github.com/stretchr/testify/require"
)

func TestHandler_CreateFriend(t *testing.T) {
	tcs := map[string]struct {
		input      string
		resultSrv  error
		expBody    string
		expErr     error
		statusCode int
	}{
		"success": {
			input:      "{\"friends\":[\"john@example.com\", \"andy\"]}",
			expBody:    `{"success":"true"}`,
			statusCode: http.StatusOK,
			expErr:     pkg.ErrInvalidEmailFormat,
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			res := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/v1/friends", strings.NewReader(tc.input))

			friendController := controller.NewBlockingController()
			handler := http.HandlerFunc(NewFriendHanlder().CreateFriendConnection)
			handler.ServeHTTP(res, req)
			if tc.expErr != nil {
				require.Equal(t, tc.statusCode, res.Code)
			} else {
				//require.Equal(t, tc.expBody, res.Body.String())
				require.Equal(t, tc.statusCode, res.Code)
			}
		})
	}
}
