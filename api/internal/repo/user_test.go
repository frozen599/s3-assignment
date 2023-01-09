package repo

import (
	"io/ioutil"
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
			userEmail: "noexist@gmai.com",
			exists:    false,
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
			repo := NewUserRepo(dbInstance)
			user, err := repo.GetUserByEmail(tc.userEmail)
			require.NoError(t, err)
			require.Equal(t, user != nil, tc.exists)
		})
	}
}

func TestRepo_GetUserByIDs(t *testing.T) {
	tcs := map[string]struct {
		userID int
		exists bool
	}{
		"user exists": {
			userID: 1,
			exists: true,
		},
		"user not exists": {
			userID: 3,
			exists: false,
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
			repo := NewUserRepo(dbInstance)
			users, err := repo.GetUserByIDs([]int{tc.userID})
			require.NoError(t, err)
			require.Equal(t, tc.exists, len(users) > 0)
		})
	}
}

func TestRepo_GetUserByEmails(t *testing.T) {
	tcs := map[string]struct {
		email  string
		exists bool
	}{
		"user exists": {
			email:  "abc@gmail.com",
			exists: true,
		},
		"user not exists": {
			email:  "noexist@gmail.com",
			exists: false,
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
			repo := NewUserRepo(dbInstance)
			users, err := repo.GetUserByEmails([]string{tc.email})
			require.NoError(t, err)
			require.Equal(t, tc.exists, len(users) > 0)
		})
	}
}
