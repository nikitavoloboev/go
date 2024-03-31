package main

import (
	"io/ioutil"
	"log"
	"net/http"

	md "github.com/JohannesKaufmann/html-to-markdown"
)

func convertUrlToMarkdown(url string) string {
	// Perform an HTTP GET request to the URL
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error fetching URL: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body to get the HTML content
	htmlBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}
	html := string(htmlBytes)

	// Convert the HTML to Markdown
	converter := md.NewConverter("", true, nil)
	markdown, err := converter.ConvertString(html)
	if err != nil {
		log.Fatal(err)
	}

	return markdown
}
