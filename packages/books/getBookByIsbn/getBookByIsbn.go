package main

import (
	"context"
	"os"
    "go.mongodb.org/mongo-driver/v2/mongo"
    "go.mongodb.org/mongo-driver/v2/mongo/options"
    "go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"biblio-api/types"
)

func Main(ctx context.Context, event types.Event) (types.Response, error) {
	client, _ := mongo.Connect(options.Client().ApplyURI(os.Getenv("MONGO_URL")))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	if event.Name == "" {
		event.Name = "stranger"
	}
	version := ctx.Value("function_version").(string)
	return types.Response {
		Body: "Hello " + event.Name + "! This is function version " + version,
	}, nil
}
