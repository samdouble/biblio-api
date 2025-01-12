package main

import (
	"context"
	"reflect"
	"testing"
	"biblio-api/types"
)

func TestSimpleReturn(t *testing.T) {
	ctx := context.Background()
    ctx = context.WithValue(ctx, "function_version", "Test")
    event := types.Event{Isbn: "0735619670"}
    msg, err := Main(ctx, event)

    var books []interface{}
	want := types.ResponseBody{
        Books: books,
    }
    if !reflect.DeepEqual(msg.Body, want) || err != nil {
        t.Fatalf(`TestSimpleReturn = %q, %v, want match for %#q, nil`, msg, err, want)
    }
}
