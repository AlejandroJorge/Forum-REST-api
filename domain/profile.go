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

	// Can return ErrNoRowsAffected
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
	Create(userID uint, displayName, tagName, picturePath, backgroundPath string) (uint, error)

	Delete(id uint) error

	UpdateTagName(id uint, tagName string) error

	UpdateDisplayName(id uint, displayName string) error

	UpdatePicturePath(id uint, picturePath string) error

	UpdateBackgroundPath(id uint, backgroundPath string) error

	GetByUserID(userId uint) (Profile, error)

	GetByTagName(tagName string) (Profile, error)

	GetFollowersByID(userId uint) ([]Profile, error)

	GetFollowersByTagName(tagName string) ([]Profile, error)

	GetFollowsByID(userId uint) ([]Profile, error)

	GetFollowsByTagName(tagName string) ([]Profile, error)

	AddFollow(followerId uint, followedId uint) error

	DeleteFollow(followerId uint, followedId uint) error
}
