package main

import (
	"flag"
	"fmt"
	"github.com/maxheckel/auto-dnd/internal/domain"
	"github.com/maxheckel/auto-dnd/internal/services/chainer"
	"github.com/maxheckel/auto-dnd/internal/services/crawler"
	"github.com/maxheckel/auto-dnd/internal/services/store"
	"github.com/maxheckel/auto-dnd/internal/services/store/chain/drivers"
)


func main(){
	auth := flag.String("auth", "", "Auth")
	rootURL := flag.String("root_url", "https://www.dndbeyond.com/sources/cos", "The root URL for the book you would like to train on")
	flag.Parse()
	fmt.Printf("Crawling %s\n", *rootURL)
	crawler, err := crawler.NewCrawler(crawler.CrawlerOptions{
		Auth:     *auth,
		BaseURL:  *rootURL,
		UseCache: true,
	})

	if err != nil{
		panic(any(err))
	}

	text, err := crawler.Crawl()
	if err != nil{
		panic(any(err))
	}

	fmt.Println("Done! Beginning training")
	articles := []string{
		"a",
		"the",
		"an",
	}
	storyChain := domain.NewChain(text.StoryText, articles, "story")
	chainer.Build(storyChain)
	readAloudChain := domain.NewChain(text.ReadAloudText, articles, "aloud")
	chainer.Build(readAloudChain)

	fmt.Println("Training Done! Writing files in this directory")
	name := text.Name()

	err = store.WriteJson(drivers.TrainedPrefix+name+".story.json", storyChain)
	if err != nil {
		panic(any(err))
	}
	err = store.WriteJson(drivers.TrainedPrefix+name+".aloud.json", readAloudChain)
	if err != nil {
		panic(any(err))
	}
}