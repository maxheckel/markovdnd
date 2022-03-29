package main

import (
	"flag"
	"fmt"
	"github.com/maxheckel/auto-dnd/internal/domain"
	"github.com/maxheckel/auto-dnd/internal/services/chainer"
	"github.com/maxheckel/auto-dnd/internal/services/store"
	"strings"

	"github.com/maxheckel/auto-dnd/internal/services/crawler"
)

func main(){
	auth := flag.String("auth", "", "Auth")
	rootURL := flag.String("root_url", "https://www.dndbeyond.com/sources/cos", "The root URL for the book you would like to train on")
	flag.Parse()
	fmt.Printf("Crawling %s", *rootURL)
	crawler, err := crawler.NewCrawler(*rootURL, *auth)

	if err != nil{
		panic(any(err))
	}

	text, err := crawler.Crawl()
	fmt.Println("Done! Beginning training")
	if err != nil{
		panic(any(err))
	}

	articles := []string{
		"a",
		"the",
		"an",
		"of",
		"of the",
	}
	storyChain := domain.NewChain(text.StoryText, articles, "story")
	chainer.Build(storyChain)
	readAloudChain := domain.NewChain(text.ReadAloudText, articles, "aloud")
	chainer.Build(readAloudChain)

	fmt.Println("Training Done! Writing files in this directory")
	url := strings.Split(*rootURL, "/")[len(strings.Split(*rootURL, "/"))-1]
	err = store.WriteJson(url+".story.json", storyChain)
	if err != nil {
		panic(any(err))
	}
	err = store.WriteJson(url+".aloud.json", readAloudChain)
	if err != nil {
		panic(any(err))
	}
}