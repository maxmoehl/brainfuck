package main

import (
	"fmt"
	"github.com/maxmoehl/brainfuck"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: brainfuck <path to brainfuck file> [-d | --debug]")
		os.Exit(1)
	}
	filePath := os.Args[1]
	code, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err.Error())
	}
	if len(os.Args) > 2 && (os.Args[2] == "--debug" || os.Args[2] == "-d") {
		brainfuck.Debug = true
	} else if len(os.Args) > 2 {
		fmt.Println("Usage: brainfuck <path to brainfuck file> [-d | --debug]")
		os.Exit(1)
	}
	brainfuck.Run(string(code))
}
