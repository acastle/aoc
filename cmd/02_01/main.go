package main

import (
	"acastle.dev/aoc/pkg/intcode"
	"fmt"
)

func main() {
	program, err := intcode.ReadProgramFromFile("inputs/02_01/input")
	if err != nil {
	  panic(err)
  }

	program.Memory[1] = 12
	program.Memory[2] = 2
	err = program.Run()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Position 0: %d", program.Memory[0])
}