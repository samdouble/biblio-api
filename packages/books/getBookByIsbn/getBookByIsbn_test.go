package main

import (
	"context"
	"testing"
	"biblio-api/types"
)

func TestSimpleReturn(t *testing.T) {
	ctx := context.Background()
    ctx = context.WithValue(ctx, "function_version", "Test")
    event := types.Event{Name: "samdouble"}
    msg, err := Main(ctx, event)

	want := "Hello samdouble! This is function version Test"
    if msg.Body != want || err != nil {
        t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
    }
}
