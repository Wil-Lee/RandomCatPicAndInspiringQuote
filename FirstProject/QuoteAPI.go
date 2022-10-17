package main

import (
	"encoding/json"
)

var quoteAPI_URL string = "https://api.quotable.io/random"
var quoteAPI_URL_byID string = "https://api.quotable.io/quotes/"
var currentQuoteID string

type QuoteAPI struct {
	ID           string `json:"_id"`
	Content      string `json:"content"`
	Author       string `json:"author"`
	Tags         string `json:"tags"`
	AuthorSlug   string `json:"length"`
	DateAdded    string `json:"dateAdded"`
	DateModified string `json:"dateModified"`
}

func randomQuote() string {
	return extractQuote(quoteAPI_URL)
}

func quoteByID(id string) string {
	return extractQuote(quoteAPI_URL_byID + id)
}

func extractQuote(url string) string {
	var quote QuoteAPI
	json.Unmarshal(extractHTMLbody(url), &quote)
	currentQuoteID = quote.ID
	return "\"" + quote.Content + "\"" + " ~ " + quote.Author
}
