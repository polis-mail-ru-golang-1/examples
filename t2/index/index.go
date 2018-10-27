package index

import (
	"strings"
)

type Index map[word]occurrences
type occurrences map[file]count
type word string
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
	rawWords := strings.Fields(data)
	for _, rawWord := range rawWords {
		newWord := cleanWord(rawWord)
		i.addWord(newWord, file(filename))
	}
	return nil
}

func (i Index) addWord(newWord word, filename file) {
	if _, ok := i[newWord]; !ok {
		i[newWord] = occurrences{}
	}
	os := i[newWord]
	os.addFile(filename)
}

func (os occurrences) addFile(filename file) {
	if _, ok := os[filename]; !ok {
		os[filename] = 1
	} else {
		os[filename]++
	}
}

func cleanWord(in string) word {
	in = strings.ToLower(in)
	in = strings.TrimFunc(in, func(c rune) bool {
		// not [0-9] and not [a-z] and not \-
		return (c < 48 || c > 57) && (c < 97 || c > 122) && c != 45
	})
	return word(in)
}

func (i Index) Search(query string) ([]Result, error) {
	rawWords := strings.Split(query, " ")
	os := occurrences{}
	for _, rawWord := range rawWords {
		qWord := cleanWord(rawWord)
		os = i.searchWord(qWord, os)
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

func (i Index) searchWord(qWord word, previous occurrences) occurrences {
	if _, ok := i[qWord]; !ok {
		return previous
	}
	os := occurrences{}
	for filename, count := range i[qWord] {
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
