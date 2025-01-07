package main

import (
	"context"
	"testing"
	"biblio-api/types"
)

func TestSimpleReturn(t *testing.T) {
	ctx := context.Background()
    ctx = context.WithValue(ctx, "function_version", "Test")
    event := types.Event{Isbn: "0735619670"}
    msg, err := Main(ctx, event)

	want := "Hello 0735619670! This is function version Test"
    if msg.Body != want || err != nil {
        t.Fatalf(`TestSimpleReturn = %q, %v, want match for %#q, nil`, msg, err, want)
    }
}
