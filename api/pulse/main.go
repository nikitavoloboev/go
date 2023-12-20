package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"net/http"

	"github.com/gocolly/colly"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// uses https://github.com/HackerNews/API#new-top-and-best-stories api to get newest stories (https://news.ycombinator.com/newest)
func getHnNewestPostsIds() ([]int, error) {
	resp, err := http.Get("https://hacker-news.firebaseio.com/v0/newstories.json?print=pretty")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var ids []int
	err = json.NewDecoder(resp.Body).Decode(&ids)
	if err != nil {
		return nil, err
	}
	return ids, nil
}

// Link represents a link
type Link struct {
	title string
	url   string
}

func processHnPost(id string) []Link {
	c := colly.NewCollector()
	var posts []Link

	c.OnHTML("span.titleline > a[href]", func(e *colly.HTMLElement) {
		post := Link{
			title: e.Text,
			url:   e.Attr("href"),
		}
		posts = append(posts, post)
	})

	hnURL := fmt.Sprintf("https://news.ycombinator.com/item?id=%s", id)
	c.Visit(hnURL)
	return posts
}

func processLobstersPost(id string) {

}

// get all urls from https://lobste.rs/newest
func getLobstersNewestUrls(last int) []string {
	c := colly.NewCollector()
	var links []string

	// find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		links = append(links, link)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL)
	})

	c.Visit("https://lobste.rs/newest")
	return links
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/hn-newest/{last}", func(w http.ResponseWriter, r *http.Request) {
		lastStr := chi.URLParam(r, "last")
		last, err := strconv.Atoi(lastStr)
		if err != nil {
			http.Error(w, "Invalid parameter", http.StatusBadRequest)
			return
		}
		links := getLobstersNewestUrls(last)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(links)
	})
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
