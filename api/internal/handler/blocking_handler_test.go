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

func TestHandler_Block(t *testing.T) {
	tcs := map[string]struct {
		input      string
		expBody    string
		expErr     error
		statusCode int
	}{
		"success": {
			input:      "{ \"requestor\": \"abc@gmail.com\", \"target\": \"def@gmail.com\"}",
			expBody:    "{\"success\":true}",
			statusCode: http.StatusOK,
			expErr:     nil,
		},
		"requestor email is invalid": {
			input:      "{ \"requestor\": \"abc\", \"target\": \"def@gmail.com\"}",
			expBody:    "{\"message\":\"invalid email format\",\"success\":false}",
			statusCode: http.StatusBadRequest,
			expErr:     pkg.ErrInvalidEmailFormat,
		},
		"target email is invalid": {
			input:      "{ \"requestor\": \"abc\", \"target\": \"def@gmail.com\"}",
			expBody:    "{\"message\":\"invalid email format\",\"success\":false}",
			statusCode: http.StatusBadRequest,
			expErr:     pkg.ErrInvalidEmailFormat,
		},
		"requestor user not exists": {
			input:      "{ \"requestor\": \"notexist@gmail.com\", \"target\": \"def@gmail.com\"}",
			expBody:    "{\"message\":\"user not found\",\"success\":false}",
			statusCode: http.StatusInternalServerError,
			expErr:     pkg.ErrUserNotFound,
		},
		"requestor is already blocking target": {
			input:      "{ \"requestor\": \"abc@gmail.com\", \"target\": \"def@gmail.com\"}",
			expBody:    "{\"message\":\"requestor is blocking target\",\"success\":false}",
			statusCode: http.StatusInternalServerError,
			expErr:     pkg.ErrUserNotFound,
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
			blockingController := controller.NewBlockingController(userRepo, relaRepo)
			handler := http.HandlerFunc(NewBlockingHandler(blockingController).Block)
			handler.ServeHTTP(res, req)
			if tc.expErr != nil {
				require.Equal(t, tc.statusCode, res.Code)
				require.Equal(t, tc.expBody, res.Body.String())
			} else {
				require.Equal(t, tc.statusCode, res.Code)
				require.Equal(t, tc.expBody, res.Body.String())

			}
		})
	}
}
