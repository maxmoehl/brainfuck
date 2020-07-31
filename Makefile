default: build run

build:
	go build -o bin/brainfuck cmd/brainfuck.go

run:
	./bin/brainfuck examples/helloworld.brainfuck

debug:
	./bin/brainfuck examples/helloworld.brainfuck -d