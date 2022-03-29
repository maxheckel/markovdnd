package crawler

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"strings"
	"unicode"
)

var bannedURLPhrases []string = []string{
	"credits",
	"forward",
	"introduction",
	"foreword",
	"appendix",
}


type Crawler interface {
	Crawl() (string, error)
}

type crawler struct {
	Auth string
	BaseURL string
	Collector *colly.Collector
	URLsToCrawl []string
}

func (c crawler) Crawl() (string, error) {
	text := ""
	aloud := ""
	for _, url := range c.URLsToCrawl{
		c.Collector.OnHTML(".p-article-content p, .p-article-content ul", func(e *colly.HTMLElement) {
			text = text + " " + SpaceStringsBuilder(e.Text)
		})
		c.Collector.OnHTML(".read-aloud-text", func(e *colly.HTMLElement) {
			aloud = aloud + " " + SpaceStringsBuilder(e.Text)
		})
		c.Collector.OnRequest(func(request *colly.Request) {
			request.Headers.Add("Cookie", c.Auth)
			request.Headers.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.83 Safari/537.36")
		})
		c.Collector.Visit(url)


	}
	c.Collector.Wait()
	return aloud, nil
}

func NewCrawler(base, auth string) (Crawler, error) {
	c := &crawler{
		Auth:    auth,
		BaseURL: base,
	}
	c.Collector = colly.NewCollector()
	c.Collector.OnHTML(".compendium-toc-full-text > h3 a", func(element *colly.HTMLElement) {
		url := element.Attr("href")
		for _, banned := range bannedURLPhrases{
			if strings.Contains(url, banned) {
				return
			}
		}
		fmt.Println(url)
		c.URLsToCrawl = append(c.URLsToCrawl, url)
	})
	c.Collector.OnRequest(func(request *colly.Request) {
		request.Headers.Add("Cookie", c.Auth)
		request.Headers.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.83 Safari/537.36")
	})
	c.Collector.Visit(c.BaseURL)
	c.Collector.Wait()
	return c, nil
}

func SpaceStringsBuilder(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	prev := rune(0)
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
			prev = ch
		// Tracking the previous rune to ensure that it wasn't a space allows us to write only one space when multiple
		// instances of whitespace occur
		} else if prev != ' '  {
			b.WriteRune(' ')
			prev = ' '
		}

	}
	return b.String()
}
