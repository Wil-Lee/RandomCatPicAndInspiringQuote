package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type CatPicAndQuote struct {
	CatPic string
	Quote  string
}

func catQuoteHandler(w http.ResponseWriter, r *http.Request) {
	var quote QuoteAPI
	json.Unmarshal(extractHTMLbody(QuoteAPI_URL), &quote)
	c := CatPicAndQuote{CatPic: string(catPicURL()), Quote: "\"" + quote.Content + "\"" + " ~ " + quote.Author}
	t, _ := template.ParseFiles("catQuote.html")
	t.Execute(w, c)
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", catQuoteHandler)
	http.ListenAndServe(":"+port, mux)
}
