package scraper

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/rovn208/df-go/ex05/util"
	"sync"
)

type AnimeScraper struct {
	userAgent       string
	numberOfChannel int
	startUrl        string
	limit           int
	wg              sync.WaitGroup
	chs             chan interface{}
	animeList       []Anime
	visitedUrls     []string
	urls            []string
	counter         int
	collector       *colly.Collector
}

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

func (s *AnimeScraper) Execute() {
	s.visit(s.startUrl)
	s.wg.Wait()
	fmt.Printf("Anime scraped: %d", len(s.animeList))
}

func (s *AnimeScraper) addUrl(url string) {
	s.urls = append(s.urls, url)
}

func (s *AnimeScraper) addAnime(anime Anime) {
	s.animeList = append(s.animeList, anime)
}

func (s *AnimeScraper) increaseCount() {
	s.counter++
}

func (s *AnimeScraper) addVisitedUrl(url string) {
	s.visitedUrls = append(s.visitedUrls, url)
}

func (s *AnimeScraper) addGoRoutine() {
	s.wg.Add(1)
	s.chs <- struct{}{}
}

func (s *AnimeScraper) visit(url string) {
	s.addVisitedUrl(url)
	s.increaseCount()
	s.addGoRoutine()

	go s.collector.Visit(url)
}
