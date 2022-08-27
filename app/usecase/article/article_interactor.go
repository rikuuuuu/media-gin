package article

import "media-gin/app/domain/model"

type ArticleInteractor struct {
	ArticleRepository ArticleRepository
}

func (interactor *ArticleInteractor) Add(a model.Article) (article model.Article, err error) {
	id, err := interactor.ArticleRepository.CreateByUser(a)
	if err != nil {
		return
	}

	article, err = interactor.ArticleRepository.FindById(id)
	if err != nil {
		return
	}

	return
}

func (interactor *ArticleInteractor) Edit(a model.UpdateArticle) (Article model.Article, err error) {
	err = interactor.ArticleRepository.UpdateByUser(a)
	if err != nil {
		return
	}
	Article, err = interactor.ArticleRepository.FindById(a.ID)
	if err != nil {
		return
	}
	return
}

func (interactor *ArticleInteractor) Articles() (Article []*model.Article, err error) {
	Article, err = interactor.ArticleRepository.FindAll()
	return
}

func (interactor *ArticleInteractor) ArticleById(identifier string) (Article model.Article, err error) {
	Article, err = interactor.ArticleRepository.FindById(identifier)
	return
}

func (interactor *ArticleInteractor) DeleteArticleById(identifier string) (err error) {
	err = interactor.ArticleRepository.DeleteById(identifier)
	return
}
