package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var inFile, outFile string
	fmt.Scanln(&inFile)
	fmt.Scanln(&outFile)

	f, err := os.Open(inFile)
	die(err)

	s, err := f.Stat()
	die(err)
	data := make([]byte, s.Size())
	_, err = f.Read(data)
	die(err)

	splitted := strings.Split(string(data), ",")
	numbers := string2int(splitted)

	numbers = sort(numbers)
	data = []byte(fmt.Sprintf("%v", numbers))

	f, err = os.Create(outFile)
	die(err)
	_, err = f.Write(data)
	die(err)
}

func string2int(in []string) []int {
	numbers := make([]int, 0, len(in))
	for _, item := range in {
		num, err := strconv.Atoi(strings.TrimSpace(item))
		if err == nil {
			numbers = append(numbers, num)
		} else {
			fmt.Printf("Illegal string: %q, skip\n", item)
		}
	}
	return numbers
}

func sort(in []int) []int {
	out := make([]int, len(in))
	copy(out, in)
	for i := 0; i < len(out); i++ {
		for j := 0; j < len(out)-i-1; j++ {
			if out[j] > out[j+1] {
				out[j], out[j+1] = out[j+1], out[j]
			}
		}
	}
	fmt.Println(out)
	return out
}

func die(err error) {
	if err != nil {
		panic(err)
	}
}
