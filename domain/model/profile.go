package model

type Profile struct {
	UserID         uint
	DisplayName    string
	TagName        string
	PicturePath    string
	BackgroundPath string
	Followers      uint
	Follows        uint
}
