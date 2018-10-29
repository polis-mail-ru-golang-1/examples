package console

import (
	"bufio"
	"fmt"
	"github.com/polis-mail-ru-golang-1/examples/webservice/index"
	"os"
)

type Console struct {
	index index.Index
}

func New(index index.Index) Console {
	return Console{
		index: index,
	}
}

func (c Console) Start() error {
	for {
		q := c.readQuery()
		results := c.index.Search(q)
		for _, result := range results {
			fmt.Printf("%s -- %d\n", result.Name, result.Score)
		}
	}
	return nil
}

func (c Console) readQuery() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter query: ")
	in, _ := reader.ReadString('\n')
	return in
}
