package main

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var CatPicAPI_URL string = "https://api.thecatapi.com/v1/images/search"
var QuoteAPI_URL string = "https://api.quotable.io/random"

type CatPicAndQuote struct {
	CatPic string
	Quote  string
}

type CatPicAPI struct {
	ID     string `json:"id"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type QuoteAPI struct {
	ID           string `json:"_id"`
	Content      string `json:"content"`
	Author       string `json:"author"`
	Tags         string `json:"tags"`
	AuthorSlug   string `json:"length"`
	DateAdded    string `json:"dateAdded"`
	DateModified string `json:"dateModified"`
}

func catQuoteHandler(w http.ResponseWriter, r *http.Request) {
	var quote QuoteAPI
	json.Unmarshal(extractHTMLbody(QuoteAPI_URL), &quote)

	c := CatPicAndQuote{CatPic: string(catPicURL()), Quote: "\"" + quote.Content + "\"" + " ~ " + quote.Author}
	t, _ := template.ParseFiles("catQuote.html")
	t.Execute(w, c)
}

func extractHTMLbody(url string) []byte {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	htmlBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return htmlBody
}

func catPicURL() []byte {
	var catPicAPI []CatPicAPI
	json.Unmarshal(extractHTMLbody(CatPicAPI_URL), &catPicAPI)

	return []byte(catPicAPI[0].URL)
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
