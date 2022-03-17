package database

import (
	"context"
	"media-gin/app/domain/model"
	"media-gin/app/utils"
)

type ArticleRepository struct {
	FirestoreHandler
}

func (repo *ArticleRepository) CreateByUser(a model.Article) (id string, err error) {
	ctx := context.Background()

	mapData, err := utils.StructToMap(a)
	if err != nil {
		return
	}

	id, err = repo.New(ctx, articles_table_name, mapData)
	if err != nil {
		return
	}
	return
}

func (repo *ArticleRepository) UpdateByUser(a model.UpdateArticle) error {
	ctx := context.Background()

	mapData, err := utils.StructToMap(a)
	if err != nil {
		return err
	}

	err = repo.Set(ctx, articles_table_name, a.ID, mapData)
	if err != nil {
		return err
	}
	return nil
}

func (repo *ArticleRepository) FindById(identifier string) (Article model.Article, err error) {
	ctx := context.Background()

	res, err := repo.Get(ctx, articles_table_name, identifier)
	if err != nil {
		return Article, err
	}

	err = utils.MapToStruct(res, &Article)
	if err != nil {
		return Article, err
	}

	return
}

func (repo *ArticleRepository) FindAll() (Articles []*model.Article, err error) {
	ctx := context.Background()

	res, err := repo.GetAll(ctx, articles_table_name)
	if err != nil {
		return nil, err
	}

	Articles = make([]*model.Article, 0, len(res))
	for _, r := range res {
		u := new(model.Article)
		err = utils.MapToStruct(r, &u)
		if err != nil {
			return nil, err
		}
		Articles = append(Articles, u)
	}

	return
}

func (repo *ArticleRepository) DeleteById(identifier string) (err error) {
	ctx := context.Background()

	err = repo.Delete(ctx, articles_table_name, identifier)
	if err != nil {
		return
	}

	return
}
