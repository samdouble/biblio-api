package isbnSearch

type Dimensions struct {
	Height string
	Width string
	Thickness string
}

type ImageLinks struct {
	SmallThumbnail string
	Thumbnail string
	Small string
	Medium string
	Large string
	ExtraLarge string
}

type IndustryIdentifier struct {
	Type string
	Identifier string
}

type VolumeInfo struct {
	Title string
	Authors []string
	Publisher string
	PublishedDate string
	Description string
	IndustryIdentifiers []IndustryIdentifier
	PageCount int
	Dimensions Dimensions
	PrintType string
	MainCategory string
	Categories []string
	AverageRating float64
	RatingsCount int
	ContentVersion string
	ImageLinks ImageLinks
	Language string
	InfoLink string
	CanonicalVolumeLink string
}
