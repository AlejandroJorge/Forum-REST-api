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
	Create(email, hashedPassword string) (uint, error)

	Delete(id uint) error

	UpdateEmail(id uint, newEmail string) error

	UpdateHashedPassword(id uint, newHashedPassword string) error

	GetByID(id uint) (User, error)

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
