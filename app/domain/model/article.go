package model

type Articles []Article

type Article struct {
	ID           int
	Title        string
	Content      string
	ThumbnailURL string
	User         *User
}

func (a *Article) Build() *Article {
	return a
}
