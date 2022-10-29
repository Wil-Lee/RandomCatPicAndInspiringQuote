package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type ComboLinkCatPicQuote struct {
	CatQuote  CatPicAndQuote
	ComboLink string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	cQuote := CatPicAndQuote{CatPic: randomCatPicURL(), Quote: randomQuote()}
	var comboLink string = r.Host + buildComboPath()
	ComboLinkAndCatPicQuote := ComboLinkCatPicQuote{CatQuote: cQuote, ComboLink: comboLink}

	t, _ := template.ParseFiles("index.html")
	t.Execute(w, ComboLinkAndCatPicQuote)
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

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc(comboPath, comboHandler)
	mux.HandleFunc(savedCombosPath, savedCombosHandler)
	http.ListenAndServe(":"+port, mux)
}
