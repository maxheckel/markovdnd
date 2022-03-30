package chainer

import (
	"github.com/maxheckel/auto-dnd/internal/domain"
	"math/rand"
	"strings"
)


func Run(chain *domain.Chain, numWords int) (string, error){
	randomStartWord := randomElementOfMap[[]string](chain.StartingWords)
	text := []string{randomStartWord}
	for i := 1; i < numWords; i++{
		text = append(text, getNextWord(text[i-1], chain))
	}
	return strings.Join(text, " "), nil
}

func getNextWord(seed string, chain *domain.Chain) string{
	options := chain.Chain[seed]
	if len(options) > 0 {
		return randomElementOfMap(options)
	}
	return ""
}

func randomElementOfMap[V any](input map[string]V) string{
	max := len(input)
	random := rand.Intn(max - 0) + 0
	i := 0
	for key, _ := range input {
		if i == random{
			return key
		}
		i++
	}
	return ""

}