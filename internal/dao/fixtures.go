package dao

import (
	"fmt"
	"sort"

	"github.com/smotes/purse"
)

// TestDBURL is the DSN used by local test suite.
const TestDBURL string = "postgres://dev_env_user:dev_env_password@localhost:5432/test_db?sslmode=disable"

// Fixtures executes the SQL file `fileName` into the DB located at `TestDBURL`.
// It returns an `error` if something bad occurs.
//
// NOTE: Only the files in the relative path `../../deployments/fixtures`
// will be loaded.
func Fixtures(fileName string) error {
	ps, err := purse.New("../../deployments/fixtures")
	if err != nil {
		return fmt.Errorf("Fixtures: %v", err)
	}

	contents, ok := ps.Get(fileName)
	if !ok {
		return fmt.Errorf("Fixtures: %s - SQL file not loaded", fileName)
	}

	db, err := NewSQLConn(TestDBURL)
	if err != nil {
		return fmt.Errorf("Fixtures: %v", err)
	}
	defer func() { _ = db.Close() }()
	return db.Exec(contents).Error
}

// Migrations executes all SQL files found in `../../deployments/migrations`
// directory into the DB located at `TestDBURL`. It returns an `error` if
// something bad occurs.
//
// NOTE: The SQL files found will be executed in name order.
func Migrations() error {
	db, err := NewSQLConn(TestDBURL)
	if err != nil {
		return fmt.Errorf("Migrations: %v", err)
	}
	defer func() { _ = db.Close() }()

	ps, err := purse.New("../../deployments/migrations")
	if err != nil {
		return fmt.Errorf("Migrations: %v", err)
	}

	files := ps.Files()
	sort.Strings(files)
	for _, f := range files {
		contents, ok := ps.Get(f)
		if !ok {
			return fmt.Errorf("Migrations: SQL file not loaded")
		}
		if err := db.Exec(contents).Error; err != nil {
			return fmt.Errorf("Migrations: %s - %v", f, err)
		}
	}
	return nil
}
