package infra

import (
	"context"
	"log"
	"media-gin/app/interfaces/database"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type FirestoreHandler struct{}

func NewFirestoreHandler() database.FirestoreHandler {
	firestoreHandler := new(FirestoreHandler)
	return firestoreHandler
}

// firebaseInit.. Firebase初期化メソッド
func firebaseInit(ctx context.Context) (client *firestore.Client, err error) {

	sa := option.WithCredentialsFile("config/dev-secret.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	client, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return client, nil
}

// Get.. Firestore 取得
func (f *FirestoreHandler) Get(ctx context.Context, collection string, doc string) (res map[string]interface{}, err error) {

	client, err := firebaseInit(ctx)
	if err != nil {
		return nil, err
	}

	dsnap, err := client.Collection(collection).Doc(doc).Get(ctx)
	if err != nil {
		return nil, err
	}

	err = dsnap.DataTo(&res)
	if err != nil {
		return nil, err
	}

	defer client.Close()
	return
}

// GetAll.. Firestore 全件取得
func (f *FirestoreHandler) GetAll(ctx context.Context, collection string) (resData []map[string]interface{}, err error) {

	client, err := firebaseInit(ctx)
	if err != nil {
		return nil, err
	}

	allData := client.Collection(collection).Documents(ctx)
	docs, err := allData.GetAll()
	if err != nil {
		log.Fatalf("Failed adding getAll: %v", err)
	}

	for _, doc := range docs {
		resData = append(resData, doc.Data())
	}

	defer client.Close()
	return
}

// Add.. Firestore 追加
func (f *FirestoreHandler) New(ctx context.Context, collection string, args map[string]interface{}) error {

	client, err := firebaseInit(ctx)
	if err != nil {
		return err
	}

	ref := client.Collection(collection).NewDoc()
	args["id"] = ref.ID
	_, err = ref.Set(ctx, args)
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}
	defer client.Close()
	return nil
}

// Set.. Firestore 追加または上書き
func (f *FirestoreHandler) Set(ctx context.Context, collection string, doc string, args map[string]interface{}) error {

	client, err := firebaseInit(ctx)
	if err != nil {
		return err
	}

	_, err = client.Collection(collection).Doc(doc).Set(ctx, args)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer client.Close()
	return nil
}

// Delete.. Firestore 削除
func (f *FirestoreHandler) Delete(ctx context.Context, collection string, doc string) error {

	client, err := firebaseInit(ctx)
	if err != nil {
		return err
	}

	_, err = client.Collection(collection).Doc(doc).Delete(ctx)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer client.Close()
	return nil
}
