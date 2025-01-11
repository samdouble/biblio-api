package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
    "io/ioutil"
    "log"
    "net/http"
	"os"
	"biblio-api/db"
	"biblio-api/types"
)

type Search struct {
    ID int `bson:"_id"`
    Items []interface{} `bson:"items"`
    Kind string `bson:"kind"`
	TotalItems int `bson:"totalItems"`
}

func Main(ctx context.Context, event types.Event) (types.Response, error) {
	client := db.ResolveClientDB(os.Getenv("MONGO_URL"))
	defer db.CloseClientDB()

	booksCollection := client.Database(os.Getenv("MONGO_DBNAME")).Collection("books")

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
    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

	// doc := interface{}{
	// 	Book{Title: "Cat's Cradle", Author: "Kurt Vonnegut Jr."},
	// 	Book{Title: "In Memory of Memory", Author: "Maria Stepanova"},
	// 	Book{Title: "Pride and Prejudice", Author: "Jane Austen"},
	// }
	var doc types.IsbnSearchResponse
	err = bson.UnmarshalExtJSON([]byte(string(responseData)), true, &doc)

	if err != nil {
		log.Fatal(err)
	}

	_, err = booksCollection.InsertOne(context.Background(), doc)
	// log.Println(result, err)

	version := ctx.Value("function_version").(string)
	return types.Response {
		Body: "Hello " + event.Isbn + "! This is function version " + version,
	}, nil
}
