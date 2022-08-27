package database

import "context"

type FirestoreHandler interface {
	Get(ctx context.Context, collection string, doc string) (map[string]interface{}, error)
	GetAll(ctx context.Context, collection string) ([]map[string]interface{}, error)
	GetAllWithPage(ctx context.Context, collection string, offset int, limit int) ([]map[string]interface{}, error)
	New(ctx context.Context, collection string, args map[string]interface{}) (string, error)
	Set(ctx context.Context, collection string, doc string, args map[string]interface{}) error
	Delete(ctx context.Context, collection string, doc string) error
}
