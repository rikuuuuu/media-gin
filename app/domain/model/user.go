package model

type Users []User

type User struct {
	ID    int
	Email string
	Name  string
}

func (u *User) Build() *User {
	return u
}
