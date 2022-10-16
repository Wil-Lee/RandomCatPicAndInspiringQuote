package main

import (
	"html/template"
	"net/http"
	"strings"
)

var comboPath string = "/combo/"
var comboSeparator string = "."

type CatPicAndQuote struct {
	CatPic string
	Quote  string
}

func comboHandler(w http.ResponseWriter, r *http.Request) {

	requestedCombo := strings.TrimPrefix(r.URL.Path, comboPath)
	catPicID, quoteID, _ := strings.Cut(requestedCombo, comboSeparator)

	c := CatPicAndQuote{CatPic: catPicURL_byID(catPicID), Quote: quoteByID(quoteID)}
	t, _ := template.ParseFiles("catAndQuote.html")
	t.Execute(w, c)
}
