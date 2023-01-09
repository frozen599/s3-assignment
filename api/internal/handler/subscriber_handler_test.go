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

func TestHandler_CreateSubscription(t *testing.T) {
	tcs := map[string]struct {
		input      string
		expBody    string
		expErr     error
		statusCode int
	}{
		"success": {
			input:      "{\"requestor\":\"abc@gmail.com\",\"target\":\"def@gmail.com\"}",
			expBody:    "{\"success\":true}",
			statusCode: http.StatusOK,
			expErr:     nil,
		},
		"first user not exist": {
			input:      "{\"requestor\":\"noexist@example.com\",\"target\":\"def@gmail.com\"}",
			expBody:    "{\"message\":\"user not found\",\"success\":false}",
			statusCode: http.StatusInternalServerError,
			expErr:     pkg.ErrUserNotFound,
		},
		"second user not exist": {
			input:      "{\"requestor\":\"abc@gmail.com\",\"target\":\"noexist@gmail.com\"}",
			expBody:    "{\"message\":\"user not found\",\"success\":false}",
			statusCode: http.StatusInternalServerError,
			expErr:     pkg.ErrUserNotFound,
		},
		"is already subscribing": {
			input:      "{\"requestor\":\"abc@gmail.com\",\"target\":\"def@gmail.com\"}",
			expBody:    "{\"message\":\"requestor is already subscribing target\",\"success\":false}",
			statusCode: http.StatusInternalServerError,
			expErr:     pkg.ErrCurrentUserIsAlreadySubscribingTarget,
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
			subscriberController := controller.NewSubscriberController(userRepo, relaRepo)
			handler := http.HandlerFunc(NewSubscriberHandler(subscriberController).CreateSubscription)
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

func TestHandler_CanReceiveUpdate(t *testing.T) {
	tcs := map[string]struct {
		input      string
		expBody    string
		expErr     error
		statusCode int
	}{
		"success": {
			input:      "{\"sender\":\"abc@gmail.com\",\"text\":\"Hello World! ghi@gmail.com\"}",
			expBody:    "{\"success\":true,\"recipients\":[\"def@gmail.com\",\"ghi@gmail.com\"]}",
			statusCode: http.StatusOK,
			expErr:     nil,
		},
		"sender not exists": {
			input:      "{\"sender\":\"notexist@gmail.com\",\"text\":\"Hello World! kate@gmail.com\"}",
			expBody:    "{\"message\":\"user not found\",\"success\":false}",
			statusCode: http.StatusInternalServerError,
			expErr:     pkg.ErrUserNotFound,
		},
		"sender does not have can receive update": {
			input:      "{\"sender\":\"def@gmail.com\",\"text\":\"Hello World!\"}",
			expBody:    "{\"success\":true,\"recipients\":null}",
			statusCode: http.StatusOK,
			expErr:     nil,
		},
	}

	cfg := config.NewConfig("./../../..")
	dbInstance := config.InitDB(cfg)
	defer dbInstance.Close()
	//initData, err := ioutil.ReadFile("./test_data/init_data.sql")
	//require.NoError(t, err)
	//_, err = dbInstance.Exec(string(initData))
	//require.NoError(t, err)
	//deleteData, err := ioutil.ReadFile("./test_data/delete_data.sql")
	//require.NoError(t, err)
	//defer dbInstance.Exec(string(deleteData))

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			res := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/v1/friends", strings.NewReader(tc.input))
			userRepo := repo.NewUserRepo(dbInstance)
			relaRepo := repo.NewRelationshipRepo(dbInstance)
			subscriberController := controller.NewSubscriberController(userRepo, relaRepo)
			handler := http.HandlerFunc(NewSubscriberHandler(subscriberController).GetCanReceiveUpdate)
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
