package run

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/maxheckel/markovdnd/internal/domain"
	"github.com/maxheckel/markovdnd/internal/services/chainer"
	"github.com/maxheckel/markovdnd/internal/services/store/chain"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

type Chain struct {
	Store chain.Store
}

func (rc Chain) ServeHTTP(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	name := params["name"]
	numStory := 10
	numAloud := 10
	var err error
	if r.URL.Query().Get("num_story") != ""{
		numStory, err = strconv.Atoi(r.URL.Query().Get("num_story"))
	}
	if r.URL.Query().Get("num_aloud") != ""{
		numAloud, err = strconv.Atoi(r.URL.Query().Get("num_aloud"))
	}
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	err = rc.Store.LoadChain(name)
	if err != nil {
		w.Write([]byte("Could not find chain for source "+name))
		return
	}
	chains, err := rc.Store.GetChains(name)
	if err != nil {
		w.Write([]byte("Could not find chain for source "+name))
		return
	}
	err = rc.Store.LoadImages(name)
	if err != nil {
		w.Write([]byte("Could not find images for source "+name))
		return
	}
	images, err := rc.Store.GetImages(name)
	if err != nil {
		w.Write([]byte("Could not find images for source "+name))
		return
	}


	resp := domain.Generated{
		Story:     []string{},
		ReadAloud: []string{},
	}
	for range make([]int, numStory){
		paragraph, err := chainer.Run(chains[0], 100)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		resp.Story = append(resp.Story, paragraph)
	}
	for range make([]int, numAloud){
		paragraph, err := chainer.Run(chains[1], 100)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		resp.ReadAloud = append(resp.ReadAloud, paragraph)
	}

	rc.AddImagesToResponse(resp.ReadAloud, images, &resp, "aloud")
	rc.AddImagesToResponse(resp.Story, images, &resp, "story")

	output, err := json.Marshal(resp)
	w.Write(output)

}

func (rc Chain) AddImagesToResponse(sentences []string, images *domain.Images, resp *domain.Generated, wordsType string) {
TOP:
	for index, strand := range sentences {
		words := strings.Split(strand, " ")
		for _, word := range words {
			if images.ImageMap[word] != nil {
				resp.Images = append(resp.Images, domain.ImageWithPosition{
					// Random element of the array
					URL:      images.ImageMap[word][rand.Intn(len(images.ImageMap[word]))],
					Type:     wordsType,
					Position: index,
				})
				continue TOP
			}
		}
	}
}