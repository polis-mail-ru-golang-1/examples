package index

import (
	"strings"
)

type Index map[token]occurrences
type occurrences map[file]count
type token string
type file string
type count int

type Result struct {
	File  string
	Count int
}

type Stat struct {
	Count int
}

func New() Index {
	return Index{}
}

func (i Index) Add(data, filename string) error {
	words := strings.Fields(data)
	for _, word := range words {
		token := cleanWord(word)
		i.addToken(token, file(filename))
	}
	return nil
}

func (i Index) addToken(newToken token, filename file) {
	if _, ok := i[newToken]; !ok {
		i[newToken] = occurrences{}
	}
	os := i[newToken]
	os.addFile(filename)
}

func (os occurrences) addFile(filename file) {
	os[filename]++
}

func cleanWord(in string) token {
	in = strings.ToLower(in)
	in = strings.TrimFunc(in, func(c rune) bool {
		// not [0-9] and not [a-z] and not \-
		return (c < 48 || c > 57) && (c < 97 || c > 122) && c != 45
	})
	return token(in)
}

func (i Index) Search(query string) ([]Result, error) {
	words := strings.Split(query, " ")
	os := occurrences{}
	for _, word := range words {
		token := cleanWord(word)
		os = i.searchWord(token, os)
	}

	results := make([]Result, 0, len(os))
	for filename, count := range os {
		results = append(results, Result{
			File:  string(filename),
			Count: int(count),
		})
	}
	return results, nil
}

func (i Index) searchWord(qToken token, previous occurrences) occurrences {
	if _, ok := i[qToken]; !ok {
		return previous
	}
	os := occurrences{}
	for filename, count := range i[qToken] {
		if in(previous, filename) || len(previous) == 0 {
			os[filename] = previous[filename] + count
		}
	}
	return os
}

func in(os occurrences, filename file) bool {
	for idxFilename := range os {
		if idxFilename == filename {
			return true
		}
	}
	return false
}

func (i Index) Info() Stat {
	return Stat{
		Count: len(i),
	}
}
