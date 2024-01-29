package service

import (
	"github.com/AlejandroJorge/forum-rest-api/domain"
	"github.com/AlejandroJorge/forum-rest-api/util"
)

type commentServiceImpl struct {
	repo domain.CommentRepository
}

func (serv commentServiceImpl) AddLike(userId uint, commentId uint) error {
	if userId == 0 || commentId == 0 {
		return util.ErrIncorrectParameters
	}

	return serv.repo.AddLike(userId, commentId)
}

func (serv commentServiceImpl) CreateNew(createInfo struct {
	UserID  uint
	PostID  uint
	Content string
}) (uint, error) {
	if createInfo.UserID == 0 ||
		createInfo.PostID == 0 ||
		createInfo.Content == "" {
		return 0, util.ErrIncorrectParameters
	}

	return serv.repo.CreateNew(domain.Comment{
		UserID:  createInfo.UserID,
		PostID:  createInfo.PostID,
		Content: createInfo.Content,
	})
}

func (serv commentServiceImpl) Delete(id uint) error {
	if id == 0 {
		return util.ErrIncorrectParameters
	}

	return serv.repo.Delete(id)
}

func (serv commentServiceImpl) DeleteLike(userId uint, commentId uint) error {
	if userId == 0 || commentId == 0 {
		return util.ErrIncorrectParameters
	}

	return serv.repo.DeleteLike(userId, commentId)
}

func (serv commentServiceImpl) GetByID(id uint) (domain.Comment, error) {
	if id == 0 {
		return domain.Comment{}, util.ErrIncorrectParameters
	}

	return serv.repo.GetByID(id)
}

func (serv commentServiceImpl) GetByPost(postID uint) ([]domain.Comment, error) {
	if postID == 0 {
		return nil, util.ErrIncorrectParameters
	}

	return serv.repo.GetByPost(postID)
}

func (serv commentServiceImpl) GetByUser(userID uint) ([]domain.Comment, error) {
	if userID == 0 {
		return nil, util.ErrIncorrectParameters
	}

	return serv.repo.GetByUser(userID)
}

func (serv commentServiceImpl) Update(id uint, updatedContent string) error {
	if id == 0 || updatedContent == "" {
		return util.ErrIncorrectParameters
	}

	return serv.repo.UpdateContent(id, updatedContent)
}

func NewCommentService(repo domain.CommentRepository) domain.CommentService {
	return commentServiceImpl{repo: repo}
}
