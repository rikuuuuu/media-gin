package usecase

import "media-gin/app/domain/model"

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u model.User) (user model.User, err error) {
	identifier, err := interactor.UserRepository.Store(u)
	if err != nil {
		return
	}
	user, err = interactor.UserRepository.FindById(identifier)
	return
}

func (interactor *UserInteractor) Users() (user model.Users, err error) {
	user, err = interactor.UserRepository.FindAll()
	return
}

func (interactor *UserInteractor) UserById(identifier int) (user model.User, err error) {
	user, err = interactor.UserRepository.FindById(identifier)
	return
}
