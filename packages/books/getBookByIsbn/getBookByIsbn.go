package main

import (
	"context"
	"fmt"
    "io/ioutil"
    "log"
    "net/http"
	"os"
	"biblio-api/db"
	"biblio-api/types"
)

func Main(ctx context.Context, event types.Event) (types.Response, error) {
	client := db.ResolveClientDB(os.Getenv("MONGO_URL"))
	fmt.Println(client)
	defer db.CloseClientDB()

	response, err := http.Get(
		fmt.Sprintf(
			"https://www.googleapis.com/books/v1/volumes?q=isbn:%s&key=%s",
			event.Isbn,
			os.Getenv("GOOGLE_BOOKS_API_TOKEN"),
		),
	)
    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }
    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(responseData))

	version := ctx.Value("function_version").(string)
	return types.Response {
		Body: "Hello " + event.Isbn + "! This is function version " + version,
	}, nil
}
