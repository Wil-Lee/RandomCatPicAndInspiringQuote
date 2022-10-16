package main

import (
	"encoding/json"
)

var catPicAPI_URL string = "https://api.thecatapi.com/v1/images/search"
var catPicAPI_URL_byID string = "https://api.thecatapi.com/v1/images/"

type CatPicAPI struct {
	ID     string `json:"id"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

func randomCatPicURL() string {
	var catPicAPI []CatPicAPI
	json.Unmarshal(extractHTMLbody(catPicAPI_URL), &catPicAPI)

	return catPicAPI[0].URL
}

func catPicURL_byID(id string) string {
	var catPicAPI CatPicAPI
	json.Unmarshal(extractHTMLbody(catPicAPI_URL_byID+id), &catPicAPI)

	return catPicAPI.URL
}
