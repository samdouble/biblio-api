package isbnSearch

type Epub struct {
	IsAvailable bool `json:"isAvailable"`
}

type Pdf struct {
	IsAvailable bool `json:"isAvailable"`
}

type AccessInfo struct {
	Country string `json:"country"`
	Viewability string `json:"viewability"`
	Embeddable bool `json:"embeddable"`
	PublicDomain bool `json:"publicDomain"`
	TextToSpeechPermission string `json:"textToSpeechPermission"`
	Epub Epub `json:"epub"`
	Pdf Pdf `json:"pdf"`
	AccessViewStatus string `json:"accessViewStatus"`
}
