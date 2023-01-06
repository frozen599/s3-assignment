package controller

import (
	"errors"
	"io/ioutil"
	"testing"

	"github.com/frozen599/s3-assignment/api/internal/config"
	"github.com/frozen599/s3-assignment/api/internal/forms"
	"github.com/frozen599/s3-assignment/api/internal/models"
	"github.com/frozen599/s3-assignment/api/internal/pkg"
	"github.com/frozen599/s3-assignment/api/internal/repo"
	"github.com/stretchr/testify/require"
)

func TestController_BlockUpdate(t *testing.T) {
	type mockGetUserByEmail struct {
		firstEmail     models.User
		errFirstEmail  error
		secondEmail    models.User
		errSecondEmail error
	}

	type mockCheckIsFriend struct {
		check bool
		err   error
	}
	type mockCheckIsBlockEachOther struct {
		check bool
		err   error
	}

	tcs := map[string]struct {
		input                     forms.CreateFriendRequest
		mockGetUser               mockGetUserByEmail
		mockCheckIsFriend         mockCheckIsFriend
		mockCheckIsBlockEachOther mockCheckIsBlockEachOther
		mockCreateFriend          error
		expResult                 error
	}{
		"success": {
			input: forms.CreateFriendRequest{
				Friends: []string{"test1@gmail.com", "test2@gmai.com"},
			},
			mockGetUser: mockGetUserByEmail{
				firstEmail: models.User{
					ID:    1,
					Email: "test1@gmail.com",
				},
				secondEmail: models.User{
					ID:    2,
					Email: "test2@gmail.com",
				},
			},
		},
		"email not exist with first email": {
			input: forms.CreateFriendRequest{
				Friends: []string{"test1@gmail.com", "test2@gmail.com"},
			},
			mockGetUser: mockGetUserByEmail{
				errFirstEmail: pkg.ErrUserNotFound,
			},
			expResult: pkg.ErrUserNotFound,
		},
		"email not exist with target email": {
			input: forms.CreateFriendRequest{
				Friends: []string{"test1@gmail.com", "test2@gmail.com"},
			},
			mockGetUser: mockGetUserByEmail{
				firstEmail: models.User{
					ID:    1,
					Email: "test@gmail.com",
				},
				errSecondEmail: pkg.ErrUserNotFound,
			},
			expResult: pkg.ErrUserNotFound,
		},
		"2 email are already friend": {
			input: forms.CreateFriendRequest{
				Friends: []string{},
			},
			mockGetUser: mockGetUserByEmail{
				firstEmail: models.User{
					ID:    1,
					Email: "test@gmail.com",
				},
				secondEmail: models.User{
					ID:    2,
					Email: "test2@gmail.com",
				},
			},
			mockCheckIsFriend: mockCheckIsFriend{
				check: true,
			},
			expResult: pkg.ErrFriendshipAlreadyExists,
		},
		"2 email blocked each other": {
			input: forms.CreateFriendRequest{
				Friends: []string{},
			},
			mockGetUser: mockGetUserByEmail{
				firstEmail: models.User{
					ID:    1,
					Email: "test@gmail.com",
				},
				secondEmail: models.User{
					ID:    2,
					Email: "test2@gmail.com",
				},
			},
			mockCheckIsBlockEachOther: mockCheckIsBlockEachOther{
				check: true,
			},
			expResult: pkg.ErrCurrentUserIsBlockingTarget,
		},
		"something went wrong": {
			input: forms.CreateFriendRequest{
				Friends: []string{},
			},
			mockGetUser: mockGetUserByEmail{
				firstEmail: models.User{
					ID:    1,
					Email: "test@gmail.com",
				},
				secondEmail: models.User{
					ID:    2,
					Email: "test2@gmail.com",
				},
			},
			expResult: errors.New("something went wrong"),
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
			deleteData, err := ioutil.ReadFile("./test_data/delet_data.sql")
			require.NoError(t, err)
			defer dbInstance.Exec(deleteData)

			userRepo := repo.NewUserRepo(dbInstance)
			relaRepo := repo.NewRelationshipRepo(dbInstance)
			blockingController := NewBlockingController(userRepo, relaRepo)
			err = blockingController.Block(tc.input.Friends[0], tc.input.Friends[1])
			if tc.expResult != nil {
				require.Equal(t, err, tc.expResult)
			} else {
				require.NoError(t, err)
			}
		})
	}

}
