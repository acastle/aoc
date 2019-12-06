package main

import (
	"acastle.dev/aoc/pkg/intcode"
	"fmt"
)

func main() {
	program, err := intcode.ReadProgramFromFile("inputs/05_01/input")
	if err != nil {
		panic(err)
	}

	var diag int
	program.In = func()int{ return 1 }
	program.Out = func(a int) {
		if program.Counter + 2 < len(program.Memory) && program.Memory[program.Counter + 2] == 99 {
			diag = a
		}
	}

	err = program.Run()
	if err != nil {
		panic(err)
	}

	fmt.Printf("(Problem 1) Diagnostic Code: %d\n", diag)

	program, err = intcode.ReadProgramFromFile("inputs/05_01/input")
	if err != nil {
		panic(err)
	}

	program.In = func()int{ return 5 }
	program.Out = func(a int) {
		if program.Counter + 2 < len(program.Memory) && program.Memory[program.Counter + 2] == 99 {
			diag = a
		}
	}

	err = program.Run()
	if err != nil {
		panic(err)
	}

	fmt.Printf("(Problem 2) Diagnostic Code: %d\n", diag)
}