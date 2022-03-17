package controllers

import (
	"media-gin/app/domain/model"
	"media-gin/app/interfaces/database"
	"media-gin/app/usecase/user"
)

type UserController struct {
	Interactor user.UserInteractor
}

func NewUserController(handler database.FirestoreHandler) *UserController {
	return &UserController{
		Interactor: user.UserInteractor{
			UserRepository: &database.UserRepository{
				FirestoreHandler: handler,
			},
		},
	}
}

func (controller *UserController) Create(c Context) {
	u := model.User{}
	err := c.Bind(&u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}

	user, err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, user)
}

func (controller *UserController) Index(c Context) {
	users, err := controller.Interactor.Users()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, users)
}

func (controller *UserController) Show(c Context) {
	id := c.Param("id")
	user, err := controller.Interactor.UserById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, user)
}

func (controller *UserController) Update(c Context) {
	u := model.UpdateUser{}
	err := c.Bind(&u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}

	user, err := controller.Interactor.UserUpdate(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, user)
}

func (controller *UserController) Delete(c Context) {
	id := c.Param("id")
	err := controller.Interactor.DeleteUserById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, nil)
}
