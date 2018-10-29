package main

import (
	"fmt"
	"os"

	"github.com/polis-mail-ru-golang-1/examples/webservice/console"
	"github.com/polis-mail-ru-golang-1/examples/webservice/index"
	"github.com/polis-mail-ru-golang-1/examples/webservice/web"
)

func main() {
	if len(os.Args) < 3 {
		panic("Not enough arguments")
	}

	address := os.Args[1]
	files := os.Args[2:]

	idx := index.New()
	for _, file := range files {
		content, err := readFile(file)
		if err != nil {
			fmt.Println("Error reading", file)
			continue
		}
		idx.AddFile(file, content)
	}

	w := web.New(idx, address)
	die(w.Start())

	console.New(idx) // line to omit not used variables
	// c := console.New(idx)
	// die(c.Start())
}

func readFile(name string) (string, error) {
	// reading file here
	return `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam orci dui, mattis eget dignissim a, sollicitudin
			quis felis. Phasellus imperdiet, diam vel fermentum convallis, massa erat hendrerit ligula, vitae
			sollicitudin nisl augue et tortor. Etiam id scelerisque felis. Nulla eget neque libero. Pellentesque ipsum
			est, dignissim at est tempor, volutpat faucibus tellus. Mauris in cursus lorem, id sagittis eros. Donec
			posuere eleifend nisi, nec tristique dolor luctus at. Sed ipsum elit, hendrerit at eros nec, tincidunt
			finibus mauris. Suspendisse ullamcorper quis lorem vitae dictum. Fusce auctor erat neque, congue lobortis
			tortor viverra vel. Proin leo augue, faucibus vel ex quis, posuere facilisis enim. Morbi commodo erat lorem,
			sed sagittis augue viverra facilisis. Integer consequat dui quam, vitae hendrerit nunc dapibus in. Aenean
			pharetra ipsum non velit sodales, sed lobortis libero mollis. Nullam facilisis magna libero, ac facilisis
			neque vehicula in.`, nil
}

func die(err error) {
	if err != nil {
		panic(err)
	}
}
