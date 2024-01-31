package domain

import "github.com/AlejandroJorge/forum-rest-api/util"

type Comment struct {
	ID      uint   `json:"ID"`
	PostID  uint   `json:"PostID"`
	UserID  uint   `json:"UserID"`
	Content string `json:"Content"`
	Likes   uint   `json:"Likes"`
}

func (c Comment) Validate() bool {
	conditions := []bool{
		c.ID != 0,
		c.PostID != 0,
		c.UserID != 0,
		c.Content != "",
	}

	return util.MergeAND(conditions)
}

type CommentRepository interface {
	Create(postID, userID uint, content string) (uint, error)

	Delete(id uint) error

	UpdateContent(id uint, newContent string) error

	GetByID(id uint) (Comment, error)

	GetByPost(postID uint) ([]Comment, error)

	GetByUser(userID uint) ([]Comment, error)

	AddLike(userId uint, commentId uint) error

	DeleteLike(userId uint, commentId uint) error
}

type CommentService interface {
	CreateNew(userID, postID uint, content string) (uint, error)

	Delete(id uint) error

	Update(id uint, updatedContent string) error

	GetByID(id uint) (Comment, error)

	GetByPost(postID uint) ([]Comment, error)

	GetByUser(userID uint) ([]Comment, error)

	AddLike(userId uint, commentId uint) error

	DeleteLike(userId uint, commentId uint) error
}
