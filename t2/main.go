package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage of the utility: ./t2 file1.txt file2.txt")
		return
	}
	fmt.Println("Reading files:", os.Args[1:])
}
