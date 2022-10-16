package main

import (
	"io"
	"log"
	"net/http"
)

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
