package isbnSearch

type Dimensions struct {
	Height string
	Width string
	Thickness string
}

type ImageLinks struct {
	SmallThumbnail string `json:"smallThumbnail"`
	Thumbnail string `json:"thumbnail"`
	Small string `json:"small"`
	Medium string `json:"medium"`
	Large string `json:"large"`
	ExtraLarge string `json:"extraLarge"`
}

type IndustryIdentifier struct {
	Type string `json:"type"`
	Identifier string `json:"identifier"`
}

type VolumeInfo struct {
	Title string `json:"title"`
	Authors []string `json:"authors"`
	Publisher string `json:"publisher"`
	PublishedDate string `json:"publishedDate"`
	Description string `json:"description"`
	IndustryIdentifiers []IndustryIdentifier `json:"industryIdentifiers"`
	PageCount int `json:"pageCount"`
	Dimensions Dimensions `json:"dimensions"`
	PrintType string `json:"printType"`
	MainCategory string `json:"mainCategory"`
	Categories []string `json:"categories"`
	AverageRating float64 `json:"averageRating"`
	RatingsCount int `json:"ratingsCount"`
	ContentVersion string `json:"contentVersion"`
	ImageLinks ImageLinks `json:"imageLinks"`
	Language string `json:"language"`
	InfoLink string `json:"infoLink"`
	CanonicalVolumeLink string `json:"canonicalVolumeLink"`
}
