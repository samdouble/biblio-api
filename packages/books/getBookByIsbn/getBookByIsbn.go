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
	// 	return types.Response{
	// 		Body: "ISBN is required",
	// 	}
	// }

	client := db.ResolveClientDB(os.Getenv("MONGO_URL"))
	defer db.CloseClientDB()

	searchesCollection := client.Database(os.Getenv("MONGO_DBNAME")).Collection("searches")

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
	search := types.Search{
		Id: uuid.New().String(),
		CreatedAt: time.Now().UTC(),
		Isbn: event.Isbn,
		Result: *isbnSearchResponse,
	}

	wc := writeconcern.Majority()
	transactionOptions := options.Transaction().SetWriteConcern(wc)
	session, err := client.StartSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.EndSession(context.TODO())

	result, err := session.WithTransaction(context.TODO(), func(ctx mongo.SessionContext) (interface{}, error) {
		result, err := searchesCollection.InsertOne(context.TODO(), search)
		return result, err
	}, transactionOptions)
	if err != nil {
		log.Fatal(err)
	}

	jsonResult, err := json.Marshal(result)
	return types.Response {
		Body: jsonResult,
	}, nil
}
