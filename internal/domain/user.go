package domain

// User is a `struct` representing a single user informations.
type User struct {
	FullName    string
	City        string
	PhoneNumber string
}

// NewUser returns a new `*domain.User`.
func NewUser(fullName, city, phoneNumber string) *User {
	return &User{
		FullName:    fullName,
		City:        city,
		PhoneNumber: phoneNumber,
	}
}
