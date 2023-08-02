package scraper

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/rovn208/df-go/ex05/util"
	"sync"
)

type AnimeScraper struct {
	userAgent       string           // User-Agent string used by HTTP requests
	numberOfChannel int              // Number of channel to scrape at a time
	startUrl        string           // Start url
	limit           int              // Number of pages limit to scrape at a time
	wg              sync.WaitGroup   // Wait group
	chs             chan interface{} // A buffered channel for scraping jobs
	animeList       []Anime          // List of anime
	visitedUrls     []string         // List of visited urls
	urls            []string         // List of urls
	counter         int              // Number of pages scraped
	collector       *colly.Collector // Colly Collector
}

// NewAnimeScraper returns a new AnimeScraper
func NewAnimeScraper(startUrl string, numberOfChannel, pageLimit int) *AnimeScraper {
	scraper := &AnimeScraper{
		userAgent:   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36",
		startUrl:    startUrl,
		limit:       pageLimit,
		chs:         make(chan interface{}, numberOfChannel),
		animeList:   make([]Anime, 0),
		visitedUrls: make([]string, 0),
		urls:        make([]string, 0),
		collector:   colly.NewCollector(),
	}
	scraper.setup()

	return scraper
}

// setup setups the scraper
func (s *AnimeScraper) setup() {
	s.collector.UserAgent = s.userAgent

	s.collector.OnRequest(func(e *colly.Request) {
		fmt.Printf("visiting %v\n", e.URL)
	})

	s.collector.OnHTML("ul.pagination > li:not(.active)", func(e *colly.HTMLElement) {
		url := util.GetAnimeURL(e)
		if len(url) > 0 && !util.Contains(s.visitedUrls, url) {
			s.addUrl(url)
		}
	})

	s.collector.OnHTML("div.movie-item", func(e *colly.HTMLElement) {
		anime := Anime{
			Title:   e.ChildAttr("a", "title"),
			URL:     util.GetAnimeURL(e),
			Episode: e.ChildAttr("span.ribbon", "text"),
		}
		s.addAnime(anime)
	})

	s.collector.OnError(func(e *colly.Response, err error) {
		defer s.wg.Done()
		<-s.chs
		fmt.Printf("Something went wrong when visiting %v: %v\n", e.Request.URL, err)
	})

	s.collector.OnScraped(func(e *colly.Response) {
		defer s.wg.Done()
		<-s.chs

		for len(s.urls) > 0 && s.counter < s.limit {
			url := s.urls[0]
			s.urls = s.urls[1:]

			if !util.Contains(s.visitedUrls, url) {
				s.visit(url)
			}
		}
	})
}

// Execute runs the scraper with the start url
func (s *AnimeScraper) Execute() {
	s.visit(s.startUrl)
	s.wg.Wait()
	fmt.Printf("Anime scraped: %d", len(s.animeList))
}

// addUrl adds a url to the list of urls
func (s *AnimeScraper) addUrl(url string) {
	s.urls = append(s.urls, url)
}

// addAnime adds an anime to the list of anime
func (s *AnimeScraper) addAnime(anime Anime) {
	s.animeList = append(s.animeList, anime)
}

// increaseCount increases the counter
func (s *AnimeScraper) increaseCount() {
	s.counter++
}

// addVisitedUrl adds a url to the list of visited urls
func (s *AnimeScraper) addVisitedUrl(url string) {
	s.visitedUrls = append(s.visitedUrls, url)
}

// addGoRoutine adds a goroutine to the wait group and the channel
func (s *AnimeScraper) addGoRoutine() {
	s.wg.Add(1)
	s.chs <- struct{}{}
}

// visit visits a url
func (s *AnimeScraper) visit(url string) {
	s.addVisitedUrl(url)
	s.increaseCount()
	s.addGoRoutine()

	go s.collector.Visit(url)
}
