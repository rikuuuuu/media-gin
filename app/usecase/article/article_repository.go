package article

import "media-gin/app/domain/model"

type ArticleRepository interface {
	CreateByUser(model.Article) error
	UpdateByUser(model.Article) error
	FindById(string) (model.Article, error)
	FindAll() ([]*model.Article, error)
	DeleteById(string) error
}
