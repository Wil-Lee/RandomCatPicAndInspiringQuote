package main

import (
	"encoding/json"
)

var CatPicAPI_URL string = "https://api.thecatapi.com/v1/images/search"

type CatPicAPI struct {
	ID     string `json:"id"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

func catPicURL() []byte {
	var catPicAPI []CatPicAPI
	json.Unmarshal(extractHTMLbody(CatPicAPI_URL), &catPicAPI)

	return []byte(catPicAPI[0].URL)
}
