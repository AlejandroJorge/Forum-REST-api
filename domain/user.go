package domain

import (
	"time"

	"github.com/AlejandroJorge/forum-rest-api/util"
)

type User struct {
	ID uint

	// For auth
	Email          string
	HashedPassword string

	// Meta
	RegistrationDate time.Time
}

func (u User) Validate() bool {
	conditions := []bool{
		u.ID != 0,
		u.Email != "",
		u.HashedPassword != "",
		!u.RegistrationDate.IsZero(),
	}

	return util.MergeAND(conditions)
}

type UserRepository interface {
	// Returns the user corresponding to the provided ID
	GetByID(id uint) (User, error)

	// Returns the user corresponding to the provided Email
	GetByEmail(email string) (User, error)

	// Creates a new user, the id and registrationDate in the model are ignored
	CreateNew(user User) (uint, error)

	// Updates the email of the user corresponding to the provided ID
	UpdateEmail(id uint, newEmail string) error

	// Updates the password of the user corresponding to the provided ID. The password must be previously hashed
	UpdateHashedPassword(id uint, newHashedEmail string) error

	// Deletes the user corresponding to the provided ID
	Delete(id uint) error
}

type UserService interface {
	// Retrieves a user by the id provided
	GetByID(id uint) (User, error)

	// Retrieves a user by the email provided
	GetByEmail(email string) (User, error)

	// Creates a new user with the info provided
	CreateNew(createInfo struct {
		NewEmail          string
		NewHashedPassword string
	}) (uint, error)

	// Updates the authentication info of the user with corresponding id
	Update(id uint, updateInfo struct {
		UpdatedEmail          string
		UpdatedHashedPassword string
	}) error

	// Deletes the user with the provided id
	Delete(id uint) error
}
