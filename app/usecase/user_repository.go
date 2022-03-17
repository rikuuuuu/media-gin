package usecase

import "media-gin/app/domain/model"

type UserRepository interface {
	Add(model.User) error
	FindById(string) (model.User, error)
	FindAll() ([]*model.User, error)
	DeleteById(string) error
}
