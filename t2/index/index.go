package index

import (
	"github.com/reiver/go-porterstemmer"
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

func (i Index) Merge(i2 Index) {
	for tkn2, os2 := range i2 {
		if _, ok := i[tkn2]; !ok {
			i[tkn2] = os2
		} else {
			i[tkn2].merge(os2)
		}
	}
}

func (os occurrences) merge(os2 occurrences) {
	for filename2, count2 := range os2 {
		os[filename2] += count2
	}
}

func (i Index) Add(data, filename string) error {
	words := strings.Fields(data)
	for _, word := range words {
		tkn := cleanWord(word)
		i.addToken(tkn, file(filename))
	}
	return nil
}

func (i Index) addToken(newTkn token, filename file) {
	if _, ok := i[newTkn]; !ok {
		i[newTkn] = occurrences{}
	}
	os := i[newTkn]
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
	in = porterstemmer.StemString(in)
	return token(in)
}

func (i Index) Search(query string) ([]Result, error) {
	words := strings.Split(query, " ")
	os := occurrences{}
	for _, word := range words {
		tkn := cleanWord(word)
		os = i.searchToken(tkn, os)
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

func (i Index) searchToken(tkn token, prev occurrences) occurrences {
	if _, ok := i[tkn]; !ok {
		return nil
	}
	os := occurrences{}
	for filename, count := range i[tkn] {
		if in(prev, filename) || len(prev) == 0 {
			os[filename] = prev[filename] + count
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
