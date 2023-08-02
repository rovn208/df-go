package main

import "github.com/rovn208/df-go/ex05/scraper"

func main() {
	animeScraper := scraper.NewAnimeScraper("https://anime47.com/danh-sach/phim-moi/1.html", 3, 40)
	animeScraper.Execute()
}
