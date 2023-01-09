package handler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/frozen599/s3-assignment/api/internal/config"
	"github.com/frozen599/s3-assignment/api/internal/controller"
	"github.com/frozen599/s3-assignment/api/internal/pkg"
	"github.com/frozen599/s3-assignment/api/internal/repo"
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
			input:      "{\"friends\":[\"abc@gmail.com\",\"def@gmail.com\"]}",
			expBody:    "{\"success\":true}",
			statusCode: http.StatusOK,
			expErr:     nil,
		},
		"first user not exist": {
			input:      "{\"friends\":[\"noexist@example.com\",\"def@gmail.com\"]}",
			expBody:    "{\"message\":\"user not found\",\"success\":false}",
			statusCode: http.StatusInternalServerError,
			expErr:     pkg.ErrUserNotFound,
		},
		"second user not exist": {
			input:      "{\"friends\":[\"abc@gmail.com\",\"noexist@gmail.com\"]}",
			expBody:    "{\"message\":\"user not found\",\"success\":false}",
			statusCode: http.StatusInternalServerError,
			expErr:     pkg.ErrUserNotFound,
		},
		"is already friend": {
			input:      "{\"friends\":[\"abc@gmail.com\",\"def@gmail.com\"]}",
			expBody:    "{\"message\":\"friendship already exists\",\"success\":false}",
			statusCode: http.StatusInternalServerError,
			expErr:     pkg.ErrFriendshipAlreadyExists,
		},
	}

	cfg := config.NewConfig("./../../..")
	dbInstance := config.InitDB(cfg)
	defer dbInstance.Close()

	initData, err := ioutil.ReadFile("./test_data/init_data.sql")
	require.NoError(t, err)
	_, err = dbInstance.Exec(string(initData))
	require.NoError(t, err)
	deleteData, err := ioutil.ReadFile("./test_data/delete_data.sql")
	require.NoError(t, err)
	defer dbInstance.Exec(string(deleteData))

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			res := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/v1/friends", strings.NewReader(tc.input))

			userRepo := repo.NewUserRepo(dbInstance)
			relaRepo := repo.NewRelationshipRepo(dbInstance)

			friendController := controller.NewFriendController(userRepo, relaRepo)
			handler := http.HandlerFunc(NewFriendHanlder(friendController).CreateFriendConnection)
			handler.ServeHTTP(res, req)
			if tc.expErr != nil {
				require.Equal(t, tc.statusCode, res.Code)
				require.Equal(t, tc.expBody, res.Body.String())
			} else {
				require.Equal(t, tc.expBody, res.Body.String())
				require.Equal(t, tc.statusCode, res.Code)
			}
		})
	}
}

func TestHandler_GetFriendList(t *testing.T) {
	tcs := map[string]struct {
		input      string
		resultSrv  error
		expBody    string
		expErr     error
		statusCode int
	}{
		"success": {
			input:      "{\"email\":\"abc@gmail.com\"}",
			expBody:    "{\"success\":true,\"friends\":[\"def@gmail.com\",\"ghi@gmail.com\"],\"count\":2}",
			statusCode: http.StatusOK,
			expErr:     nil,
		},
		"user not exist": {
			input:      "{\"email\":\"john@gmail.com\"}",
			expBody:    "{\"message\":\"user not found\",\"success\":false}",
			statusCode: http.StatusInternalServerError,
		},
		"user does not have any friends": {
			input:      "{\"email\":\"jkl@gmail.com\"}",
			expBody:    "{\"success\":true,\"friends\":null,\"count\":0}",
			statusCode: http.StatusOK,
		},
	}

	cfg := config.NewConfig("./../../..")
	dbInstance := config.InitDB(cfg)
	defer dbInstance.Close()

	initData, err := ioutil.ReadFile("./test_data/init_friend_data.sql")
	require.NoError(t, err)
	_, err = dbInstance.Exec(string(initData))
	require.NoError(t, err)
	deleteData, err := ioutil.ReadFile("./test_data/delete_data.sql")
	require.NoError(t, err)
	defer dbInstance.Exec(string(deleteData))

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			res := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/v1/friends", strings.NewReader(tc.input))

			userRepo := repo.NewUserRepo(dbInstance)
			relaRepo := repo.NewRelationshipRepo(dbInstance)

			friendController := controller.NewFriendController(userRepo, relaRepo)
			handler := http.HandlerFunc(NewFriendHanlder(friendController).GetFriendList)
			handler.ServeHTTP(res, req)
			if tc.expErr != nil {
				require.Equal(t, tc.statusCode, res.Code)
				require.Equal(t, tc.expBody, res.Body.String())
			} else {
				require.Equal(t, tc.expBody, res.Body.String())
				require.Equal(t, tc.statusCode, res.Code)
			}
		})
	}
}

func TestHandler_GetMutualFriendList(t *testing.T) {
	tcs := map[string]struct {
		input      string
		resultSrv  error
		expBody    string
		expErr     error
		statusCode int
	}{
		"success": {
			input:      "{\"friends\":[\"def@gmail.com\",\"abc@gmail.com\"]}",
			expBody:    "{\"success\":true,\"friends\":[\"ghi@gmail.com\"],\"count\":1}",
			statusCode: http.StatusOK,
			expErr:     nil,
		},
		"first user not exist": {
			input:      "{\"friends\":[\"noexist@example.com\",\"def@gmail.com\"]}",
			expBody:    "{\"message\":\"user not found\",\"success\":false}",
			statusCode: http.StatusInternalServerError,
			expErr:     pkg.ErrUserNotFound,
		},
		"second user not exist": {
			input:      "{\"friends\":[\"abc@gmail.com\",\"noexist@gmail.com\"]}",
			expBody:    "{\"message\":\"user not found\",\"success\":false}",
			statusCode: http.StatusInternalServerError,
			expErr:     pkg.ErrUserNotFound,
		},
		"two users do not have mutual friends": {
			input:      "{\"friends\":[\"def@gmail.com\",\"jkl@gmail.com\"]}",
			expBody:    "{\"success\":true,\"friends\":null,\"count\":0}",
			statusCode: http.StatusOK,
			expErr:     nil,
		},
	}

	cfg := config.NewConfig("./../../..")
	dbInstance := config.InitDB(cfg)
	defer dbInstance.Close()

	initData, err := ioutil.ReadFile("./test_data/init_friend_data.sql")
	require.NoError(t, err)
	_, err = dbInstance.Exec(string(initData))
	require.NoError(t, err)
	deleteData, err := ioutil.ReadFile("./test_data/delete_data.sql")
	require.NoError(t, err)
	defer dbInstance.Exec(string(deleteData))

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			res := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/v1/friends", strings.NewReader(tc.input))
			userRepo := repo.NewUserRepo(dbInstance)
			relaRepo := repo.NewRelationshipRepo(dbInstance)
			friendController := controller.NewFriendController(userRepo, relaRepo)
			handler := http.HandlerFunc(NewFriendHanlder(friendController).GetMutualFriendList)
			handler.ServeHTTP(res, req)
			if tc.expErr != nil {
				require.Equal(t, tc.statusCode, res.Code)
			} else {
				require.Equal(t, tc.expBody, res.Body.String())
				require.Equal(t, tc.statusCode, res.Code)
			}
		})
	}
}
