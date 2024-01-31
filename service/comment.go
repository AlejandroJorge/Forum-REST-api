package service

import (
	"github.com/AlejandroJorge/forum-rest-api/domain"
	"github.com/AlejandroJorge/forum-rest-api/logging"
	"github.com/AlejandroJorge/forum-rest-api/repository"
	"github.com/AlejandroJorge/forum-rest-api/util"
)

type commentServiceImpl struct {
	repo domain.CommentRepository
}

func (serv commentServiceImpl) AddLike(userId uint, commentId uint) error {
	if userId == 0 || commentId == 0 {
		logging.LogDomainError(ErrIncorrectParameters)
		return ErrIncorrectParameters
	}

	err := serv.repo.AddLike(userId, commentId)
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

func (serv commentServiceImpl) Create(userID, postID uint, content string) (uint, error) {
	if userID == 0 || postID == 0 || content == "" {
		logging.LogDomainError(ErrIncorrectParameters)
		return 0, ErrIncorrectParameters
	}

	id, err := serv.repo.Create(postID, userID, content)
	if err == repository.ErrNoMatchingDependency {
		logging.LogDomainError(ErrDependencyNotSatisfied)
		return 0, ErrDependencyNotSatisfied
	}
	if err != nil {
		logging.LogUnexpectedDomainError(err)
		return 0, ErrUnknown
	}

	return id, nil
}

func (serv commentServiceImpl) Delete(id uint) error {
	if id == 0 {
		logging.LogDomainError(ErrIncorrectParameters)
		return ErrIncorrectParameters
	}

	err := serv.repo.Delete(id)
	if err == repository.ErrNoRowsAffected {
		logging.LogDomainError(ErrNotExistingEntity)
		return ErrIncorrectParameters
	}
	if err != nil {
		logging.LogUnexpectedDomainError(err)
		return ErrUnknown
	}

	return nil
}

func (serv commentServiceImpl) DeleteLike(userId uint, commentId uint) error {
	if userId == 0 || commentId == 0 {
		logging.LogDomainError(ErrIncorrectParameters)
		return ErrIncorrectParameters
	}

	err := serv.repo.DeleteLike(userId, commentId)
	if err == repository.ErrNoRowsAffected {
		logging.LogDomainError(ErrNotExistingEntity)
		return ErrIncorrectParameters
	}
	if err != nil {
		logging.LogUnexpectedDomainError(err)
		return ErrUnknown
	}

	return nil
}

func (serv commentServiceImpl) GetByID(id uint) (domain.Comment, error) {
	if id == 0 {
		logging.LogDomainError(ErrIncorrectParameters)
		return domain.Comment{}, ErrIncorrectParameters
	}

	comment, err := serv.repo.GetByID(id)
	if err == repository.ErrEmptySelection {
		logging.LogDomainError(ErrNotExistingEntity)
		return domain.Comment{}, ErrNotExistingEntity
	}
	if err != nil {
		logging.LogUnexpectedDomainError(err)
		return domain.Comment{}, ErrUnknown
	}

	return comment, nil
}

func (serv commentServiceImpl) GetByPost(postID uint) ([]domain.Comment, error) {
	if postID == 0 {
		logging.LogDomainError(ErrIncorrectParameters)
		return nil, ErrIncorrectParameters
	}

	comments, err := serv.repo.GetByPost(postID)
	if err == repository.ErrEmptySelection {
		logging.LogDomainError(ErrNotExistingEntity)
		return nil, ErrNotExistingEntity
	}
	if err != nil {
		logging.LogUnexpectedDomainError(err)
		return nil, ErrUnknown
	}

	return comments, nil
}

func (serv commentServiceImpl) GetByUser(userID uint) ([]domain.Comment, error) {
	if userID == 0 {
		logging.LogDomainError(ErrIncorrectParameters)
		return nil, ErrIncorrectParameters
	}

	comments, err := serv.repo.GetByUser(userID)
	if err == repository.ErrEmptySelection {
		logging.LogDomainError(ErrNotExistingEntity)
		return nil, ErrNotExistingEntity
	}
	if err != nil {
		logging.LogUnexpectedDomainError(err)
		return nil, ErrUnknown
	}

	return comments, nil
}

func (serv commentServiceImpl) Update(id uint, updatedContent string) error {
	if id == 0 || updatedContent == "" {
		logging.LogDomainError(ErrIncorrectParameters)
		return ErrIncorrectParameters
	}

	err := serv.repo.UpdateContent(id, updatedContent)
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

func NewCommentService(repo domain.CommentRepository) domain.CommentService {
	return commentServiceImpl{repo: repo}
}
