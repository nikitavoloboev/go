package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"net/http"

	"github.com/gocolly/colly"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// get all urls from https://news.ycombinator.com/newest
func getHnNewestUrls(last int) []string {
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

	c.Visit("https://news.ycombinator.com/newest")
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
		links := getHnNewestUrls(last)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(links)
	})
	http.ListenAndServe(":3000", r)
}
