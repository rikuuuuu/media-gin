package model

import "time"

type Article struct {
	ID           string    `db:"id" form:"id" json:"id" firestore:"id,omitempty"`
	Title        string    `db:"title" form:"title" json:"title" firestore:"title,omitempty"`
	Content      string    `db:"content" form:"content" json:"content" firestore:"content,omitempty"`
	ThumbnailURL string    `db:"thumbnail_url" form:"thumbnail_url" json:"thumbnail_url" firestore:"thumbnail_url,omitempty"`
	UserID       string    `db:"user_id" form:"user_id" json:"user_id" firestore:"user_id,omitempty"`
	CreatedAt    time.Time `db:"created_at" json:"created_at" firestore:"created_at,omitempty"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at" firestore:"updated_at,omitempty"`
}

type UpdateArticle struct {
	ID           string    `db:"id" form:"id" json:"id" firestore:"id,omitempty"`
	Title        string    `db:"title" form:"title" json:"title" firestore:"title,omitempty"`
	Content      string    `db:"content" form:"content" json:"content" firestore:"content,omitempty"`
	ThumbnailURL string    `db:"thumbnail_url" form:"thumbnail_url" json:"thumbnail_url" firestore:"thumbnail_url,omitempty"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at" firestore:"updated_at,omitempty"`
}
