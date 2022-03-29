package main

import (
	"fmt"
	"github.com/maxheckel/auto-dnd/internal/services/crawler"
)

func main(){
	crawler, err := crawler.NewCrawler("https://www.dndbeyond.com/sources/cos", "")
	if err != nil{
		panic(any(err))
	}
	text, err := crawler.Crawl()
	if err != nil{
		panic(any(err))
	}
	fmt.Println(text)
}