package domain

type Profile struct {
	UserID         uint
	DisplayName    string
	TagName        string
	PicturePath    string
	BackgroundPath string
	Followers      uint
	Follows        uint
}

type ProfileRepository interface {
	// Returns the profile corresponding to the provided userID
	GetByUserID(userId uint) (Profile, error)

	// Returns the profile corresponding to the provided tagName
	GetByTagName(tagName string) (Profile, error)

	// Creates a new profile, the id in the model should correspond to a valid user
	CreateNew(profile Profile) (uint, error)

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
