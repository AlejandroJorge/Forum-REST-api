package service

import (
	"time"

	"github.com/AlejandroJorge/forum-rest-api/domain"
	"github.com/AlejandroJorge/forum-rest-api/logging"
	"github.com/AlejandroJorge/forum-rest-api/repository"
)

type postServiceImpl struct {
	repo domain.PostRepository
}

// Can return ErIncorrectParameters, ErrAlreadyExisting, ErrDependencyNotSatisfied
func (serv postServiceImpl) AddLike(userId uint, postId uint) error {
	if userId == 0 || postId == 0 {
		logging.LogDomainError(ErrIncorrectParameters)
		return ErrIncorrectParameters
	}

	err := serv.repo.AddLike(userId, postId)
	if err == repository.ErrRepeatedEntity {
		logging.LogDomainError(ErrAlreadyExisting)
		return ErrAlreadyExisting
	}
	if err == repository.ErrNoMatchingDependency {
		logging.LogDomainError(ErrDependencyNotSatisfied)
		return ErrDependencyNotSatisfied
	}
	if err != nil {
		logging.LogUnexpectedDomainError(err)
		return ErrUnknown
	}

	return nil
}

// Returns the ID of the created post, can return ErrIncorrectParameters, ErrDependencyNotSatisfied, ErrAlreadyExisting
func (serv postServiceImpl) Create(ownerID uint, title, description, content string) (uint, error) {
	if ownerID == 0 || title == "" || description == "" || content == "" {
		logging.LogDomainError(ErrIncorrectParameters)
		return 0, ErrIncorrectParameters
	}

	id, err := serv.repo.Create(ownerID, title, description, content)
	if err == repository.ErrNoMatchingDependency {
		logging.LogDomainError(ErrDependencyNotSatisfied)
		return 0, ErrDependencyNotSatisfied
	}
	if err == repository.ErrRepeatedEntity {
		logging.LogDomainError(ErrAlreadyExisting)
		return 0, ErrAlreadyExisting
	}
	if err != nil {
		logging.LogUnexpectedDomainError(err)
		return 0, ErrUnknown
	}

	return id, nil
}

// Can return ErrIncorrectParameters, ErrNotExistingEntity
func (serv postServiceImpl) Delete(id uint) error {
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
func (serv postServiceImpl) DeleteLike(userId uint, postId uint) error {
	if userId == 0 || postId == 0 {
		logging.LogDomainError(ErrIncorrectParameters)
		return ErrIncorrectParameters
	}

	err := serv.repo.DeleteLike(userId, postId)
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

// Returns a valid post, can return ErrIncorrectParameters, ErrNotExistingEntity
func (serv postServiceImpl) GetByID(id uint) (domain.Post, error) {
	if id == 0 {
		logging.LogDomainError(ErrIncorrectParameters)
		return domain.Post{}, ErrIncorrectParameters
	}

	post, err := serv.repo.GetByID(id)
	if err == repository.ErrEmptySelection {
		logging.LogDomainError(ErrNotExistingEntity)
		return domain.Post{}, ErrNotExistingEntity
	}
	if err != nil {
		logging.LogUnexpectedDomainError(err)
		return domain.Post{}, ErrUnknown
	}

	return post, nil
}

// Returns a slice of valid posts, can return ErrIncorrectParameters, ErrNotExistingEntity
func (serv postServiceImpl) GetByUser(userId uint) ([]domain.Post, error) {
	if userId == 0 {
		logging.LogDomainError(ErrIncorrectParameters)
		return nil, ErrIncorrectParameters
	}

	posts, err := serv.repo.GetByUser(userId)
	if err == repository.ErrEmptySelection {
		logging.LogDomainError(ErrNotExistingEntity)
		return nil, ErrNotExistingEntity
	}
	if err != nil {
		logging.LogUnexpectedDomainError(err)
		return nil, ErrUnknown
	}

	return posts, nil
}

// Returns a slice of valid posts, can return ErrNotExistingEntity
func (serv postServiceImpl) GetPopularAllTime() ([]domain.Post, error) {
	posts, err := serv.repo.GetPopularAfter(time.Time{}, 20)
	if err == repository.ErrEmptySelection {
		logging.LogDomainError(ErrNotExistingEntity)
		return nil, ErrNotExistingEntity
	}
	if err != nil {
		logging.LogUnexpectedDomainError(err)
		return nil, ErrUnknown
	}

	return posts, nil
}

// Returns a slice of valid posts, can return ErrNotExistingEntity
func (serv postServiceImpl) GetPopularLastMonth() ([]domain.Post, error) {
	posts, err := serv.repo.GetPopularAfter(time.Now().AddDate(0, -1, 0), 20)
	if err == repository.ErrEmptySelection {
		logging.LogDomainError(ErrNotExistingEntity)
		return nil, ErrNotExistingEntity
	}
	if err != nil {
		logging.LogUnexpectedDomainError(err)
		return nil, ErrUnknown
	}

	return posts, nil
}

// Returns a slice of valid posts, can return ErrNotExistingEntity
func (serv postServiceImpl) GetPopularLastWeek() ([]domain.Post, error) {
	posts, err := serv.repo.GetPopularAfter(time.Now().AddDate(0, 0, -7), 20)
	if err == repository.ErrEmptySelection {
		logging.LogDomainError(ErrNotExistingEntity)
		return nil, ErrNotExistingEntity
	}
	if err != nil {
		logging.LogUnexpectedDomainError(err)
		return nil, ErrUnknown
	}

	return posts, nil
}

// Returns a slice of valid posts, can return ErrNotExistingEntity
func (serv postServiceImpl) GetPopularToday() ([]domain.Post, error) {
	posts, err := serv.repo.GetPopularAfter(time.Now().AddDate(0, 0, -1), 20)
	if err == repository.ErrEmptySelection {
		logging.LogDomainError(ErrNotExistingEntity)
		return nil, ErrNotExistingEntity
	}
	if err != nil {
		logging.LogUnexpectedDomainError(err)
		return nil, ErrUnknown
	}

	return posts, nil
}

// Can return ErrIncorrectParameters, ErrNotExistingEntity,
func (serv postServiceImpl) UpdateTitle(id uint, title string) error {
	if id == 0 || title == "" {
		logging.LogDomainError(ErrIncorrectParameters)
		return ErrIncorrectParameters
	}

	err := serv.repo.UpdateTitle(id, title)
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

// Can return ErrIncorrectParameters, ErrNotExistingEntity,
func (serv postServiceImpl) UpdateDescription(id uint, description string) error {
	if id == 0 || description == "" {
		logging.LogDomainError(ErrIncorrectParameters)
		return ErrIncorrectParameters
	}

	err := serv.repo.UpdateDescription(id, description)
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

// Can return ErrIncorrectParameters, ErrNotExistingEntity,
func (serv postServiceImpl) UpdateContent(id uint, content string) error {
	if id == 0 || content == "" {
		logging.LogDomainError(ErrIncorrectParameters)
		return ErrIncorrectParameters
	}

	err := serv.repo.UpdateDescription(id, content)
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

func NewPostService(repo domain.PostRepository) domain.PostService {
	return postServiceImpl{repo: repo}
}
