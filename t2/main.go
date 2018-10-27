package main

import (
	"bufio"
	"fmt"
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

	index := readingLoad(files, 100000)

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

func readingLoad(files []string, buffer int) index.Index {
	index := index.New(buffer)
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			fmt.Printf("error reading file %q, skip\n", file)
			continue
		}
		index.Read(bufio.NewReader(f), file)
	}
	index.Wait()
	return index
}

func readQuery() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter query: ")
	in, _ := reader.ReadString('\n')
	return in
}
