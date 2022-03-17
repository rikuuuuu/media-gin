package database

import (
	"context"
	"media-gin/app/domain/model"
	"media-gin/app/utils"
)

type UserRepository struct {
	FirestoreHandler
}

func (repo *UserRepository) Add(u model.User) error {
	ctx := context.Background()

	mapData, err := utils.StructToMap(u)
	if err != nil {
		return err
	}

	err = repo.Set(ctx, users_table_name, u.ID, mapData)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) UpdateSelf(u model.UpdateUser) error {
	ctx := context.Background()

	mapData, err := utils.StructToMap(u)
	if err != nil {
		return err
	}

	err = repo.Set(ctx, users_table_name, u.ID, mapData)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) FindById(identifier string) (user model.User, err error) {
	ctx := context.Background()

	res, err := repo.Get(ctx, users_table_name, identifier)
	if err != nil {
		return user, err
	}

	err = utils.MapToStruct(res, &user)
	if err != nil {
		return user, err
	}

	return
}

func (repo *UserRepository) FindAll() (users []*model.User, err error) {
	ctx := context.Background()

	res, err := repo.GetAll(ctx, users_table_name)
	if err != nil {
		return nil, err
	}

	users = make([]*model.User, 0, len(res))
	for _, r := range res {
		u := new(model.User)
		err = utils.MapToStruct(r, &u)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return
}

func (repo *UserRepository) DeleteById(identifier string) (err error) {
	ctx := context.Background()

	err = repo.Delete(ctx, users_table_name, identifier)
	if err != nil {
		return
	}

	return
}
