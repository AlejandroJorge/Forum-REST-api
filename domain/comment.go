package domain

type Comment struct {
	ID      uint
	PostID  uint
	UserID  uint
	Content string
	Likes   uint
}

type CommentRepository interface {
	// Returns the comment corresponding to the provided id
	GetByID(id uint) (Comment, error)

	// Returns the comments corresponding to the provided postID, they're sorted by likes
	GetByPost(postID uint) ([]Comment, error)

	// Returns the comments corresponding to the provided userID, they're sorted by likes
	GetByUser(userID uint) ([]Comment, error)

	// Creates a new comment, the id in the model is ignored
	CreateNew(comment Comment) (uint, error)

	// Updates the content of the comment corresponding to the provided ID
	UpdateContent(id uint, newContent string) error

	// Deletes the comment corresponding to the provided ID
	Delete(id uint) error
}
