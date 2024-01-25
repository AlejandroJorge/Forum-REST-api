package model

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
