package main

import (
	spew "github.com/davecgh/go-spew/spew"
	flags "github.com/jessevdk/go-flags"
)

var opts struct {
	NotSet   []string `long:"notset"`
	Empty    []string `long:"empty"`
	Single   []string `long:"single"`
	Multiple []string `long:"multiple"`
}

func main() {
	args := []string{
		"--empty=",
		"--single=string",
		"--multiple=stringA,stringB,stringC",
	}
	flags.ParseArgs(&opts, args)
	spew.Dump(opts)
}
