package main

type options struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type topic struct {
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Options []options `json:"options"`
}
