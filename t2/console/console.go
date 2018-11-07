package console

import (
	"bufio"
	"fmt"
	"github.com/polis-mail-ru-golang-1/examples/t2/index"
	"os"
)

// Console provides cli interface
type Console struct {
	index index.Index
}

// New console iface
func New(index index.Index) Console {
	return Console{
		index: index,
	}
}

// Start console interaction
func (c Console) Start() error {
	for {
		q := c.readQuery()
		results, err := c.index.Search(q)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			continue
		}
		for _, result := range results {
			fmt.Printf("%s -- %d\n", result.File, result.Count)
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
