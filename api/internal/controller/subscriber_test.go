package controller

import (
	"io/ioutil"
	"testing"

	"github.com/frozen599/s3-assignment/api/internal/config"
	"github.com/frozen599/s3-assignment/api/internal/forms"
	"github.com/frozen599/s3-assignment/api/internal/repo"
	"github.com/stretchr/testify/require"
)

func TestController_CreateSubscription(t *testing.T) {
	tcs := map[string]struct {
		input            []string
		mockCreateFriend error
		expResult        error
	}{
		"success": {
			input: []string{"abc@gmail.com", "def@gmail.com"},
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			cfg := config.NewConfig("./../../..")
			dbInstance := config.InitDB(cfg)
			defer dbInstance.Close()

			initData, err := ioutil.ReadFile("./test_data/init_data.sql")
			require.NoError(t, err)
			_, err = dbInstance.Exec(string(initData))
			require.NoError(t, err)
			deleteData, err := ioutil.ReadFile("./../test_data/delete_data.sql")
			require.NoError(t, err)
			defer dbInstance.Exec(deleteData)

			userRepo := repo.NewUserRepo(dbInstance)
			relaRepo := repo.NewRelationshipRepo(dbInstance)
			subscriberController := NewSubscriberController(userRepo, relaRepo)
			err = subscriberController.CreateSubScription(tc.input[0], tc.input[1])
			if tc.expResult != nil {
				require.Equal(t, err, tc.expResult)
			} else {
				require.NoError(t, err)
			}
		})
	}

}

func TestController_CanReceiveUpdate(t *testing.T) {
	tcs := map[string]struct {
		input            forms.CanReceiveUpdateRequest
		mockCreateFriend error
		expResult        error
	}{
		"success": {
			input: forms.CanReceiveUpdateRequest{
				Sender: "abc@gmail.com",
				Text:   "ssss",
			},
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			cfg := config.NewConfig("./../../..")
			dbInstance := config.InitDB(cfg)
			defer dbInstance.Close()

			initData, err := ioutil.ReadFile("./test_data/init_data.sql")
			require.NoError(t, err)
			_, err = dbInstance.Exec(string(initData))
			require.NoError(t, err)
			deleteData, err := ioutil.ReadFile("./../test_data/delete_data.sql")
			require.NoError(t, err)
			defer dbInstance.Exec(deleteData)

			userRepo := repo.NewUserRepo(dbInstance)
			relaRepo := repo.NewRelationshipRepo(dbInstance)
			subscriberController := NewSubscriberController(userRepo, relaRepo)
			_, err = subscriberController.CanReceiveUpdate(tc.input.Sender, []string{tc.input.Text})
			if tc.expResult != nil {
				require.Equal(t, err, tc.expResult)
			} else {
				require.NoError(t, err)
			}
		})
	}

}
