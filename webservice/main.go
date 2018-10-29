package main

import (
	"github.com/polis-mail-ru-golang-1/examples/webservice/console"
	"github.com/polis-mail-ru-golang-1/examples/webservice/index"
	"github.com/polis-mail-ru-golang-1/examples/webservice/web"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		panic("Not enough arguments")
	}

	address := os.Args[1]
	files := os.Args[2:]

	idx := index.New()
	for _, file := range files {
		idx.AddFile(file)
	}

	w := web.New(idx, address)
	die(w.Start())

	console.New(idx) // line to omit not used variables
	// c := console.New(idx)
	// die(c.Start())
}

func die(err error) {
	if err != nil {
		panic(err)
	}
}
