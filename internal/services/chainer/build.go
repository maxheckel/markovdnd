package chainer

import (
	"github.com/maxheckel/markovdnd/internal/domain"
)


func Build(c domain.Chain)  {
	textArr := Tokenize(c.Input)
	inputLength := len(textArr)
	for i := 0; i < inputLength; i++{
		if indexOutOfBounds(i+1, inputLength) || indexOutOfBounds(i, inputLength) {
			continue
		}
		word := textArr[i]
		if c.Chain[word] == nil {
			c.Chain[word] = map[string]int{}
		}
		c.Chain[word][textArr[i+1]]++
	}

	for i, word := range textArr{
		if i+5 > len(textArr){
			break
		}
		if len(word) > 0 && word[len(word)-1] == '.' {
			c.StartingWords[textArr[i+1]] = textArr[i+1:i+5]
		}
	}

}

func indexOutOfBounds(index int, length int) bool{
	return index + 1 >= length || index - 1 < 0
}

