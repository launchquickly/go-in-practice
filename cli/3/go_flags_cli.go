package main

import (
	"fmt"

	flags "github.com/jessevdk/go-flags"
)

var opts struct {
	Name    string `short:"n" long:"name" description:"A name to say hello to." default:"World"`
	Spanish bool   `short:"s" long:"spanish" description:"Use Spanish language."`
}

func main() {
	flags.Parse(&opts)

	if opts.Spanish {
		fmt.Printf("Hola %s!\n", opts.Name)
	} else {
		fmt.Printf("Hello %s!\n", opts.Name)
	}
}
