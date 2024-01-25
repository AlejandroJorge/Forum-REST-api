package domain

import "time"

type User struct {
	ID uint

	// For auth
	Email          string
	HashedPassword string

	// Meta
	RegistrationDate time.Time
}

type UserRepository interface {
	// Returns the user corresponding to the provided ID
	GetByID(id uint) (User, error)

	// Returns the user corresponding to the provided Email
	GetByEmail(email string) (User, error)

	// Creates a new user, the id and registrationDate in the model are ignored
	CreateNew(user User) error

	// Updates the email of the user corresponding to the provided ID
	UpdateEmail(id uint, newEmail string) error

	// Updates the password of the user corresponding to the provided ID. The password must be previously hashed
	UpdateHashedPassword(id uint, newHashedEmail string) error

	// Deletes the user corresponding to the provided ID
	Delete(id uint) error
}
