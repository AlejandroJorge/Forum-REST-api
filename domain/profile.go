package domain

import "github.com/AlejandroJorge/forum-rest-api/util"

type Profile struct {
	UserID         uint   `json:"UserID"`
	DisplayName    string `json:"DisplayName"`
	TagName        string `json:"TagName"`
	PicturePath    string `json:"PicturePath"`
	BackgroundPath string `json:"BackgroundPath"`
	Followers      uint   `json:"Followers"`
	Follows        uint   `json:"Follows"`
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
	// Returns the id of the created profile, can return ErrNoMatchingDependency, ErrRepeatedEntity
	Create(userID uint, tagName, displayName string) (uint, error)

	// Can return ErrNoRowsAffected
	Delete(id uint) error

	// Can return ErrNoRowsAffected, ErrRepeatedEntity
	UpdateTagName(id uint, newTagName string) error

	// Can return ErrNoRowsAffected
	UpdateDisplayName(id uint, newDisplayName string) error

	// Can return ErrNoRowsAffected
	UpdatePicturePath(id uint, newPicturePath string) error

	// Can return ErrNoRowsAffected
	UpdateBackgroundPath(id uint, newBackgroundPath string) error

	// Returns a valid profile and can return ErrEmptySelection
	GetByUserID(userId uint) (Profile, error)

	// Returns a valid profile and can return ErrEmptySelection
	GetByTagName(tagName string) (Profile, error)

	// Returns an slice of valid profiles, can return ErrEmptySelection
	GetFollowersByID(userId uint) ([]Profile, error)

	// Returns an slice of valid profiles, can return ErrEmptySelection
	GetFollowersByTagName(tagName string) ([]Profile, error)

	// Returns an slice of valid profiles, can return ErrEmptySelection
	GetFollowsByID(userId uint) ([]Profile, error)

	// Returns an slice of valid profiles, can return ErrEmptySelection
	GetFollowsByTagName(tagName string) ([]Profile, error)

	// Can return ErrRepeatedEntity, ErrNoMatchingDependency
	AddFollow(followerId uint, followedId uint) error

	// Can return ErrNoRowsAffected
	DeleteFollow(followerId uint, followedId uint) error
}

type ProfileService interface {
	// Returns the ID corresponding to the created profile, can return ErrDependencyNotSatisfied, ErrProfileExistsOrTagNameIsRepeated, ErrIncorrectParameters
	Create(userID uint, displayName, tagName string) (uint, error)

	// Can return ErrIncorrectParameters, ErrNotExistingEntity
	Delete(id uint) error

	// Can return ErrNotExistingEntity, ErrAlreadyExisting
	UpdateTagName(id uint, tagName string) error

	// Can return ErrNotExistingEntity
	UpdateDisplayName(id uint, displayName string) error

	// Can return ErrNotExistingEntity
	UpdatePicturePath(id uint, picturePath string) error

	// Can return ErrNotExistingEntity
	UpdateBackgroundPath(id uint, backgroundPath string) error

	// Returns a valid profile, can return ErrIncorrectParameters, ErrNotExistingEntity
	GetByUserID(userId uint) (Profile, error)

	// Returns a valid profile, can return ErrIncorrectParameters, ErrNotExistingEntity
	GetByTagName(tagName string) (Profile, error)

	// Returns a slice of valid profiles, can return ErrIncorrectParameters, ErrNotExistingEntity
	GetFollowersByID(userId uint) ([]Profile, error)

	// Returns a slice of valid profiles, can return ErrIncorrectParameters, ErrNotExistingEntity
	GetFollowersByTagName(tagName string) ([]Profile, error)

	// Returns a slice of valid profiles, can return ErrIncorrectParameters, ErrNotExistingEntity
	GetFollowsByID(userId uint) ([]Profile, error)

	// Returns a slice of valid profiles, can return ErrIncorrectParameters, ErrNotExistingEntity
	GetFollowsByTagName(tagName string) ([]Profile, error)

	// Can return ErrAlreadyExisting, ErrIncorrectParameters, ErrDependencyNotSatisfied
	AddFollow(followerId uint, followedId uint) error

	// Can return ErrIncorrectParameters, ErrNotExistingEntity
	DeleteFollow(followerId uint, followedId uint) error
}
