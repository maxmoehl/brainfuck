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
	if os.Args[2] == "--debug" || os.Args[2] == "-d" {
		brainfuck.Debug = true
	} else {
		fmt.Println("Usage: brainfuck <path to brainfuck file> [-d | --debug]")
	}
	brainfuck.Run(string(code))
}