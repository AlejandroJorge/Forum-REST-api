package service

import (
	"github.com/AlejandroJorge/forum-rest-api/domain"
	"github.com/AlejandroJorge/forum-rest-api/util"
)

type profileServiceImpl struct {
	repo domain.ProfileRepository
}

func (serv profileServiceImpl) AddFollow(followerId uint, followedId uint) error {
	if followedId == 0 || followerId == 0 {
		return util.ErrIncorrectParameters
	}

	return serv.repo.AddFollow(followerId, followedId)
}

func (serv profileServiceImpl) CreateNew(createInfo struct {
	UserID         uint
	DisplayName    string
	TagName        string
	PicturePath    string
	BackgroundPath string
}) (uint, error) {
	if createInfo.UserID == 0 ||
		createInfo.DisplayName == "" ||
		util.IsAlphanumeric(createInfo.TagName) {
		return 0, util.ErrIncorrectParameters
	}

	return serv.repo.CreateNew(domain.Profile{
		UserID:         createInfo.UserID,
		DisplayName:    createInfo.DisplayName,
		TagName:        createInfo.TagName,
		PicturePath:    createInfo.PicturePath,
		BackgroundPath: createInfo.BackgroundPath,
	})
}

func (serv profileServiceImpl) Delete(id uint) error {
	if id == 0 {
		return util.ErrIncorrectParameters
	}

	return serv.Delete(id)
}

func (serv profileServiceImpl) DeleteFollow(followerId uint, followedId uint) error {
	if followedId == 0 || followerId == 0 {
		return util.ErrIncorrectParameters
	}

	return serv.repo.DeleteFollow(followerId, followedId)
}

func (serv profileServiceImpl) GetByTagName(tagName string) (domain.Profile, error) {
	if !util.IsAlphanumeric(tagName) {
		return domain.Profile{}, util.ErrIncorrectParameters
	}

	return serv.repo.GetByTagName(tagName)
}

func (serv profileServiceImpl) GetByUserID(userId uint) (domain.Profile, error) {
	if userId == 0 {
		return domain.Profile{}, util.ErrIncorrectParameters
	}

	return serv.repo.GetByUserID(userId)
}

func (serv profileServiceImpl) GetFollowersByID(userId uint) ([]domain.Profile, error) {
	if userId == 0 {
		return nil, util.ErrIncorrectParameters
	}

	return serv.repo.GetFollowersByID(userId)
}

func (serv profileServiceImpl) GetFollowersByTagName(tagName string) ([]domain.Profile, error) {
	if !util.IsAlphanumeric(tagName) {
		return nil, util.ErrIncorrectParameters
	}

	return serv.repo.GetFollowersByTagName(tagName)
}

func (serv profileServiceImpl) GetFollowsByID(userId uint) ([]domain.Profile, error) {
	if userId == 0 {
		return nil, util.ErrIncorrectParameters
	}

	return serv.repo.GetFollowsByID(userId)
}

func (serv profileServiceImpl) GetFollowsByTagName(tagName string) ([]domain.Profile, error) {
	if !util.IsAlphanumeric(tagName) {
		return nil, util.ErrIncorrectParameters
	}

	return serv.repo.GetFollowsByTagName(tagName)
}

func (serv profileServiceImpl) Update(id uint, updateInfo struct {
	UpdatedTagName        string
	UpdatedDisplayName    string
	UpdatedPicturePath    string
	UpdatedBackgroundPath string
}) error {
	if id == 0 {
		return util.ErrIncorrectParameters
	}

	if updateInfo.UpdatedTagName != "" {
		if !util.IsAlphanumeric(updateInfo.UpdatedTagName) {
			return util.ErrIncorrectParameters
		}

		err := serv.repo.UpdateTagName(id, updateInfo.UpdatedTagName)
		if err != nil {
			return err
		}
	}

	if updateInfo.UpdatedDisplayName != "" {
		err := serv.repo.UpdateDisplayName(id, updateInfo.UpdatedDisplayName)
		if err != nil {
			return err
		}
	}

	if updateInfo.UpdatedPicturePath != "" {
		err := serv.repo.UpdatePicturePath(id, updateInfo.UpdatedPicturePath)
		if err != nil {
			return err
		}
	}

	if updateInfo.UpdatedBackgroundPath != "" {
		err := serv.repo.UpdateBackgroundPath(id, updateInfo.UpdatedBackgroundPath)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewProfileService(repo domain.ProfileRepository) domain.ProfileService {
	return profileServiceImpl{repo: repo}
}
