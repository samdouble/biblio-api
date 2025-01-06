package main

import (
	"context"
	"biblio-api/types"
)

func Main(ctx context.Context, event types.Event) (types.Response, error) {
	if event.Name == "" {
		event.Name = "stranger"
	}
	version := ctx.Value("function_version").(string)
	return types.Response {
		Body: "Hello " + event.Name + "! This is function version " + version,
	}, nil
}
