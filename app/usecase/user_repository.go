package usecase

import "media-gin/app/domain/model"

type UserRepository interface {
	Store(model.User) (int, error)
	FindById(int) (model.User, error)
	FindAll() (model.Users, error)
}
