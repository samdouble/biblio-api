package main

import (
	"context"
	"fmt"
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
	want := types.Response{
        Body: types.ResponseBody{
            Books: books,
        },
    }
    fmt.Printf("unmarshalled: %#v\n", msg)
    fmt.Printf("unmarshalled: %#v\n", want)
    if !reflect.DeepEqual(msg, want) || err != nil {
        t.Fatalf(`TestSimpleReturn = %#v, %v want match for %#v, nil`, msg, err, want)
    }
}
