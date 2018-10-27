package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/polis-mail-ru-golang-1/examples/t2/index"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage of the utility: ./t2 file1.txt file2.txt")
		return
	}

	files := os.Args[1:]
	fmt.Println("Reading files: ", files)

	// index := singleLoad(files)
	index := parallelLoad(files)

	fmt.Printf("%+v\n", index.Info())

	for {
		q := readQuery()
		results, err := index.Search(q)
		if err != nil {
			fmt.Printf("error searching: %q\n", err)
			continue
		}
		for _, result := range results {
			fmt.Printf("%s -- %d\n", result.File, result.Count)
		}
	}
}

func singleLoad(files []string) index.Index {
	index := index.New()
	for _, file := range files {
		reed, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Printf("error reading file %q, skip\n", file)
			continue
		}
		index.Add(string(reed), file)
	}
	return index
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

func readQuery() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter query: ")
	in, _ := reader.ReadString('\n')
	return in
}
