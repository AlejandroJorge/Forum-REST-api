package repository

import "github.com/AlejandroJorge/forum-rest-api/data/model"

type ProfileRepository interface {
	// Returns the profile corresponding to the provided userID
	GetByUserID(userId uint) (model.Profile, error)

	// Returns the profile corresponding to the provided tagName
	GetByTagName(tagName string) (model.Profile, error)

	// Creates a new profile, the id in the model should correspond to a valid user
	CreateNew(profile model.Profile) error

	// Updates the tagName of the profile corresponding to the provided userID
	UpdateTagName(id uint, newTagName string) error

	// Updates the displayName of the profile corresponding to the provided userID
	UpdateDisplayName(id uint, newDisplayName string) error

	// Updates the picturePath of the profile corresponding to the provided userID
	UpdatePicturePath(id uint, newPicturePath string) error

	// Updates the backgroundPath of the profile corresponding to the provided userID
	UpdateBackgroundPath(id uint, newBackgroundPath string) error

	// Deletes the profile corresponding to the provided userID
	Delete(id uint) error
}
