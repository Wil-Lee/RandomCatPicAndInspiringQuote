package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {

	c := CatPicAndQuote{CatPic: randomCatPicURL(), Quote: randomQuote()}
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, c)
}

func buildComboID(catPicID string, quoteID string) string {
	return catPicID + comboSeparator + quoteID
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
	http.ListenAndServe(":"+port, mux)
}
