package main

var QuoteAPI_URL string = "https://api.quotable.io/random"

type QuoteAPI struct {
	ID           string `json:"_id"`
	Content      string `json:"content"`
	Author       string `json:"author"`
	Tags         string `json:"tags"`
	AuthorSlug   string `json:"length"`
	DateAdded    string `json:"dateAdded"`
	DateModified string `json:"dateModified"`
}
