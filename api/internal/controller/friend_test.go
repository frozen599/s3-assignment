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

func TestController_CreateFriend(t *testing.T) {
	tcs := map[string]struct {
		input            forms.CreateFriendRequest
		mockCreateFriend error
		expResult        error
	}{
		"success": {
			input: forms.CreateFriendRequest{
				Friends: []string{"abc@gmail.com", "def@gmail.com"},
			},
			expResult: nil,
		},
		"email not exist with first email": {
			input: forms.CreateFriendRequest{
				Friends: []string{"notexist@gmail.com", "def@gmail.com"},
			},
			expResult: pkg.ErrUserNotFound,
		},
		"email not exist with target email": {
			input: forms.CreateFriendRequest{
				Friends: []string{"abc@gmail.com", "notexist@gmail.com"},
			},
			expResult: pkg.ErrUserNotFound,
		},
		"2 email are already friend": {
			input: forms.CreateFriendRequest{
				Friends: []string{"abc@gmail.com", "def@gmail.com"},
			},
			expResult: pkg.ErrFriendshipAlreadyExists,
		},
		"2 email are blocking each other": {
			input: forms.CreateFriendRequest{
				Friends: []string{"abc@gmail.com", "ghi@gmail.com"},
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
			friendController := NewFriendController(userRepo, relaRepo)
			err = friendController.CreateFriendConnection(tc.input.Friends[0], tc.input.Friends[1])
			if tc.expResult != nil {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestController_GetFriendList(t *testing.T) {
	tcs := map[string]struct {
		input     forms.FriendListRequest
		expResult error
	}{
		"user have friends": {
			input: forms.FriendListRequest{
				Email: "abc@gmail.com",
			},
			expResult: nil,
		},
		"user not exists": {
			input: forms.FriendListRequest{
				Email: "notexist@gmai.com",
			},
			expResult: pkg.ErrUserNotFound,
		},
		"user does have any friends": {
			input: forms.FriendListRequest{
				Email: "def@gmail.com",
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
			friendController := NewFriendController(userRepo, relaRepo)
			results, err := friendController.GetFriendList(tc.input.Email)
			if tc.expResult != nil {
				require.Error(t, err)
				require.Empty(t, results)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, results)
			}
		})
	}

}

func TestController_GetMutualFriendList(t *testing.T) {
	tcs := map[string]struct {
		input     forms.MutualFriendListRequest
		expResult error
	}{
		"success": {
			input: forms.MutualFriendListRequest{
				Friends: []string{"abc@gmail.com", "def@gmail.com"},
			},
			expResult: nil,
		},
		"user not exists with first email": {
			input: forms.MutualFriendListRequest{
				Friends: []string{"notexist@gmail.com", "def@gmail.com"},
			},
			expResult: pkg.ErrUserNotFound,
		},
		"user not exists with sencond email": {
			input: forms.MutualFriendListRequest{
				Friends: []string{"abc@gmail.com", "notexist@gmail.com"},
			},
			expResult: pkg.ErrUserNotFound,
		},
		"two users do not have mutual friends": {
			input: forms.MutualFriendListRequest{
				Friends: []string{"abc@gmail.com", "def@gmail.com"},
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
			friendController := NewFriendController(userRepo, relaRepo)
			results, err := friendController.GetMutualFriendList(tc.input.Friends[0], tc.input.Friends[1])
			if tc.expResult != nil {
				require.Error(t, err)
				require.Empty(t, results)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, results)
			}
		})
	}

}
