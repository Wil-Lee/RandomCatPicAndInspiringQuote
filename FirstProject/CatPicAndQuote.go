package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var tpl = template.Must(template.ParseFiles("index.html"))
var APIurl string = "https://api.thecatapi.com/v1/images/search"

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
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
	body := extractHTMLbody(APIurl)
	startIndex := bytes.Index(body, []byte("http"))
	endIndex := bytes.Index(body, []byte(".jpg")) + 4 // +4 to get index of g

	return body[startIndex:endIndex]
}

func main() {
	fmt.Print(string(catPicURL()))

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
	http.ListenAndServe(":"+port, mux)
}
