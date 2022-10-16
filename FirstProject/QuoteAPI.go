package main

import (
	"encoding/json"
)

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

func quote() string {
	var quote QuoteAPI
	json.Unmarshal(extractHTMLbody(QuoteAPI_URL), &quote)

	return "\"" + quote.Content + "\"" + " ~ " + quote.Author
}
