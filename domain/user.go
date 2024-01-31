package domain

import (
	"time"

	"github.com/AlejandroJorge/forum-rest-api/util"
)

type User struct {
	ID               uint      `json:"ID"`
	Email            string    `json:"Email"`
	HashedPassword   string    `json:"HashedPassword"`
	RegistrationDate time.Time `json:"RegistrationDate"`
}

func (u User) Validate() bool {
	conditions := []bool{
		u.ID != 0,
		u.Email != "",
		u.HashedPassword != "",
		!u.RegistrationDate.IsZero(),
		util.IsEmailFormat(u.Email),
	}

	return util.MergeAND(conditions)
}

type UserRepository interface {
	// Returns the ID of the created user and can return ErrRepeatedEntity
	Create(email, hashedPassword string) (uint, error)

	// Can return ErrNoRowsAffected
	Delete(id uint) error

	// Can return ErrNoRowsAffected
	UpdateEmail(id uint, newEmail string) error

	// Can return ErrNoRowsAffected
	UpdateHashedPassword(id uint, newHashedPassword string) error

	// Returns a valid user and can return ErrEmptySelection
	GetByID(id uint) (User, error)

	// Returns a valid user and can return ErrEmptySelection
	GetByEmail(email string) (User, error)
}

type UserService interface {
	// Returns the ID of the created user, can return ErrIncorrectParameters, ErrPasswordUnableToHash, ErrExistingEmail
	Create(email, password string) (uint, error)

	// Can return ErrNotExistingEntity
	Delete(id uint) error

	// Can return ErrNotExistingEntity, ErrIncorrectParameters
	UpdateEmail(id uint, email string) error

	// Can return ErrNotExistingEntity, ErrIncorrectParameters, ErrPasswordUnableToHash
	UpdatePassword(id uint, password string) error

	// Returns a valid user, can return ErrIncorrectParameters, ErrNotExistingEntity
	GetByID(id uint) (User, error)

	// Returns a valid user, can return ErrIncorrectParameters, ErrNotExistingEntity
	GetByEmail(email string) (User, error)

	// Returns nil if credentials are OK, can return ErrIncorrectParameters, ErrNotValidCredentials, ErrNotExistingEntity
	CheckCredentials(email, password string) error
}
