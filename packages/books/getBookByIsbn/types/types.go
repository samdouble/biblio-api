package types

import (
	"encoding/json"
	"time"
)

type Search struct {
    Id string `bson:"id"`
	CreatedAt time.Time `bson:"createdAt"`
	Isbn string `bson:"isbn"`
	Result IsbnSearchResponse `bson:"result"`
}

type Event struct {
	Isbn string `json:"isbn"`
}

type Response struct {
	Body json.RawMessage `json:"body"`
}
