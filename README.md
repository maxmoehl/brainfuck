# Brainfuck interpreter

## What is brainfuck?

[Wikipedia](https://en.wikipedia.org/wiki/Brainfuck) says:  
> Brainfuck is an esoteric programming language created in 1993
> by Urban Müller.  
> Notable for its extreme minimalism, the language consists of
> only eight simple commands and an instruction pointer. While
> it is fully Turing complete, it is not intended for practical
> use, but to challenge and amuse programmers. Brainfuck simply
> requires one to break commands into microscopic steps.

### Instructions (again thanks to [Wikipedia](https://en.wikipedia.org/wiki/Brainfuck))

| Instruction | Meaning                                                                                                                                                                           |
|-------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| >           | increment the data pointer (to point to the next cell to the right).                                                                                                              |
| <           | decrement the data pointer (to point to the next cell to the left).                                                                                                               |
| +           | increment (increase by one) the byte at the data pointer.                                                                                                                         |
| -           | decrement (decrease by one) the byte at the data pointer.                                                                                                                         |
| .           | output the byte at the data pointer.                                                                                                                                              |
| ,           | accept one byte of input, storing its value in the byte at the data pointer.                                                                                                      |
| [           | if the byte at the data pointer is zero, then instead of moving the instruction pointer forward to the next command, jump it forward to the command after the matching ] command. |
| ]           | if the byte at the data pointer is nonzero, then instead of moving the instruction pointer forward to the next command, jump it back to the command after the matching [ command. |

## What the heck is this?

This is a brainfuck interpreter written in Go (aka GoLang), the only thing missing is the read function.

The main part is in `interpreter.go`, `cmd/brainfuck.go` parses
the command line arguments and starts execution.

## TODO

- [ ] Implement read function
- [ ] Add documentation
- [ ] Add unit tests