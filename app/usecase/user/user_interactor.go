package user

import "media-gin/app/domain/model"

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u model.User) (user model.User, err error) {
	err = interactor.UserRepository.Add(u)
	if err != nil {
		return
	}
	user, err = interactor.UserRepository.FindById(u.ID)
	if err != nil {
		return
	}
	return
}

func (interactor *UserInteractor) UserUpdate(u model.UpdateUser) (user model.User, err error) {
	err = interactor.UserRepository.UpdateSelf(u)
	if err != nil {
		return
	}
	user, err = interactor.UserRepository.FindById(u.ID)
	if err != nil {
		return
	}
	return
}

func (interactor *UserInteractor) Users() (user []*model.User, err error) {
	user, err = interactor.UserRepository.FindAll()
	return
}

func (interactor *UserInteractor) UserById(identifier string) (user model.User, err error) {
	user, err = interactor.UserRepository.FindById(identifier)
	return
}

func (interactor *UserInteractor) DeleteUserById(identifier string) (err error) {
	err = interactor.UserRepository.DeleteById(identifier)
	return
}
