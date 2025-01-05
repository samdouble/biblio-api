package main

import (
	"context"
    "fmt"
)

type Event struct {
	Name string `json:"name"`
}

type Response struct {
	Body string `json:"body"`
}

func Main(ctx context.Context, event Event) Response {
	if event.Name == "" {
		event.Name = "stranger"
	}
	version := ctx.Value("function_version").(string)
	return Response {
		Body: "Hello " + event.Name + "! This is function version " + version,
	}
}

func main() {
    ctx := context.Background()
    ctx = context.WithValue(ctx, "function_version", "Test")
    event := Event{Name: "samdouble"}
    fmt.Println(Main(ctx, event))
}
