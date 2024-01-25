package repository

import "github.com/AlejandroJorge/forum-rest-api/data/model"

type UserRepository interface {
	// Returns the user corresponding to the provided ID
	GetByID(id uint) (model.User, error)

	// Returns the user corresponding to the provided Email
	GetByEmail(email string) (model.User, error)

	// Creates a new user, the id in the model is ignored
	CreateNew(user model.User) error

	// Updates the email of the user corresponding to the provided ID
	UpdateEmail(id uint, newEmail string) error

	// Updates the password of the user corresponding to the provided ID
	UpdatePassword(id uint, newEmail string) error

	// Deletes the user corresponding to the provided ID
	Delete(id uint) error
}
