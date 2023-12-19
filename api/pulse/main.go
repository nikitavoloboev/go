package main

import (
	"fmt"
	"strconv"

	"net/http"

	"github.com/gocolly/colly"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// get all urls from https://news.ycombinator.com/newest
func getHnNewestUrls(last int) {
	fmt.Println(last, "last")
}

func main() {
	fmt.Println("wow")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/hn-newest/{last}", func(w http.ResponseWriter, r *http.Request) {
		lastStr := chi.URLParam(r, "last")
		last, err := strconv.Atoi(lastStr)
		if err != nil {
			http.Error(w, "Invalid parameter", http.StatusBadRequest)
			return
		}
		getHnNewestUrls(last)
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":3000", r)

	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("http://go-colly.org/")
}
