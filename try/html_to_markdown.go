package main

import (
	"fmt"
	"log"

	md "github.com/JohannesKaufmann/html-to-markdown"
)

func convertToMarkdown() {
	converter := md.NewConverter("", true, nil)
	html := `<strong>Important</strong>`

	markdown, err := converter.ConvertString(html)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("md ->", markdown)
}
