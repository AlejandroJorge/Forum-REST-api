package domain

import "github.com/AlejandroJorge/forum-rest-api/util"

type Profile struct {
	UserID         uint
	DisplayName    string
	TagName        string
	PicturePath    string
	BackgroundPath string
	Followers      uint
	Follows        uint
}

func (p Profile) Validate() bool {
	conditions := []bool{
		p.UserID != 0,
		p.DisplayName != "",
		p.TagName != "",
	}

	return util.MergeAND(conditions)
}

type ProfileRepository interface {
	// Returns the profile corresponding to the provided userID
	GetByUserID(userId uint) (Profile, error)

	// Returns the profile corresponding to the provided tagName
	GetByTagName(tagName string) (Profile, error)

	// Returns the profiles that follows the profile correpsonding to the provided ID
	GetFollowersByID(userId uint) ([]Profile, error)

	// Returns the profiles that follows the profile correpsonding to the provided tagName
	GetFollowersByTagName(tagName string) ([]Profile, error)

	// Returns the profiles that follows the profile correpsonding to the provided ID
	GetFollowsByID(userId uint) ([]Profile, error)

	// Returns the profiles that follows the profile correpsonding to the provided ID
	GetFollowsByTagName(tagName string) ([]Profile, error)

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

	// Creates the relation of following between a follower and a followed
	AddFollow(followerId uint, followedId uint) error

	// Deletes the relation of following between a follower and a followed
	DeleteFollow(followerId uint, followedId uint) error

	// Deletes the profile corresponding to the provided userID
	Delete(id uint) error
}

type ProfileService interface {
	// Returns the profile corresponding to the provided userID
	GetByUserID(userId uint) (Profile, error)

	// Returns the profile corresponding to the provided tagName
	GetByTagName(tagName string) (Profile, error)

	// Returns the profiles that follows the profile correpsonding to the provided ID
	GetFollowersByID(userId uint) ([]Profile, error)

	// Returns the profiles that follows the profile correpsonding to the provided tagName
	GetFollowersByTagName(tagName string) ([]Profile, error)

	// Returns the profiles that follows the profile correpsonding to the provided ID
	GetFollowsByID(userId uint) ([]Profile, error)

	// Returns the profiles that follows the profile correpsonding to the provided ID
	GetFollowsByTagName(tagName string) ([]Profile, error)

	// Creates a new profile, the id in the model should correspond to a valid user
	CreateNew(createInfo struct {
		UserID         uint
		DisplayName    string
		TagName        string
		PicturePath    string
		BackgroundPath string
	}) (uint, error)

	// Updates the profile with the info provided
	Update(id uint, updateInfo struct {
		UpdatedTagName        string
		UpdatedDisplayName    string
		UpdatedPicturePath    string
		UpdatedBackgroundPath string
	})

	// Creates the relation of following between a follower and a followed
	AddFollow(followerId uint, followedId uint) error

	// Deletes the relation of following between a follower and a followed
	DeleteFollow(followerId uint, followedId uint) error

	// Deletes the profile corresponding to the provided userID
	Delete(id uint) error
}
