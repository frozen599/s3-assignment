package controller

import (
	"io/ioutil"
	"testing"

	"github.com/frozen599/s3-assignment/api/internal/config"
	"github.com/frozen599/s3-assignment/api/internal/forms"
	"github.com/frozen599/s3-assignment/api/internal/pkg"
	"github.com/frozen599/s3-assignment/api/internal/repo"
	"github.com/stretchr/testify/require"
)

func TestController_CreateSubscription(t *testing.T) {
	tcs := map[string]struct {
		input     []string
		expResult error
	}{
		"success": {
			input:     []string{"abc@gmail.com", "def@gmail.com"},
			expResult: nil,
		},
		"user not exist with first email": {
			input:     []string{"notexist@gmail.com", "def@gmail.com"},
			expResult: pkg.ErrUserNotFound,
		},
		"user not exist with second email": {
			input:     []string{"abc@gmail.com", "notexist@gmail.com"},
			expResult: pkg.ErrUserNotFound,
		},
		"requestor is already subscribing to target": {
			input:     []string{"abc@gmail.com", "def@gmail.com"},
			expResult: pkg.ErrCurrentUserIsAlreadySubscribingTarget,
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
			userRepo := repo.NewUserRepo(dbInstance)
			relaRepo := repo.NewRelationshipRepo(dbInstance)
			subscriberController := NewSubscriberController(userRepo, relaRepo)
			err = subscriberController.CreateSubScription(tc.input[0], tc.input[1])
			if tc.expResult != nil {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}

}

func TestController_CanReceiveUpdate(t *testing.T) {
	tcs := map[string]struct {
		input     forms.CanReceiveUpdateRequest
		expResult error
	}{
		"sender have recipients": {
			input: forms.CanReceiveUpdateRequest{
				Sender: "abc@gmail.com",
				Text:   "Hello World! kate@example.com josh@example.com",
			},
			expResult: nil,
		},
		"sender user not exist": {
			input: forms.CanReceiveUpdateRequest{
				Sender: "noexist@gmail.com",
				Text:   "Hello World! kate@example.com",
			},
			expResult: pkg.ErrUserNotFound,
		},
		"sender can not have any recipients": {
			input: forms.CanReceiveUpdateRequest{
				Sender: "def@gmail.com",
				Text:   "Hello World!",
			},
			expResult: nil,
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
			userRepo := repo.NewUserRepo(dbInstance)
			relaRepo := repo.NewRelationshipRepo(dbInstance)
			subscriberController := NewSubscriberController(userRepo, relaRepo)
			recipients, err := subscriberController.GetCanReceiveUpdate(tc.input.Sender, []string{tc.input.Text})
			if tc.expResult != nil {
				require.Error(t, err)
				require.Empty(t, recipients)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, recipients)
			}
		})
	}

}
