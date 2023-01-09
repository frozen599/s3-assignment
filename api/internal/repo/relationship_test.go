package repo

import (
	"errors"
	"fmt"
	"io/ioutil"
	"testing"
	"time"

	"github.com/frozen599/s3-assignment/api/internal/config"
	"github.com/frozen599/s3-assignment/api/internal/models"
	"github.com/stretchr/testify/require"
)

func TestRepo_CreateFriend(t *testing.T) {
	type mockFriend struct {
		requester_id int
		target_id    int
	}

	tcs := map[string]struct {
		input     mockFriend
		expResult error
	}{
		"success": {
			input: mockFriend{
				requester_id: 4,
				target_id:    5,
			},
		},
		"fail": {
			input: mockFriend{
				requester_id: 4,
				target_id:    5,
			},
			expResult: errors.New("ERROR #23502 null value in column \"relationship_type\" of relation \"relationships\" violates not-null constraint"),
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

			repo := NewRelationshipRepo(dbInstance)
			err = repo.CreateRelationship(models.Relationship{
				UserID1:          tc.input.requester_id,
				UserID2:          tc.input.target_id,
				RelationshipType: models.RelationshipTypeFriend,
				CreatedAt:        time.Now(),
				UpdatedAt:        time.Now(),
			})
			if tc.expResult != nil {
				fmt.Println(err.Error())
				require.Error(t, err, err.Error())
			} else {
				require.NoError(t, err)
			}
		},
		)
	}
}

func TestRepo_CheckIfFriendConnectionExists(t *testing.T) {
	type mockFriend struct {
		userID1 int
		userID2 int
	}

	tcs := map[string]struct {
		input  mockFriend
		result bool
	}{
		"friendship exists": {
			input: mockFriend{
				userID1: 3,
				userID2: 4,
			},
			result: true,
		},
		"friendship not exists": {
			input: mockFriend{
				userID1: 4,
				userID2: 6,
			},
			result: false,
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
			repo := NewRelationshipRepo(dbInstance)
			exists, err := repo.CheckIfFriendConnectionExists(tc.input.userID1, tc.input.userID2)
			require.NoError(t, err)
			require.Equal(t, tc.result, exists)
		})
	}

}

func TestRepo_CheckIfIsBlockingTarget(t *testing.T) {
	type mockBlocking struct {
		userID1 int
		userID2 int
	}

	tcs := map[string]struct {
		input  mockBlocking
		result bool
	}{
		"is blocking": {
			input: mockBlocking{
				userID1: 1,
				userID2: 2,
			},
			result: true,
		},
		"is not blocking": {
			input: mockBlocking{
				userID1: 3,
				userID2: 4,
			},
			result: false,
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
			repo := NewRelationshipRepo(dbInstance)
			exists, err := repo.CheckIfFriendConnectionExists(tc.input.userID1, tc.input.userID2)
			require.NoError(t, err)
			require.Equal(t, tc.result, exists)
		})
	}
}

// func TestRepo_GetCanReceiveUpdate(t *testing.T) {
// 	tcs := map[string]struct {
// 		userID int
// 		result int
// 	}{
// 		"can receive update": {
// 			userID: 1,
// 			result: 2,
// 		},
// 		"cannot receive update": {
// 			userID: 2,
// 			result: 3,
// 		},
// 	}

// 	for desc, tc := range tcs {
// 		t.Run(desc, func(t *testing.T) {
// 			cfg := config.NewConfig("./../../..")
// 			dbInstance := config.InitDB(cfg)
// 			defer dbInstance.Close()

// 			initData, err := ioutil.ReadFile("./test_data/init_data.sql")
// 			require.NoError(t, err)
// 			_, err = dbInstance.Exec(string(initData))
// 			require.NoError(t, err)
// 			deleteData, err := ioutil.ReadFile("./test_data/delet_data.sql")
// 			require.NoError(t, err)
// 			defer dbInstance.Exec(deleteData)

// 			repo := NewRelationshipRepo(dbInstance)
// 			relas, err := repo.CanReceiveUpdate(tc.userID)
// 			require.NoError(t, err)
// 			require.Equal(t, tc.result, len(relas))
// 		})
// 	}
// }
