package chainer

import (
	"github.com/maxheckel/auto-dnd/internal/domain"
	"strings"
	"unicode"
)


func Build(c domain.Chain)  {
	c.Input = CleanText(c.Input)
	textArr := strings.Split(c.Input, " ")
	inputLength := len(textArr)
	for i := 0; i < inputLength; i++{
		if indexOutOfBounds(i+1, inputLength) || indexOutOfBounds(i, inputLength) {
			continue
		}

		word := textArr[i]
		if isArticleWord(word, c.ArticleWords) {
			if indexOutOfBounds(i+2, inputLength) || indexOutOfBounds(i+1, inputLength) {
				continue
			}
			// For compound article words i.e. "of the"
			if isArticleWord(word+" "+textArr[i+1], c.ArticleWords){
				if indexOutOfBounds(i+3, inputLength) {
					continue
				}
				word = word + " " + textArr[i+1] + " " + textArr[i+2]
				if c.Chain[word] == nil {
					c.Chain[word] = map[string]int{}
				}
				c.Chain[word][textArr[i+3]]++
				i+=2
				continue
			}
			word = word + " " + textArr[i+1]
			if c.Chain[word] == nil {
				c.Chain[word] = map[string]int{}
			}
			c.Chain[word][textArr[i+2]]++
			// Skip the next word
			i++
			continue
		}



		if isArticleWord(textArr[i+1], c.ArticleWords) {
			if isArticleWord(textArr[i+1] + " " + textArr[i+2], c.ArticleWords){

				if indexOutOfBounds(i+3, inputLength) {
					continue
				}
				if c.Chain[word] == nil {
					c.Chain[word] = map[string]int{}
				}
				c.Chain[word][textArr[i+1] + " " + textArr[i+2] + " " + textArr[i+3]]++
				i+=2
				continue
			}

			if indexOutOfBounds(i+2, inputLength) {
				continue
			}
			if c.Chain[word] == nil {
				c.Chain[word] = map[string]int{}
			}
			i++
			c.Chain[word][textArr[i+1] + " " + textArr[i+2]]++
			continue
		}
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

func isArticleWord(word string, rules []string) bool{
	for _, check := range rules{
		if check == word{
			return true
		}
	}
	return false
}


func CleanText(str string) string {
	str = strings.ToLower(str)
	var b strings.Builder
	b.Grow(len(str))
	prev := rune(0)
	for _, ch := range str {
		if isForbiddenRune(ch) {
			continue
		}
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
			prev = ch
			// Tracking the previous rune to ensure that it wasn't a space allows us to write only one space when multiple
			// instances of whitespace occur
		} else if prev != ' '  {
			b.WriteRune(' ')
			prev = ' '
		}

	}
	return b.String()
}

func isForbiddenRune(ch rune) bool {
	forbiddenRunes := []rune{
		'"',
		'“',
		'”',
		'(',
		')',
	}

	for _, r := range forbiddenRunes{
		if ch == r{
			return true
		}
	}
	return false
}