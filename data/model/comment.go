package model

type Comment struct {
	ID      uint
	PostID  uint
	UserID  uint
	Content string
	Likes   uint
}
