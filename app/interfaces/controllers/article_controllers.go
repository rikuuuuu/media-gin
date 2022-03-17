package controllers

import (
	"media-gin/app/domain/model"
	"media-gin/app/interfaces/database"
	"media-gin/app/usecase/article"
)

type ArticleController struct {
	Interactor article.ArticleInteractor
}

func NewArticleController(handler database.FirestoreHandler) *ArticleController {
	return &ArticleController{
		Interactor: article.ArticleInteractor{
			ArticleRepository: &database.ArticleRepository{
				FirestoreHandler: handler,
			},
		},
	}
}

func (controller *ArticleController) Create(c Context) {
	a := model.Article{}
	err := c.Bind(&a)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}

	article, err := controller.Interactor.Add(a)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, article)
}

func (controller *ArticleController) Index(c Context) {
	article, err := controller.Interactor.Articles()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, article)
}

func (controller *ArticleController) Show(c Context) {
	id := c.Param("id")
	article, err := controller.Interactor.ArticleById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, article)
}

func (controller *ArticleController) Update(c Context) {
	a := model.Article{}
	err := c.Bind(&a)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}

	article, err := controller.Interactor.Edit(a)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, article)
}

func (controller *ArticleController) Delete(c Context) {
	id := c.Param("id")
	err := controller.Interactor.DeleteArticleById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, nil)
}
