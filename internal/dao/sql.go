package dao

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/timtosi/testo/internal/domain"

	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres database driver
)

// SQLConn is a `struct` implementing several DB querying methods related to SQL
// database.
type SQLConn struct {
	*gorm.DB
}

// NewSQLConn returns a new `*dao.SQLConn`.
func NewSQLConn(url string) (*SQLConn, error) {
	db, err := gorm.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &SQLConn{DB: db}, nil
}

// GetUsers returns a `[]*domain.User` representing all the users found in the
// `users` SQL table.
func (db *SQLConn) GetUsers() ([]*domain.User, error) {
	var usrs []*domain.User

	rows, err := db.Raw(`SELECT fullname, city, phone_number FROM users`).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var fullName string
		var city string
		var phoneNumber string

		if err := rows.Scan(&fullName, &city, &phoneNumber); err != nil {
			log.Println(fmt.Errorf("GetUsers: %v", err))
			continue
		}
		usrs = append(usrs, domain.NewUser(fullName, city, phoneNumber))
	}
	return usrs, nil
}

// AddUser inserts `usr` in the `users` table. It returns an `error` if
// something bad occurs.
func (db *SQLConn) AddUser(usr *domain.User) error {
	if usr == nil {
		return fmt.Errorf("AddUser: usr is nil")
	}

	if err := db.Exec(
		`INSERT INTO users (
			fullname,
			city,
			phone_number
		)
		VALUES ( ?, ?, ? )
		ON CONFLICT (phone_number)
		DO UPDATE SET (
			fullname,
			city,
			phone_number
		) = ( ?, ?, ? )`,
		usr.FullName,
		usr.City,
		usr.PhoneNumber,
		usr.FullName,
		usr.City,
		usr.PhoneNumber,
	).Error; err != nil {
		return fmt.Errorf("AddUser: %v", err)
	}
	return nil
}
