package service

import (
	"time"

	"github.com/AlejandroJorge/forum-rest-api/domain"
	"github.com/AlejandroJorge/forum-rest-api/util"
)

type postServiceImpl struct {
	repo domain.PostRepository
}

func (serv postServiceImpl) AddLike(userId uint, postId uint) error {
	if userId == 0 || postId == 0 {
		return util.ErrIncorrectParameters
	}

	return serv.repo.AddLike(userId, postId)
}

func (serv postServiceImpl) CreateNew(createInfo struct {
	OwnerID     uint
	Title       string
	Description string
	Content     string
}) (uint, error) {
	if createInfo.OwnerID == 0 ||
		createInfo.Title == "" ||
		createInfo.Description == "" ||
		createInfo.Content == "" {
		return 0, util.ErrIncorrectParameters
	}

	return serv.repo.CreateNew(domain.Post{
		OwnerID:     createInfo.OwnerID,
		Title:       createInfo.Title,
		Description: createInfo.Description,
		Content:     createInfo.Content,
	})
}

func (serv postServiceImpl) Delete(id uint) error {
	if id == 0 {
		return util.ErrIncorrectParameters
	}

	return serv.repo.Delete(id)
}

func (serv postServiceImpl) DeleteLike(userId uint, postId uint) error {
	if userId == 0 || postId == 0 {
		return util.ErrIncorrectParameters
	}

	return serv.repo.DeleteLike(userId, postId)
}

func (serv postServiceImpl) GetByID(id uint) (domain.Post, error) {
	if id == 0 {
		return domain.Post{}, util.ErrIncorrectParameters
	}

	return serv.repo.GetByID(id)
}

func (serv postServiceImpl) GetByUser(userId uint) ([]domain.Post, error) {
	if userId == 0 {
		return nil, util.ErrIncorrectParameters
	}

	return serv.repo.GetByUser(userId)
}

func (serv postServiceImpl) GetPopularAllTime() ([]domain.Post, error) {
	return serv.repo.GetPopularAfter(time.Time{}, 20)
}

func (serv postServiceImpl) GetPopularLastMonth() ([]domain.Post, error) {
	return serv.repo.GetPopularAfter(time.Now().AddDate(0, -1, 0), 20)
}

func (serv postServiceImpl) GetPopularLastWeek() ([]domain.Post, error) {
	return serv.repo.GetPopularAfter(time.Now().AddDate(0, 0, -7), 20)
}

func (serv postServiceImpl) GetPopularToday() ([]domain.Post, error) {
	return serv.repo.GetPopularAfter(time.Now().AddDate(0, 0, -1), 20)
}

func (serv postServiceImpl) Update(id uint, updateInfo struct {
	UpdatedTitle       string
	UpdatedDescription string
	UpdatedContent     string
}) error {
	if id == 0 {
		return util.ErrIncorrectParameters
	}

	if updateInfo.UpdatedTitle != "" {
		err := serv.repo.UpdateTitle(id, updateInfo.UpdatedTitle)
		if err != nil {
			return err
		}
	}

	if updateInfo.UpdatedDescription != "" {
		err := serv.repo.UpdateDescription(id, updateInfo.UpdatedDescription)
		if err != nil {
			return err
		}
	}

	if updateInfo.UpdatedContent != "" {
		err := serv.repo.UpdateContent(id, updateInfo.UpdatedContent)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewPostService(repo domain.PostRepository) domain.PostService {
	return postServiceImpl{repo: repo}
}
