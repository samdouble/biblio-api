package types

import (
	"time"
	"biblio-api/types/isbnSearch"
)

type Book struct {
	Id string `bson:"id"`
	CreatedAt time.Time `bson:"createdAt"`
	Isbn string `bson:"isbn"`
	SearchId string `bson:"searchId"`
	VolumeInfo isbnSearch.VolumeInfo `bson:"volumeInfo"`
}

type Search struct {
    Id string `bson:"id"`
	CreatedAt time.Time `bson:"createdAt"`
	Isbn string `bson:"isbn"`
	Result IsbnSearchResponse `bson:"result"`
}

type Event struct {
	Isbn string `json:"isbn"`
}

type ResponseBody struct {
	Books []interface{} `json:"books"`
}

type Response struct {
	Body ResponseBody `json:"body"`
}
