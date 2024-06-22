// TODO: https://www.reddit.com/r/golang/comments/1dl4q0f/what_are_your_musthave_libraries/

package main

import "fmt"

func main() {
	url := "https://github.com/teamhanko/hanko/blob/main/frontend/elements/README.md"
	md := convertUrlToMarkdown(url)
	fmt.Println(md)
}
