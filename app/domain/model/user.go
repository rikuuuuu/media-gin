package model

import "time"

type User struct {
	ID        string    `db:"id" form:"id" json:"id" firestore:"id,omitempty"`
	Email     string    `db:"email" form:"email" json:"email" firestore:"email,omitempty"`
	Name      string    `db:"name" form:"name" json:"name" firestore:"name,omitempty"`
	CreatedAt time.Time `db:"created_at" json:"created_at" firestore:"created_at,omitempty"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at" firestore:"updated_at,omitempty"`
}

type UpdateUser struct {
	ID        string    `db:"id" form:"id" json:"id" firestore:"id,omitempty"`
	Name      string    `db:"name" form:"name" json:"name" firestore:"name,omitempty"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at" firestore:"updated_at,omitempty"`
}
