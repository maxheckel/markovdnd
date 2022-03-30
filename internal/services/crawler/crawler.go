package crawler

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"strings"
)

var bannedURLPhrases = []string{
	"credits",
	"forward",
	"introduction",
	"foreword",
	"appendix",
}

type CrawlerOptions struct {
	Auth string
	BaseURL string
	UseCache bool
}


type Crawler interface {
	Crawl() (BookData, error)
}

type crawler struct {
	Auth string
	BaseURL string
	Collector *colly.Collector
	URLsToCrawl []string
	UseCache bool
}

func (c crawler) Crawl() (BookData, error) {
	res := BookData{
		StoryText: "",
		ReadAloudText: "",
		BaseURL: c.BaseURL,
	}

	if c.UseCache {
		err := res.load()
		return res, err
	}

	for _, url := range c.URLsToCrawl{
		c.Collector.OnHTML(".p-article-content p, .p-article-content ul", func(e *colly.HTMLElement) {
			res.StoryText = res.StoryText + " " + e.Text
		})
		c.Collector.OnHTML(".read-aloud-text", func(e *colly.HTMLElement) {
			res.ReadAloudText = res.ReadAloudText + " " + e.Text
		})

		c.Collector.OnRequest(func(request *colly.Request) {
			request.Headers.Add("Cookie", c.Auth)
			request.Headers.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.83 Safari/537.36")
		})
		c.Collector.Visit(url)


	}

	c.Collector.Wait()
	return res, res.store()
}

func NewCrawler(options CrawlerOptions) (Crawler, error) {
	c := &crawler{
		Auth:    options.Auth,
		BaseURL: options.BaseURL,
		UseCache: options.UseCache,
	}
	if options.UseCache {
		fmt.Println("Skipping craw, cache is enabled!")
		return c, nil
	}

	err := buildURLsToCrawl(c)
	if err != nil {

		return nil, err
	}
	return c, nil
}

func buildURLsToCrawl(c *crawler) error {
	c.Collector = colly.NewCollector()
	c.Collector.OnHTML(".compendium-toc-full-text > h3 a", func(element *colly.HTMLElement) {
		url := element.Attr("href")
		for _, banned := range bannedURLPhrases {
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
	res := c.Collector.Visit(c.BaseURL)

	c.Collector.Wait()
	return res
}
