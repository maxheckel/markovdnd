package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/maxheckel/markovdnd/internal/domain"
	"github.com/maxheckel/markovdnd/internal/services/chainer"
	"github.com/maxheckel/markovdnd/internal/services/crawler"
	"github.com/maxheckel/markovdnd/internal/services/images"
	"github.com/maxheckel/markovdnd/internal/services/store"
	"github.com/maxheckel/markovdnd/internal/services/store/chain/drivers"
	"strings"
)


func main(){
	auth := flag.String("auth", "Preferences=undefined; Preferences=undefined; optimizelyEndUserId=oeu1618614554837r0.37376565303291254; ResponsiveSwitch.DesktopMode=1; _rdt_uuid=1618614555432.6099c0d0-ee3c-48e1-a1b3-21a28cc706be; _pxvid=c2f2d121-9f08-11eb-b384-d16d27a49fa9; G_ENABLED_IDPS=google; User=UserID=100459034&UserName=okawei&AuthType=2&ExternalUserID=56966634&ExternalDisplayName=okawei; Preferences.TimeZoneID=1; CobaltSession=eyJhbGciOiJkaXIiLCJlbmMiOiJBMTI4Q0JDLUhTMjU2In0..MfsU07GizGfDi_PhGSCD_A.TRR6m8zQQNBMenfxcW4ZzMALGpncmcj05ALIq9AhewVI4nhJtAxwK5M9USDTxh62.PwUwDT-54GiGQLHceujvzA; sailthru_content=c612f60f905ccac04409c342665f6c55086887834a75211e380c1f6cc01cb5cbf6daf07ff2c20c14e209660c248025c7e330dad3688c8c7c1400e0d5df5721f60deb5d9a01f02fa37b32e8bdd689b89f17199cb57122e519d512ff4a01f6e7fad6d05b03afdcf8966321bdc8101b9973d5ded70a5ae8a043846d565b52f003b9dc825c92a5ccd641d1eccd0cc34071fd7043b5fa210b8295a3fa228d27205f02418d07b58024c07dd05a7019c693392628d28336c31f1eddb3b01976dedbbb8fc728b4eff5058239015426d0876eeb0e8a6f51f2ec5d86947de0cd454b120ae2b4d0999ad34ff27ba26bc41d474d3aa1a45aa54013e60dc8fb6e3290a4e1351d; sailthru_visitor=86ffb421-2909-402a-bb4e-c9e7fd62c8ce; ddbSiteBanner:2af480dd-983d-4c84-9f46-592e6536f99b=true; _pin_unauth=dWlkPU9XVTJZemMwWW1ZdFpqUm1ZUzAwTm1FMkxXSmpaVEF0WlRkak1UUmtOREpqTVdVeg; Ratings=null; ddbSiteBanner:5f53d64b-e6f6-42b6-b58b-245081f1b570=true; ddbSiteBanner:a8605075-5591-451e-9ddd-6f3f81544b2e=true; ddbSiteBanner:6e1322ba-b1dc-4288-9e41-39f4f261626a=true; _hjid=7aaf1174-8244-4014-8559-5f7581b92d50; ddb.toast.magic-item.homebrew-create.hide-toast=true; ddb.toast.magic-item.homebrew-edit.hide-toast=true; ddbSiteBanner:33a7f687-9c2e-4f5c-8589-89ba55cf2d4d=true; fs_uid=rs.fullstory.com#14G1F1#6432490009567232:6405794674155520/1662764442; ddbSiteBanner:59b5c92f-c856-4b2d-91c5-cec2b05d8bd6=true; ddbSiteBanner:dddc53bb-e35f-4514-ab94-d3e3e759b4b4=true; ddbSiteBanner:2fe97b88-c7ad-4748-bf49-8cf86f836ac2=true; ddbSiteBanner:75fae628-5515-48e5-b358-58d9bc594312=true; ddbSiteBanner:00523ca3-81a4-4d2a-8f86-9e40273af2e2=true; _hjSessionUser_2578953=eyJpZCI6Ijc1ZTI0MDc3LTI0NDYtNTc3OS04Nzg0LWVhOTk2ODJjZTFlMSIsImNyZWF0ZWQiOjE2Mzc0NDkwMDE3ODksImV4aXN0aW5nIjp0cnVlfQ==; ddbSiteBanner:34d227e9-4a83-4205-8967-e7f42a77f054=true; LoginState=32f6585d-a705-4747-b178-b48e4e492926; _cc_id=74122c82f8043aa17a47cc719f33f9ec; ddbSiteBanner:201a8bf0-b7ea-41ae-b490-e42061022b2e=true; ddbSiteBanner:a9a9d7eb-1586-4434-97f1-15ac559234b7=true; _gcl_au=1.1.1128065689.1642033900; _pxhd=tqspTd6z7bgFQfXn7I28X5/C15ys2lUHDnyI3dr0kMrGvzNTP-ibcaDR9h9B4aQqtlWiyopNlpQj9lb8tFP7lg==:CbauAWnM4evKLdDrEnLjCTtiik0FI5wO24/wCwmoJtUXIVtcZw5ux2mNQOz6iSU-oybbbB1wMFANpsyOc-aVN/AZAUiCjgR41/wThm/aTAo=; _gid=GA1.2.1981797960.1648510263; panoramaId_expiry=1649115065914; panoramaId=0209dc5905d23a38e4f89764c16e4945a70211c7194df8db7b5c7ee2e0b354b8; _clck=kd55wq|1|f06|0; AWSELB=17A593B6CA59C3C4856B812F84CD401A582EF083467276F6C69E1D56867D29F5D7C31B8EB284A9F525C1AA0DF220CB30AEE9DCF697CC4B9A586C45206F2F6B04BD612C57; AWSELBCORS=17A593B6CA59C3C4856B812F84CD401A582EF083467276F6C69E1D56867D29F5D7C31B8EB284A9F525C1AA0DF220CB30AEE9DCF697CC4B9A586C45206F2F6B04BD612C57; Geo={\"region\":\"OH\",\"country\":\"US\",\"continent\":\"NA\"}; Preferences=undefined; sublevel=MASTER; pxcts=db1e8d95-af71-11ec-b4ce-566c74586e46; cebs=1; _ce.s=v~6e14ac09cc6b3edff49ece264bd774699869e054~vpv~1; RequestVerificationToken=caeacdf1-b0ad-43aa-9ff2-a48f31cdc8e1; _ga_8P5GQ3C7YC=GS1.1.1648565765.143.1.1648568270.60; _ga=GA1.1.1719940889.1618614555; _uetsid=22bc9de0aeef11eca339e743e6a9c1e6; _uetvid=d1b3e680b9ba11eb8ec9f73802bf81cf; _derived_epik=dj0yJnU9aGRpalVRNmU1eU1vZ0c1SExNWTBlUzlybUFUY3JLd3Mmbj1oRVV5Qk42Y0FfQWtTeVB1ZkI0cEVBJm09MSZ0PUFBQUFBR0pESjg4JnJtPTEmcnQ9QUFBQUFHSkRKODg; _clsk=fmxb1i|1648568271595|13|0|e.clarity.ms/collect; _px2=eyJ1IjoiMzE3OTdkNTAtYWY3Ni0xMWVjLTg2Y2ItYjc2ZjA3YjVlNTgyIiwidiI6ImMyZjJkMTIxLTlmMDgtMTFlYi1iMzg0LWQxNmQyN2E0OWZhOSIsInQiOjE2NDg1Njg1NzIyMTgsImgiOiJiODM3YzUwYjMwMmVmYzdhNjI5YjhiMTVmOTg3MmIyN2YyNGM4ZjNhYWU0OTMwMzBjOGUyZGNjZGMwNTZhMGM0In0=", "Auth")
	rootURL := flag.String("root_url", "", "The root URL for the book you would like to train on")
	useCache := flag.Bool("use_cache", true, "Weather or not to use cache")
	flag.Parse()
	fmt.Printf("Crawling %s\n", *rootURL)
	crawler, err := crawler.NewCrawler(crawler.CrawlerOptions{
		Auth:     *auth,
		BaseURL:  *rootURL,
		UseCache: *useCache,
	})

	if err != nil{
		panic(any(err))
	}

	text, err := crawler.Crawl()
	if err != nil{
		panic(any(err))
	}

	fmt.Println("Done! Beginning training")
	storyChain := domain.NewChain(text.StoryText, "story")
	chainer.Build(storyChain)
	readAloudChain := domain.NewChain(text.ReadAloudText, "aloud")
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

	client := images.NewImageDescriber("vision.googleapis.com:443", "dndai-346223-ef10bf69a9d1.json")
	imagesArr := strings.Split(text.Images, " ")
	describerToImage, err := client.GetDescriptionWords(context.Background(), imagesArr)
	if err != nil {
		panic(any(err))
	}
	err = store.WriteJson(drivers.TrainedPrefix+name+".images.json", describerToImage)
	if err != nil {
		panic(any(err))
	}

}