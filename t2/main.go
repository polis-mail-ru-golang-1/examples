package main

import (
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
	fmt.Println("Reading files:", os.Args[1:])

	index := index.New()
	for _, file := range os.Args[1:] {
		readed, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Printf("error reading file %q, skip\n", file)
			continue
		}
		index.Add(string(readed), file)
	}

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

func readQuery() string {
	fmt.Print("Enter query: ")
	var in string
	fmt.Scanln(&in)
	return in
}
