package database

import "context"

type FirestoreHandler interface {
	Get(ctx context.Context, collection string, doc string) (map[string]interface{}, error)
	GetAll(ctx context.Context, collection string) ([]map[string]interface{}, error)
	Add(ctx context.Context, collection string, doc string, args ...interface{}) error
	Set(ctx context.Context, collection string, doc string, args map[string]interface{}) error
	Delete(ctx context.Context, collection string, doc string) error
}
