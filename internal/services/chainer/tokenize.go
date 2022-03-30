package chainer

import (
	"strings"
	"unicode"
)

func Tokenize(input string) []string{
	input = cleanText(input)
	output := []string{}
	rawSplits := strings.Split(input, " ")
	for i := 0; i < len(rawSplits); i++{
		if i+1 > len(rawSplits) {
			break
		}
		word := rawSplits[i]
		for isArticleWord(word) {
			i++
			word += " " + rawSplits[i]
		}
		output = append(output, word)
	}
	return output
}

func cleanText(str string) string {
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


func isArticleWord(word string) bool{
	rules := []string{
		"a",
		"the",
		"an",
		"and",
		"of",
		"is",
		"of the",
		"and the",
		"is of",
	}
	for _, check := range rules{
		if check == word{
			return true
		}
	}
	return false
}

