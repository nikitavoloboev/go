package main

import (
	"flag"
	"os"

	"github.com/atotto/clipboard"
)

func main() {
	saveTopLink := flag.Bool("top", false, "save top link")
	flag.Parse()

	if *saveTopLink {
		save("/Users/nikiv/Dropbox/Data/top-links.md")
		return
	}
	save("/Users/nikiv/Dropbox/Data/links.md")
}

// TODO: include title of the URL
// save link from clipboard to a list
func save(path string) {
	// get url from clipboard
	clip, _ := clipboard.ReadAll()
	clip = clip + "\n"

	// TODO: check it's url

	// append url to file (on new line)
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err = f.WriteString(clip); err != nil {
		panic(err)
	}
}

// open a number of links from a list in browser & delete them
func pop(list string, number int) {

}
