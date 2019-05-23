package dao

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/timtosi/mimick-testing/internal/domain"
)

func TestSQLConn_NewSQLConn(t *testing.T) {
	testCases := []struct {
		name              string
		mockDBURL         string
		expectedErrorFunc func(assert.TestingT, interface{}, ...interface{}) bool
	}{
		{"ok", "postgres://dev_env_user:dev_env_password@localhost:5432/test_db?sslmode=disable", assert.Nil},
		{"bad_URL", "rgergerhg", assert.NotNil},
		{"no_remote_db", "postgres://hey:ho@lets.amazonaws.com:5432/go", assert.NotNil},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewSQLConn(tc.mockDBURL)
			tc.expectedErrorFunc(t, err)
		})
	}
}

func TestSQLConn_GetUsers(t *testing.T) {
	testCases := []struct {
		name              string
		fixtureFile       string
		expectedUsers     []*domain.User
		expectedErrorFunc func(assert.TestingT, interface{}, ...interface{}) bool
	}{
		{
			"ok_one",
			"getusers_ok_one.up.sql",
			[]*domain.User{
				&domain.User{
					FullName:    "Jeanne Dupont",
					City:        "Paris",
					PhoneNumber: "0625000000",
				},
			}, assert.Nil,
		},
		{
			"ok_multiple",
			"getusers_ok_multiple.up.sql",
			[]*domain.User{
				&domain.User{
					FullName:    "Sean Case",
					City:        "Paris",
					PhoneNumber: "0625000010",
				},
				&domain.User{
					FullName:    "Roberto Polo",
					City:        "Montpellier",
					PhoneNumber: "0625000011",
				},
				&domain.User{
					FullName:    "Lucas Robert",
					City:        "Vitry-sur-Seine",
					PhoneNumber: "0625000012",
				},
			}, assert.Nil,
		},
		{
			"error_notFound",
			"getusers_error_notFound.up.sql",
			[]*domain.User{},
			assert.Nil,
		},
	}

	db, err := NewSQLConn(TestDBURL)
	require.Nil(t, err)

	defer func() { _ = db.Close() }()

	for _, tc := range testCases {
		require.Nil(t, Fixtures(tc.fixtureFile))

		t.Run(tc.name, func(t *testing.T) {
			usrs, err := db.GetUsers()
			require.ElementsMatch(t, tc.expectedUsers, usrs)
			tc.expectedErrorFunc(t, err)
		})
		require.Nil(t, Fixtures("getusers.down.sql"))
	}
}

func TestSQLConn_AddUser(t *testing.T) {
	testCases := []struct {
		name              string
		fixtureFile       string
		mockUser          *domain.User
		expectedErrorFunc func(assert.TestingT, interface{}, ...interface{}) bool
	}{
		{
			"ok_regular",
			"adduser_ok_regular.up.sql",
			&domain.User{
				FullName:    "New User",
				City:        "Montpellier",
				PhoneNumber: "0625000020",
			},
			assert.Nil,
		},
		{
			"ok_overwriting",
			"adduser_ok_overwriting.up.sql",
			&domain.User{
				FullName:    "Already Exist",
				City:        "London",
				PhoneNumber: "0625000021",
			},
			assert.Nil,
		},
		{
			"error_nilUser",
			"adduser_error_nilUser.up.sql",
			nil,
			assert.NotNil,
		},
	}

	db, err := NewSQLConn(TestDBURL)
	require.Nil(t, err)

	defer func() { _ = db.Close() }()

	for _, tc := range testCases {
		require.Nil(t, Fixtures(tc.fixtureFile))

		t.Run(tc.name, func(t *testing.T) {
			err := db.AddUser(tc.mockUser)
			tc.expectedErrorFunc(t, err)
		})
		require.Nil(t, Fixtures("adduser.down.sql"))
	}
}

func TestMain(m *testing.M) {
	if err := Migrations(); err != nil {
		os.Exit(1)
	}
	os.Exit(m.Run())
}
