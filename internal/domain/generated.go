package domain

type Generated struct {
	Story     []string            `json:"story"`
	ReadAloud []string            `json:"read_aloud"`
	Images    []ImageWithPosition `json:"images"`
}

type ImageWithPosition struct {
	URL      string `json:"url"`
	Type     string `json:"type"`
	Position int    `json:"position"`
}
