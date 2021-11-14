package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/TASnomad/bf/pkg/compiler"
	"github.com/TASnomad/bf/pkg/vm"
)

var Build = "unknown"
var Version = "unknown"

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

func main() {

	flag.Usage = func() {
		usage := fmt.Sprintf("Usage: %s [option] [arg]\n\nBrainfuck VM\n\nOptions and arguments:", os.Args[0])
		fmt.Fprint(flag.CommandLine.Output(), usage)
		flag.PrintDefaults()
	}

	var file string
	var showExecTime bool = false
	flag.StringVar(&file, "f", "", "Brainfuck program file")
	ver := flag.Bool("v", false, "Print the version number and exit")
	perf := flag.Bool("perf", false, "Print VM execution time")

	flag.Parse()

	if *ver {
		fmt.Printf("%s@%s", Version, Build)
		os.Exit(0)
	}

	if *perf {
		showExecTime = true
	}

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
	if showExecTime {
		defer elapsed("bf VM execution")()
	}
	v.Execute()
}
