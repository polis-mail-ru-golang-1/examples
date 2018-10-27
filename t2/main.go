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
	fmt.Println("Reading files:", os.Args[1:])

	index := index.New()
	for _, file := range os.Args[1:] {
		reed, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Printf("error reading file %q, skip\n", file)
			continue
		}
		index.Add(string(reed), file)
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
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter query: ")
	in, _ := reader.ReadString('\n')
	return in
}
