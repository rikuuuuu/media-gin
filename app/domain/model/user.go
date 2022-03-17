package model

type Users []User

type User struct {
	ID    string `db:"id" form:"id" json:"id" firestore:"id,omitempty"`
	Email string `db:"email" form:"email" json:"email" firestore:"email,omitempty"`
	Name  string `db:"name" form:"name" json:"name" firestore:"name,omitempty"`
}