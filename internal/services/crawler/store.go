package crawler

import (
	"github.com/maxheckel/markovdnd/internal/services/store"
	"io/ioutil"
	"strings"
)

const DataPrefix = "data/text/"

type BookData struct {
	StoryText string
	ReadAloudText string
	BaseURL string
}

func (b BookData) store() error{
	err := store.WriteText(DataPrefix+b.Name()+".story.text", b.StoryText)
	if err != nil {
		return err
	}
	return store.WriteText(DataPrefix+b.Name()+".aloud.text", b.ReadAloudText)

}

func (b *BookData) load() error{
	story, err := ioutil.ReadFile(DataPrefix+b.Name()+".story.text")
	if err != nil {
		return err
	}

	aloud, err := ioutil.ReadFile(DataPrefix+b.Name()+".aloud.text")
	if err != nil {
		return err
	}
	b.StoryText = string(story)
	b.ReadAloudText = string(aloud)
	return nil
}



func (b BookData) Name() string{
	return strings.Split(b.BaseURL, "/")[len(strings.Split(b.BaseURL, "/"))-1]
}