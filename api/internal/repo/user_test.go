package repo

import (
	"testing"

	"github.com/frozen599/s3-assignment/api/internal/config"
	"github.com/stretchr/testify/require"
)

func TestRepo_GetUserByEmail(t *testing.T) {
	tcs := map[string]struct {
		userEmail string
		exists    bool
	}{
		"user exists": {
			userEmail: "abc@gmail.com",
			exists:    true,
		},
		"user not exists": {
			userEmail: "def@gmai.com",
			exists:    false,
		},
	}
	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			cfg := config.NewConfig("./../../..")
			dbInstance := config.InitDB(cfg)
			defer dbInstance.Close()

			tx, err := dbInstance.Begin()
			require.NoError(t, err)
			defer tx.Rollback()

			repo := NewUserRepo(dbInstance)
			user, err := repo.GetUserByEmail(tc.userEmail)
			require.NoError(t, err)
			require.Equal(t, tc.exists, user != nil)
		})
	}
}

func TestRepo_GetUserByID(t *testing.T) {
	tcs := map[string]struct {
		userID int
		exists bool
	}{
		"user exists": {
			userID: 1,
			exists: true,
		},
		"user not exists": {
			userID: 2,
			exists: false,
		},
	}
	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			cfg := config.NewConfig("./../../..")
			dbInstance := config.InitDB(cfg)
			defer dbInstance.Close()

			tx, err := dbInstance.Begin()
			require.NoError(t, err)
			defer tx.Rollback()

			repo := NewUserRepo(dbInstance)
			users, err := repo.GetUserByIDs([]int{tc.userID})
			require.NoError(t, err)
			require.Equal(t, tc.exists, users[0].ID != 0)
		})
	}
}
