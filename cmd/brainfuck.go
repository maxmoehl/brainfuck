package main

import (
	"io/ioutil"
	"os"

	"github.com/maxmoehl/brainfuck"
	"github.com/mkideal/cli"
)

type Args struct {
	cli.Helper
	Debug       bool   `cli:"d,debug" usage:"enables the debug mode to print additional information" dft:"false"`
	Interactive bool   `cli:"i,interactive" usage:"if a file is given this enables the console after interpreting the file" dft:"false"`
	File        string `cli:"f,file" usage:"path to a file containing brainfuck code"`
}

func main() {
	os.Exit(cli.Run(new(Args), func(ctx *cli.Context) error {
		args := ctx.Argv().(*Args)
		if args.File == "" {
			brainfuck.RunShell(args.Debug)
		} else {
			readAndRun(args.File, args.Debug, args.Interactive)
		}
		return nil
	}))
}

func readAndRun(path string, debug bool, interactive bool) {
	code, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}
	brainfuck.Run(string(code), debug, interactive)
}
