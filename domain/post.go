package domain

import "time"

type Post struct {
	PostID       uint
	OwnerID      uint
	Title        string
	Description  string
	Content      string
	CreationDate time.Time
	Likes        uint
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

	// Deletes the post corresponding to the provided ID
	Delete(id uint) error
}
