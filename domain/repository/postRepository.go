package repository

import (
	"time"

	"github.com/AlejandroJorge/forum-rest-api/data/model"
)

type PostRepository interface {
	// Returns the post corresponding to the provided id
	GetByID(id uint) (model.Post, error)

	// Returns the posts corresponding to the provided userID, they're sorted by likes
	GetByUser(userId uint) ([]model.Post, error)

	// Returns an amount of posts that occured after certain moment, they're sorted by likes
	GetPopularAfter(moment time.Time, amount uint) ([]model.Post, error)

	// Creates a new user, the id in the model is ignored
	CreateNew(post model.Post) error

	// Updates the title of the post corresponding to the provided id
	UpdateTitle(id uint, newTitle string) error

	// Updates the description of the post corresponding to the provided id
	UpdateDescription(id uint, newDescription string) error

	// Updates the content of the post corresponding to the provided id
	UpdateContent(id uint, newContent string) error

	// Deletes the post corresponding to the provided ID
	Delete(id uint) error
}
