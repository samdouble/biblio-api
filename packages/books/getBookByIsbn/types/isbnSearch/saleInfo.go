package isbnSearch

type ListPrice struct {
	Amount float64 `json:"amount"`
	CurrencyCode string `json:"currencyCode"`
}

type RentPrice struct {
	Amount float64 `json:"amount"`
	CurrencyCode string `json:"currencyCode"`
}

type SaleInfo struct {
	Country string `json:"country"`
	IsEbook bool `json:"isEbook"`
	Saleability string `json:"saleability"`
	ListPrice ListPrice `json:"listPrice"`
	RentPrice RentPrice `json:"rentPrice"`
	BuyLink string `json:"buyLink"`
}
