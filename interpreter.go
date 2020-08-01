package brainfuck

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

var (
	memory        []int
	memoryPointer int
	commands      []byte
	cursor        int
	reader        *bufio.Reader
	debug         bool
)

const (
	dot      byte = '.'
	comma    byte = ','
	plus     byte = '+'
	minus    byte = '-'
	lt       byte = '<'
	gt       byte = '>'
	lBracket byte = '['
	rBracket byte = ']'
)

func init() {
	memoryPointer = 0
	memory = make([]int, 1, 100)
	reader = bufio.NewReader(os.Stdin)
}

func Run(input string, d bool, i bool) {
	debug = d
	run(stringToByteArray(input))
	if i {
		RunShell(d)
	}
}

func RunShell(d bool) {
	debug = d
	for {
		fmt.Print("- ")
		b, err := reader.ReadBytes('\n')
		if err != nil {
			panic(err.Error())
		}
		s := string(b)
		if s == "exit\n" {
			fmt.Print("Bye")
			os.Exit(0)
		}
		run(stringToByteArray(s))
		// If there was at least one print operation append a new line
		if bytes.Contains(b, []byte{comma}) {
			fmt.Print("\n")
		}

	}
}

// Executes a sequence of bytes one by one.
// Each call to run sets the commands slice to the
// input passed and resets the cursor to 0
func run(input []byte) {
	commands = input
	cursor = 0
	for ; cursor < len(commands); cursor++ {
		exec(commands[cursor])
	}
}

// Converts the string to a byte array and drops any rune
// that is not an instruction. The input needs to be a string
// because in a byte array we wouldn't be able to determine
// the bounds of a single character because Go uses UTF-8.
func stringToByteArray(s string) (commands []byte) {
	for _, r := range s {
		b := []byte(string(r))
		// Brainfuck commands only have a single byte
		if len(b) > 1 {
			continue
		}
		// Check if byte is part of out instruction set
		switch b[0] {
		case dot:
		case comma:
		case plus:
		case minus:
		case lt:
		case gt:
		case lBracket:
		case rBracket:
		default:
			continue
		}
		commands = append(commands, b[0])
	}
	return
}

// Executes a single instruction. If debug is true, the action to be executed,
// the memory and the memoryPointer will be printed.
func exec(b byte) {
	if debug {
		fmt.Printf("action: %v\n", string(b))
		fmt.Printf("memory: %v\n", memory)
		fmt.Printf("memoryPointer: %v\n", memoryPointer)
	}
	switch b {
	case dot:
		print()
	case comma:
		read()
	case plus:
		increase()
	case minus:
		decrease()
	case lt:
		moveLeft()
	case gt:
		moveRight()
	case lBracket:
		startLoop()
	case rBracket:
		endLoop()
	}
}

// Operation for '.'
//
// Prints a single character to the standard output. If debug is true
// the output will be prefixed with "out: " because the console is already
// filled with output from each step.
func print() {
	if debug {
		fmt.Printf("out: %v\n", string(rune(memory[memoryPointer])))
	} else {
		fmt.Printf("%v", string(rune(memory[memoryPointer])))
	}
}

// Operation for ','
//
// Reads a line form the standard input. If more than one byte
// was provided only the first byte will be recognized, if debug
// is true a warning is printed
func read() {
	b, err := reader.ReadBytes('\n')
	if err != nil {
		panic(err.Error())
	}
	// Remove last byte since it is the newline character
	b = b[:len(b)-1]
	if len(b) > 1 && debug {
		fmt.Print("warning: multiple bytes were submitted\n")
	}
	memory[memoryPointer] = int(b[0])
}

func increase() {
	memory[memoryPointer]++
}

func decrease() {
	memory[memoryPointer]--
}

func moveRight() {
	// If memory pointer would be out of bounds, append a new int
	if !(memoryPointer+1 < len(memory)) {
		memory = append(memory, 0)
	}
	memoryPointer++
}

func moveLeft() {
	if memoryPointer > 0 {
		memoryPointer--
	} else if debug {
		fmt.Print("warning: tried to move pointer below 0\n")
	}
}

// Operation for '['
//
// If the value of the current address is 0 the cursor is moved
// to the corresponding rBracket otherwise the program
// moves to the next command. Panics if closing bracket cannot
// be found.
func startLoop() {
	if memory[memoryPointer] == 0 {
		oldAddress := cursor
		brackets := 0
		cursor++
		for ; cursor < len(commands); cursor++ {
			switch commands[cursor] {
			case '[':
				brackets++
			case ']':
				if brackets == 0 {
					return
				} else {
					brackets--
				}
			}
		}
		panic("missing closing bracket matching to bracket on position " + strconv.Itoa(oldAddress))
	}
}

// Operation for ']'
//
// If the value of the current address is 0, this function does nothing
// and the pointer gets moved to the next position. If it is not zero
// the cursor is moved backwards to the corresponding opening lBracket.
// Panics is opening bracket cannot be found.
func endLoop() {
	if memory[memoryPointer] != 0 {
		oldAddress := cursor
		brackets := 0
		cursor--
		for ; cursor >= 0; cursor-- {
			switch commands[cursor] {
			case '[':
				if brackets == 0 {
					return
				} else {
					brackets--
				}
			case ']':
				brackets++
			}
		}
		panic("missing opening bracket matching to bracket on position " + strconv.Itoa(oldAddress))
	}
}
