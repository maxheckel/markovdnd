package drivers

import (
	"encoding/json"
	"github.com/maxheckel/markovdnd/internal/domain"
	"io/ioutil"
)

const TrainedPrefix = "data/trained/"

type FilesystemDriver struct {
	loadedChains map[string][]*domain.Chain
}


func (t *FilesystemDriver) LoadChain(name string) error {
	if t.loadedChains == nil {
		t.loadedChains = map[string][]*domain.Chain{}
	}
	if len(t.loadedChains[name]) > 0 {
		return nil
	}
	story, err := ioutil.ReadFile(TrainedPrefix+name+".story.json")
	if err != nil {
		return err
	}
	aloud, err := ioutil.ReadFile(TrainedPrefix+name+".aloud.json")
	if err != nil {
		return err
	}
	storyChain := &domain.Chain{}
	aloudChain := &domain.Chain{}
	err = json.Unmarshal(story, &storyChain)
	if err != nil{
		return err
	}
	err = json.Unmarshal(aloud, &aloudChain)
	if err != nil{
		return err
	}
	t.loadedChains[name] = []*domain.Chain{
		storyChain,
		aloudChain,
	}
	return nil
}

func (t *FilesystemDriver) GetChains(name string) ([]*domain.Chain, error) {
	return t.loadedChains[name], nil
}


