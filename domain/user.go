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
	Create(email, password string) (uint, error)

	Delete(id uint) error

	UpdateEmail(id uint, email string) error

	UpdatePassword(id uint, password string) error

	GetByID(id uint) (User, error)

	GetByEmail(email string) (User, error)
}
