package main

import "fmt"

func main() {
	url := "https://github.com/teamhanko/hanko/blob/main/frontend/elements/README.md"
	md := convertUrlToMarkdown(url)
	fmt.Println(md)
}
