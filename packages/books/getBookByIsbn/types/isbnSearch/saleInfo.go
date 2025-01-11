package isbnSearch

type ListPrice struct {
	Amount float64
	CurrencyCode string
}

type RentPrice struct {
	Amount float64
	CurrencyCode string
}

type SaleInfo struct {
	Country string
	IsEbook bool
	Saleability string
	ListPrice ListPrice
	RentPrice RentPrice
	BuyLink string
}
