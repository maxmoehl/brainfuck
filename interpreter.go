package brainfuck

import (
	"fmt"
	"strconv"
)

var (
	memory         []int
	memoryPointer  int
	commands       []rune
	commandPointer int
	Debug          bool
)

func init() {
	memoryPointer = 0
	memory = make([]int, 1, 100)
	commandPointer = 0
}

func Run(input string) {
	commands = []rune(input)
	for ; commandPointer < len(commands); commandPointer++ {
		readRune(commands[commandPointer])
	}
}

func readRune(c rune) {
	if Debug {
		fmt.Printf("action: %v\n", string(c))
		fmt.Printf("memory: %v\n", memory)
	}
	switch c {
	case '.':
		print()
	case ',':
		read()
	case '+':
		increase()
	case '-':
		decrease()
	case '<':
		moveLeft()
	case '>':
		moveRight()
	case '[':
		startLoop()
	case ']':
		endLoop()
	}
}

func print() {
	if Debug {
		fmt.Printf("out: %v\n", string(rune(memory[memoryPointer])))
	} else {
		fmt.Printf("%v", string(rune(memory[memoryPointer])))
	}
}

func read() {
	// TODO implement
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
	} else if Debug {
		fmt.Printf("warning: tried to move pointer below 0")
	}
}

// Operation for '['
//
// If the current address is 0 the cursor is moved to the corresponding
// closing bracket otherwise the program moves to the next command. Panics
// if closing bracket cannot be found.
func startLoop() {
	if memory[memoryPointer] == 0 {
		oldAddress := commandPointer
		brackets := 0
		commandPointer++
		for ; commandPointer < len(commands); commandPointer++ {
			switch commands[commandPointer] {
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

func endLoop() {
	if memory[memoryPointer] != 0 {
		oldAddress := commandPointer
		brackets := 0
		commandPointer--
		for ; commandPointer >= 0; commandPointer-- {
			switch commands[commandPointer] {
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
