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
	// Returns the id of the created comment, can return ErrNoMatchingDependency
	Create(postID, userID uint, content string) (uint, error)

	// Can return ErrNoRowsAffected
	Delete(id uint) error

	// Can return ErrNoRowsAffected
	UpdateContent(id uint, newContent string) error

	// Returns a valid comment and can return ErrEmptySelection
	GetByID(id uint) (Comment, error)

	// Returns an slice of valid comments, can return ErrEmptySelection
	GetByPost(postID uint) ([]Comment, error)

	// Returns an slice of valid comments, can return ErrEmptySelection
	GetByUser(userID uint) ([]Comment, error)

	// Can return ErrNoMatchingDependency
	AddLike(userId uint, commentId uint) error

	// Can return ErrNoRowsAffected
	DeleteLike(userId uint, commentId uint) error
}

type CommentService interface {
	// Returns the ID of the generated comment, can return ErrIncorrectParameters, ErrDependencySatisfied
	Create(userID, postID uint, content string) (uint, error)

	// Can return ErrIncorrectParameters, ErrNotExistingEntity
	Delete(id uint) error

	// Can return ErrIncorrectParameters, ErrNotExistingEntity
	Update(id uint, updatedContent string) error

	// Returns a valid comment, can return ErrIncorrectParameters, ErrNotExistingEntity
	GetByID(id uint) (Comment, error)

	// Returns a slice of valid comments, can return ErrIncorrectParameters, ErrNotExistingEntity
	GetByPost(postID uint) ([]Comment, error)

	// Returns a slice of valid comments, can return ErrIncorrectParameters, ErrNotExistingEntity
	GetByUser(userID uint) ([]Comment, error)

	// Can return ErrIncorrectParameters, ErrDependencyNotSatisfied
	AddLike(userId uint, commentId uint) error

	// Can return ErrIncorrectParameters, ErrNotExistingEntity
	DeleteLike(userId uint, commentId uint) error
}
