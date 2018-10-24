package splitspace

import (
	"regexp"
	"strings"
)

var (
	gre = regexp.MustCompile("[\\w]+")
)

func splitTrim(in string) []string {
	words := strings.Split(in, " ") //выделение слов и удаление знаков препинания
	for i := 0; i < len(words); i++ {
		words[i] = strings.ToLower(words[i])
		words[i] = strings.TrimFunc(words[i], func(r rune) bool {
			return ((r >= 0 && r <= 64) || (r >= 91 && r <= 96) || (r >= 123))
		})
		if words[i] == "" {
			words = append(words[:i], words[i+1:]...)
		}
	}
	return words
}

func tokenize(str string) []string {
	re := regexp.MustCompile("[\\w]+")
	tokenPositions := re.FindAllStringIndex(str, -1)
	tokens := make([]string, len(tokenPositions))

	for i, pos := range tokenPositions {
		tokens[i] = strings.ToLower(str[pos[0]:pos[1]])
	}

	return tokens
}

func tokenizeGlobal(str string) []string {
	tokenPositions := gre.FindAllStringIndex(str, -1)
	tokens := make([]string, len(tokenPositions))

	for i, pos := range tokenPositions {
		tokens[i] = strings.ToLower(str[pos[0]:pos[1]])
	}

	return tokens
}
