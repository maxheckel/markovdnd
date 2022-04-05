package drivers

import (
	"encoding/json"
	"github.com/maxheckel/markovdnd/internal/domain"
	"io/ioutil"
	"strings"
)

const TrainedPrefix = "data/trained/"

type FilesystemDriver struct {
	loadedChains map[string][]*domain.Chain
	loadedImages map[string]*domain.Images
}

func (t *FilesystemDriver) LoadImages(name string) error {
	if t.loadedImages == nil {
		t.loadedImages = map[string]*domain.Images{}
	}
	if t.loadedImages[name] != nil{
		return nil
	}
	images, err := ioutil.ReadFile(TrainedPrefix+name+".images.json")
	if err != nil {
		return err
	}
	imagesOutput := domain.Images{}

	err = json.Unmarshal(images, &imagesOutput.ImageMap)
	if err != nil{
		return err
	}
	t.loadedImages[name] = &imagesOutput
	return nil
}

func (t *FilesystemDriver) GetImages(name string) (*domain.Images, error) {
	return t.loadedImages[name], nil
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

func (t *FilesystemDriver) GetCrawled() ([]string, error) {
	files, err := ioutil.ReadDir(TrainedPrefix)
	if err != nil {
		return nil, err
	}

	var names []string
	namesMap := map[string]bool{}
	for _, file := range files {
		nameArr := strings.Split(file.Name(), ".")
		namesMap[nameArr[0]] = true
	}
	for name := range namesMap{
		names = append(names, name)
	}
	return names, nil


}


