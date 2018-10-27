package index

import (
	"bufio"
	"github.com/reiver/go-porterstemmer"
	"io"
	"strings"
	"sync"
)

type Index struct {
	idx    IndexMap
	tokenc chan pair
	wg     *sync.WaitGroup
}
type IndexMap map[token]occurrences
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

type pair struct {
	Filename file
	Token    token
}

func New(buffer int) Index {
	idx := Index{
		idx:    IndexMap{},
		tokenc: make(chan pair, buffer),
		wg:     &sync.WaitGroup{},
	}
	go idx.readTokens()
	return idx
}

func (i Index) Wait() {
	i.wg.Wait()
}

func (i Index) Read(data *bufio.Reader, filename string) error {
	i.wg.Add(1)
	go func(r *bufio.Reader, filename string) {
		defer i.wg.Done()
		for {
			raw, err := r.ReadString(' ')
			if err == io.EOF && raw == "" {
				break
			}
			words := strings.Split(raw, "\n")
			for _, word := range words {
				tkn := cleanWord(word)
				i.tokenc <- pair{
					Filename: file(filename),
					Token:    tkn,
				}
			}
		}
	}(data, filename)
	return nil
}

func (i Index) readTokens() {
	for p := range i.tokenc {
		i.addToken(p.Token, p.Filename)
	}
}

func (i Index) addToken(newTkn token, filename file) {
	if _, ok := i.idx[newTkn]; !ok {
		i.idx[newTkn] = occurrences{}
	}
	os := i.idx[newTkn]
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
	if _, ok := i.idx[tkn]; !ok {
		return nil
	}
	os := occurrences{}
	for filename, count := range i.idx[tkn] {
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
		Count: len(i.idx),
	}
}
