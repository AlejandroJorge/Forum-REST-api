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
	// Returns the post corresponding to the provided id
	GetByID(id uint) (Post, error)

	// Returns the posts corresponding to the provided userID, they're sorted by likes
	GetByUser(userId uint) ([]Post, error)

	// Returns an amount of posts that occured after certain moment, they're sorted by likes
	GetPopularAfter(moment time.Time, amount uint) ([]Post, error)

	// Creates a new user, the id in the model is ignored
	CreateNew(post Post) (uint, error)

	// Updates the title of the post corresponding to the provided id
	UpdateTitle(id uint, newTitle string) error

	// Updates the description of the post corresponding to the provided id
	UpdateDescription(id uint, newDescription string) error

	// Updates the content of the post corresponding to the provided id
	UpdateContent(id uint, newContent string) error

	// Creates the relation of liking between a profile and a post
	AddLike(userId uint, postId uint) error

	// Deletes the relation of liking between a profile and a post
	DeleteLike(userId uint, postId uint) error

	// Deletes the post corresponding to the provided ID
	Delete(id uint) error
}

type PostService interface {
	// Returns the post corresponding to the provided id
	GetByID(id uint) (Post, error)

	// Returns the posts corresponding to the provided userID, they're sorted by likes
	GetByUser(userId uint) ([]Post, error)

	GetPopularToday() ([]Post, error)

	GetPopularLastWeek() ([]Post, error)

	GetPopularLastMonth() ([]Post, error)

	GetPopularAllTime() ([]Post, error)

	// Creates a new user, the id in the model is ignored
	CreateNew(createInfo struct {
		OwnerID     uint
		Title       string
		Description string
		Content     string
	}) (uint, error)

	//
	Update(id uint, updateInfo struct {
		UpdatedTitle       string
		UpdatedDescription string
		UpdatedContent     string
	}) error

	// Creates the relation of liking between a profile and a post
	AddLike(userId uint, postId uint) error

	// Deletes the relation of liking between a profile and a post
	DeleteLike(userId uint, postId uint) error

	// Deletes the post corresponding to the provided ID
	Delete(id uint) error
}
