package isbnSearch

type Epub struct {
	IsAvailable bool
}

type Pdf struct {
	IsAvailable bool
}

type AccessInfo struct {
	Country string
	Viewability string
	Embeddable bool
	PublicDomain bool
	TextToSpeechPermission string
	Epub Epub
	Pdf Pdf
	AccessViewStatus string
}
