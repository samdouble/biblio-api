package main

import (
	"context"
	"encoding/json"
	"fmt"
    "log"
    "net/http"
	"os"
	"time"
    "github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"biblio-api/db"
	"biblio-api/types"
)

func Main(ctx context.Context, event types.Event) (types.Response, error) {
	// if event.Isbn == "" {
	// 	log.Println("ISBN is required")
	// 	return nil, fmt.Errorf("ISBN is required")
	// }

	client := db.ResolveClientDB(os.Getenv("MONGO_URL"))
	defer db.CloseClientDB()

	database := client.Database(os.Getenv("MONGO_DBNAME"))
	booksCollection := database.Collection("books")
	searchesCollection := database.Collection("searches")

	response, err := http.Get(
		fmt.Sprintf(
			"https://www.googleapis.com/books/v1/volumes?q=isbn:%s&key=%s",
			event.Isbn,
			os.Getenv("GOOGLE_BOOKS_API_TOKEN"),
		),
	)
    if err != nil {
        log.Fatal(err)
    }

	isbnSearchResponse := &types.IsbnSearchResponse{}
	json.NewDecoder(response.Body).Decode(isbnSearchResponse)
	searchId := uuid.New().String()
	search := types.Search{
		Id: searchId,
		CreatedAt: time.Now().UTC(),
		Isbn: event.Isbn,
		Result: *isbnSearchResponse,
	}
	var books []interface{}
	for _, searchItem := range search.Result.Items {
		books = append(
			books,
			types.Book{
				Id: uuid.New().String(),
				CreatedAt: time.Now().UTC(),
				Isbn: event.Isbn,
				SearchId: searchId,
				VolumeInfo: searchItem.VolumeInfo,
			},
		)
	}

	wc := writeconcern.Majority()
	transactionOptions := options.Transaction().SetWriteConcern(wc)
	session, err := client.StartSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.EndSession(context.TODO())

	_, err = session.WithTransaction(context.TODO(), func(ctx mongo.SessionContext) (interface{}, error) {
		_, err := searchesCollection.InsertOne(context.TODO(), search)
		if err != nil {
			return nil, err
		}
		var booksInsertResult *mongo.InsertManyResult
		if len(books) > 0 {
			booksInsertResult, err = booksCollection.InsertMany(context.TODO(), books)
			if err != nil {
				return nil, err
			}
		}
		return booksInsertResult, nil
	}, transactionOptions)
	if err != nil {
		log.Fatal(err)
	}

	return types.Response {
		Body: types.ResponseBody{
			Books: books,
		},
	}, nil
}
