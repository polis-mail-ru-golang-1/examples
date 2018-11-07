package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/polis-mail-ru-golang-1/examples/t2/index"
	"github.com/polis-mail-ru-golang-1/examples/t2/web"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage of the utility: ./t2 localhost:8080 file1.txt file2.txt")
		return
	}

	files := os.Args[2:]
	fmt.Println("Reading files: ", files)
	index := parallelLoad(files)
	fmt.Printf("%+v\n", index.Info())

	w := web.New(index, os.Args[1])
	w.Start()
}

func parallelLoad(files []string) index.Index {
	indexes := make(chan index.Index, len(files))

	for _, file := range files {
		go func(filename string) {
			index := index.New()
			defer func() {
				indexes <- index
			}()
			reed, err := ioutil.ReadFile(filename)
			if err != nil {
				fmt.Printf("error reading file %q, skip\n", file)
				return
			}
			index.Add(string(reed), filename)
		}(file)
	}

	var index index.Index

	for i := 0; i < len(files); i++ {
		if index == nil {
			index = <-indexes
		} else {
			index.Merge(<-indexes)
		}
	}
	return index
}
