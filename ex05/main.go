package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/gocolly/colly"
)

func main() {
	animes := make([]Anime, 0)
	urls := make([]string, 0)
	visited := make([]string, 0)
	var wg sync.WaitGroup
	chs := make(chan interface{}, 4)
	i := 1
	pageLimit := 40

	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

	c.OnRequest(func(e *colly.Request) {
		fmt.Printf("visiting %v\n", e.URL)
	})

	c.OnHTML("ul.pagination > li:not(.active)", func(e *colly.HTMLElement) {
		url := getURL(e)
		if len(url) > 0 && !contains(visited, url) {
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

	c.OnError(func(e *colly.Response, err error) {
		defer wg.Done()
		fmt.Printf("Something went wrong when visiting %v: %v\n", e.Request.URL, err)
		<-chs
	})

	c.OnScraped(func(e *colly.Response) {
		defer wg.Done()
		<-chs

		for len(urls) > 0 && i < pageLimit {
			url := urls[0]
			urls = urls[1:]
			visited = append(visited, url)
			i++

			wg.Add(1)
			chs <- struct{}{}
			go c.Visit(url)
		}
	})

	wg.Add(1)
	chs <- struct{}{}
	go c.Visit("https://anime47.com/danh-sach/phim-moi/1.html")
	visited = append(visited, "https://anime47.com/danh-sach/phim-moi/1.html")
	wg.Wait()

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
	href := e.ChildAttr("a", "href")

	if len(href) > 0 {
		return strings.Replace(href, ".", fmt.Sprintf("%s%s", "https://", e.Request.URL.Host), 1)
	}

	return ""
}
