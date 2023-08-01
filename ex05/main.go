package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"strings"
)

func main() {
	animes := make([]Anime, 0)
	urls := make([]string, 0)
	pageLimit := 3
	i := 0

	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

	c.OnHTML("ul.pagination > li:not(.active)", func(e *colly.HTMLElement) {
		url := getURL(e)
		if len(url) > 0 && !contains(urls, url) {
			urls = append(urls, url)
		}
	})

	c.OnHTML("div.movie-item", func(e *colly.HTMLElement) {
		anime := Anime{
			Title:   e.ChildAttr("a", "title"),
			URL:     getURL(e),
			Episode: e.ChildAttr("span.ribbon", "text"),
		}
		animes = append(animes, anime)
	})

	c.OnScraped(func(_ *colly.Response) {
		if len(urls) > 0 && i < pageLimit {
			url := urls[0]
			urls = urls[1:]
			i++
			c.Visit(url)
		}
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong: ", err)
	})

	c.OnRequest(func(e *colly.Request) {
		fmt.Printf("visiting %v\n", e.URL)
	})

	err := c.Visit("https://anime47.com/danh-sach/phim-moi/1.html")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Anime scraped: %d", len(animes))
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func getURL(e *colly.HTMLElement) string {
	return "https://" + strings.Replace(e.ChildAttr("a", "href"), ".", e.Request.URL.Host, 1)
}
