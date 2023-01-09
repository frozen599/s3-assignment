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

func TestController_Block(t *testing.T) {

	tcs := map[string]struct {
		input     forms.BlockingRequest
		expResult error
	}{
		"success": {
			input: forms.BlockingRequest{
				Requestor: "abc@gmail.com",
				Target:    "def@gmail.com",
			},
			expResult: nil,
		},
		"email not exist with first email": {
			input: forms.BlockingRequest{
				Requestor: "notexist@gmail.com",
				Target:    "def@gmail.com",
			},
			expResult: pkg.ErrUserNotFound,
		},
		"email not exist with target email": {
			input: forms.BlockingRequest{
				Requestor: "abc@gmail.com",
				Target:    "notexist@gmail.com",
			},
			expResult: pkg.ErrUserNotFound,
		},
		"requestor is already blocking target": {
			input: forms.BlockingRequest{
				Requestor: "abc@gmail.com",
				Target:    "def@gmail.com",
			},
			expResult: pkg.ErrCurrentUserIsBlockingTarget,
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
			blockingController := NewBlockingController(userRepo, relaRepo)
			err = blockingController.Block(tc.input.Requestor, tc.input.Target)
			if tc.expResult != nil {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}

}
