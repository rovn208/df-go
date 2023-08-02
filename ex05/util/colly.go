package util

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

// GetAnimeURL returns the anime url from a hyperlink element
func GetAnimeURL(element *colly.HTMLElement) string {
	href := element.ChildAttr("a", "href")

	if len(href) > 0 {
		return strings.Replace(href, ".", fmt.Sprintf("%s%s", "https://", element.Request.URL.Host), 1)
	}

	return ""
}
