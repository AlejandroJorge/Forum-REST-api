package domain

import (
	"time"

	"github.com/AlejandroJorge/forum-rest-api/util"
)

type Post struct {
	PostID       uint      `json:"PostID"`
	OwnerID      uint      `json:"OwnerID"`
	Title        string    `json:"Title"`
	Description  string    `json:"Description"`
	Content      string    `json:"Content"`
	CreationDate time.Time `json:"CreationDate"`
	Likes        uint      `json:"Likes"`
}

func (p Post) Validate() bool {
	conditions := []bool{
		p.PostID != 0,
		p.OwnerID != 0,
		p.Title != "",
		p.Description != "",
		p.Content != "",
		!p.CreationDate.IsZero(),
	}

	return util.MergeAND(conditions)
}

type PostRepository interface {
	CreateNew(ownerID uint, title, description, content string) (uint, error)

	Delete(id uint) error

	UpdateTitle(id uint, newTitle string) error

	UpdateDescription(id uint, newDescription string) error

	UpdateContent(id uint, newContent string) error

	GetByID(id uint) (Post, error)

	GetByUser(userId uint) ([]Post, error)

	GetPopularAfter(moment time.Time, amount uint) ([]Post, error)

	AddLike(userId uint, postId uint) error

	DeleteLike(userId uint, postId uint) error
}

type PostService interface {
	Create(ownerID uint, title, description, content string) (uint, error)

	Delete(id uint) error

	UpdateTitle(id uint, title string) error

	UpdateDescription(id uint, description string) error

	UpdateContent(id uint, content string) error

	GetByID(id uint) (Post, error)

	GetByUser(userId uint) ([]Post, error)

	GetPopularToday() ([]Post, error)

	GetPopularLastWeek() ([]Post, error)

	GetPopularLastMonth() ([]Post, error)

	GetPopularAllTime() ([]Post, error)

	AddLike(userId uint, postId uint) error

	DeleteLike(userId uint, postId uint) error
}
