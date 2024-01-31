package service

import (
	"github.com/AlejandroJorge/forum-rest-api/domain"
	"github.com/AlejandroJorge/forum-rest-api/logging"
	"github.com/AlejandroJorge/forum-rest-api/repository"
	"github.com/AlejandroJorge/forum-rest-api/util"
)

type profileServiceImpl struct {
	repo domain.ProfileRepository
}

// Can return ErrAlreadyExisting, ErrNotExistingEntity,
func (serv profileServiceImpl) AddFollow(followerId uint, followedId uint) error {
	if followedId == 0 || followerId == 0 {
		logging.LogDomainError(ErrIncorrectParameters)
		return ErrIncorrectParameters
	}

	err := serv.repo.AddFollow(followerId, followedId)
	if err == repository.ErrRepeatedEntity {
		logging.LogDomainError(ErrAlreadyExisting)
		return ErrAlreadyExisting
	}
	if err == repository.ErrNoMatchingDependency {
		logging.LogDomainError(ErrDependencyNotSatisfied)
		return ErrDependencyNotSatisfied
	}
	if err != nil {
		logging.LogUnexpectedRepositoryError(err)
		return ErrUnknown
	}

	return nil
}

// Returns the ID corresponding to the created profile, can return ErrDependencyNotSatisfied, ErrProfileExistsOrTagNameIsRepeated, ErrIncorrectParameters
func (serv profileServiceImpl) Create(userID uint, tagName, displayName string) (uint, error) {
	if userID == 0 || displayName == "" || !util.IsAlphanumeric(tagName) {
		logging.LogDomainError(ErrIncorrectParameters)
		return 0, ErrIncorrectParameters
	}

	id, err := serv.repo.Create(userID, tagName, displayName)
	if err == repository.ErrNoMatchingDependency {
		logging.LogDomainError(ErrDependencyNotSatisfied)
		return 0, ErrDependencyNotSatisfied
	}
	if err == repository.ErrRepeatedEntity {
		logging.LogDomainError(ErrProfileExistsOrTagNameIsRepeated)
		return 0, ErrProfileExistsOrTagNameIsRepeated
	}
	if err != nil {
		logging.LogUnexpectedDomainError(err)
		return 0, ErrUnknown
	}

	return id, nil
}

// Can return ErrIncorrectParameters, ErrNotExistingEntity
func (serv profileServiceImpl) Delete(id uint) error {
	if id == 0 {
		logging.LogDomainError(ErrIncorrectParameters)
		return ErrIncorrectParameters
	}

	err := serv.repo.Delete(id)
	if err == repository.ErrNoRowsAffected {
		logging.LogDomainError(ErrNotExistingEntity)
		return ErrNotExistingEntity
	}
	if err != nil {
		logging.LogUnexpectedDomainError(err)
		return ErrUnknown
	}

	return nil
}

// Can return ErrIncorrectParameters, ErrNotExistingEntity
func (serv profileServiceImpl) DeleteFollow(followerId uint, followedId uint) error {
	if followedId == 0 || followerId == 0 {
		logging.LogDomainError(ErrIncorrectParameters)
		return ErrIncorrectParameters
	}

	err := serv.repo.DeleteFollow(followerId, followedId)
	if err == repository.ErrNoRowsAffected {
		logging.LogDomainError(ErrNotExistingEntity)
		return ErrNotExistingEntity
	}
	if err != nil {
		logging.LogUnexpectedDomainError(err)
		return ErrUnknown
	}

	return nil
}

// Returns a valid profile, can return ErrIncorrectParameters, ErrNotExistingEntity
func (serv profileServiceImpl) GetByTagName(tagName string) (domain.Profile, error) {
	if !util.IsAlphanumeric(tagName) {
		logging.LogDomainError(ErrIncorrectParameters)
		return domain.Profile{}, ErrIncorrectParameters
	}

	profile, err := serv.repo.GetByTagName(tagName)
	if err == repository.ErrEmptySelection {
		logging.LogDomainError(ErrNotExistingEntity)
		return domain.Profile{}, ErrNotExistingEntity
	}
	if err != nil {
		logging.LogUnexpectedDomainError(err)
		return domain.Profile{}, ErrUnknown
	}

	return profile, nil
}

// Returns a valid profile, can return ErrIncorrectParameters, ErrNotExistingEntity
func (serv profileServiceImpl) GetByUserID(userId uint) (domain.Profile, error) {
	if userId == 0 {
		logging.LogDomainError(ErrIncorrectParameters)
		return domain.Profile{}, ErrIncorrectParameters
	}

	profile, err := serv.repo.GetByUserID(userId)
	if err == repository.ErrEmptySelection {
		logging.LogDomainError(ErrNotExistingEntity)
		return domain.Profile{}, ErrNotExistingEntity
	}
	if err != nil {
		logging.LogUnexpectedDomainError(err)
		return domain.Profile{}, ErrUnknown
	}

	return profile, nil
}

// Returns a slice of valid profiles, can return ErrIncorrectParameters, ErrNotExistingEntity
func (serv profileServiceImpl) GetFollowersByID(userId uint) ([]domain.Profile, error) {
	if userId == 0 {
		logging.LogDomainError(ErrIncorrectParameters)
		return nil, ErrIncorrectParameters
	}

	profiles, err := serv.repo.GetFollowersByID(userId)
	if err == repository.ErrEmptySelection {
		logging.LogDomainError(ErrNotExistingEntity)
		return nil, ErrNotExistingEntity
	}

	if err != nil {
		logging.LogUnexpectedDomainError(err)
		return nil, ErrUnknown
	}

	return profiles, nil
}

// Returns a slice of valid profiles, can return ErrIncorrectParameters, ErrNotExistingEntity
func (serv profileServiceImpl) GetFollowersByTagName(tagName string) ([]domain.Profile, error) {
	if !util.IsAlphanumeric(tagName) {
		logging.LogDomainError(ErrIncorrectParameters)
		return nil, ErrIncorrectParameters
	}

	profiles, err := serv.repo.GetFollowersByTagName(tagName)
	if err == repository.ErrEmptySelection {
		logging.LogDomainError(ErrNotExistingEntity)
		return nil, ErrNotExistingEntity
	}

	if err != nil {
		logging.LogUnexpectedDomainError(err)
		return nil, ErrUnknown
	}

	return profiles, nil
}

// Returns a slice of valid profiles, can return ErrIncorrectParameters, ErrNotExistingEntity
func (serv profileServiceImpl) GetFollowsByID(userId uint) ([]domain.Profile, error) {
	if userId == 0 {
		logging.LogDomainError(ErrIncorrectParameters)
		return nil, ErrIncorrectParameters
	}

	profiles, err := serv.repo.GetFollowsByID(userId)
	if err == repository.ErrEmptySelection {
		logging.LogDomainError(ErrNotExistingEntity)
		return nil, ErrNotExistingEntity
	}

	if err != nil {
		logging.LogUnexpectedDomainError(err)
		return nil, ErrUnknown
	}

	return profiles, nil
}

// Returns a slice of valid profiles, can return ErrIncorrectParameters, ErrNotExistingEntity
func (serv profileServiceImpl) GetFollowsByTagName(tagName string) ([]domain.Profile, error) {
	if !util.IsAlphanumeric(tagName) {
		logging.LogDomainError(ErrIncorrectParameters)
		return nil, ErrIncorrectParameters
	}

	profiles, err := serv.repo.GetFollowsByTagName(tagName)
	if err == repository.ErrEmptySelection {
		logging.LogDomainError(ErrNotExistingEntity)
		return nil, ErrNotExistingEntity
	}

	if err != nil {
		logging.LogUnexpectedDomainError(err)
		return nil, ErrUnknown
	}

	return profiles, nil
}

// Can return ErrNotExistingEntity
func (serv profileServiceImpl) UpdateTagName(id uint, tagName string) error {
	if id == 0 || !util.IsAlphanumeric(tagName) {
		logging.LogDomainError(ErrIncorrectParameters)
		return ErrIncorrectParameters
	}

	err := serv.repo.UpdateTagName(id, tagName)
	if err == repository.ErrNoRowsAffected {
		logging.LogDomainError(ErrNotExistingEntity)
		return ErrNotExistingEntity
	}
	if err != nil {
		logging.LogUnexpectedDomainError(err)
		return ErrUnknown
	}

	return nil
}

// Can return ErrNotExistingEntity
func (serv profileServiceImpl) UpdateDisplayName(id uint, displayName string) error {
	if id == 0 || displayName == "" {
		logging.LogDomainError(ErrIncorrectParameters)
		return ErrIncorrectParameters
	}

	err := serv.repo.UpdateDisplayName(id, displayName)
	if err == repository.ErrNoRowsAffected {
		logging.LogDomainError(ErrNotExistingEntity)
		return ErrNotExistingEntity
	}
	if err != nil {
		logging.LogUnexpectedDomainError(err)
		return ErrUnknown
	}

	return nil
}

// Can return ErrNotExistingEntity
func (serv profileServiceImpl) UpdatePicturePath(id uint, picturePath string) error {
	if id == 0 || picturePath == "" {
		logging.LogDomainError(ErrIncorrectParameters)
		return ErrIncorrectParameters
	}

	err := serv.repo.UpdatePicturePath(id, picturePath)
	if err == repository.ErrNoRowsAffected {
		logging.LogDomainError(ErrNotExistingEntity)
		return ErrNotExistingEntity
	}
	if err != nil {
		logging.LogUnexpectedDomainError(err)
		return ErrUnknown
	}

	return nil
}

// Can return ErrNotExistingEntity
func (serv profileServiceImpl) UpdateBackgroundPath(id uint, backgroundPath string) error {
	if id == 0 || backgroundPath == "" {
		logging.LogDomainError(ErrIncorrectParameters)
		return ErrIncorrectParameters
	}

	err := serv.repo.UpdateBackgroundPath(id, backgroundPath)
	if err == repository.ErrNoRowsAffected {
		logging.LogDomainError(ErrNotExistingEntity)
		return ErrNotExistingEntity
	}
	if err != nil {
		logging.LogUnexpectedDomainError(err)
		return ErrUnknown
	}

	return nil
}

func NewProfileService(repo domain.ProfileRepository) domain.ProfileService {
	return profileServiceImpl{repo: repo}
}
