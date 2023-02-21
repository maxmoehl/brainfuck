package main

import (
	"flag"
	"os"

	"github.com/maxmoehl/brainfuck"
)

var (
	flagDebug       = false
	flagInteractive = false
	flagFile        = ""
)

func init() {
	flag.BoolVar(&flagDebug, "debug", false, "enables the debug mode to print additional information")
	flag.BoolVar(&flagInteractive, "interactive", false, "if -file is given this enables the console after interpreting the file")
	flag.StringVar(&flagFile, "file", "", "path to a file containing brainfuck code")
}

func main() {
	flag.Parse()

	if flagFile == "" {
		brainfuck.RunShell(flagDebug)
	} else {
		readAndRun(flagFile, flagDebug, flagInteractive)
	}
}

func readAndRun(path string, debug bool, interactive bool) {
	code, err := os.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}
	brainfuck.Run(string(code), debug, interactive)
}
