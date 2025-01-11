package types

import (
	"biblio-api/types/isbnSearch"
)

type Item struct {
	AccessInfo isbnSearch.AccessInfo
	Etag string
	GoogleBooksId string
	SaleInfo isbnSearch.SaleInfo
	SelfLink string
	VolumeInfo isbnSearch.VolumeInfo
}

type IsbnSearchResponse struct {
    Items []Item
    Kind string
	TotalItems int
}
