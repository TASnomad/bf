package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/TASnomad/bf/pkg/compiler"
	"github.com/TASnomad/bf/pkg/vm"
)

func main() {
	var file string
	flag.StringVar(&file, "f", "", "Brainfuck program file")

	flag.Parse()

	if len(file) == 0 {
		fmt.Fprintf(os.Stderr, "Usage: %s:\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	code, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
	c := compiler.NewCompiler(string(code))
	v := vm.NewVm(c.Compile(), os.Stdin, os.Stdout)
	v.Execute()
}
