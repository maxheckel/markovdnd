package chainer

import (
	"github.com/maxheckel/auto-dnd/internal/domain"
	"math/rand"
	"strings"
	"unicode"
)


func Run(chain *domain.Chain, numWords int) (string, error){
	randomStartWord := randomElementOfMap[[]string](chain.StartingWords)
	text := []string{randomStartWord}
	i := 0
	for i = 1; i < numWords; i++{
		wordToAdd := getNextWord(text[i-1], chain)
		text = append(text, wordToAdd)
	}
	//End the final sentence
	nextWord := getNextWord(text[i - 1], chain)
	for !strings.Contains(nextWord, ".") {
		nextWord = getNextWord(text[i-1], chain)
		text = append(text, nextWord)
		i++
	}

	// Adding some basic capitalization
	text[0] = capitalize(text[0])
	for i = 1; i < numWords; i++{
		if strings.Contains(text[i-1], ".") {
			text[i] = capitalize(text[i])
		}
	}

	return strings.Join(text, " "), nil
}


// So roundabout but this is fastest
func capitalize(str string) string {
	str = strings.ToLower(str)
	var b strings.Builder
	b.Grow(len(str))
	for i, ch := range str {
		if i == 0{
			b.WriteRune(unicode.ToUpper(ch))
			continue
		}
		b.WriteRune(ch)
	}
	return b.String()
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