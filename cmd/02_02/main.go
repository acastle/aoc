package main

import (
	"acastle.dev/aoc/pkg/intcode"
	"fmt"
)

func main() {
	base, err := intcode.ReadProgramFromFile("inputs/02_01/input")
	if err != nil {
	  panic(err)
  }

  for noun := 0; noun < 100; noun++ {
  	for verb := 0; verb < 100; verb++ {
			program := intcode.Program{
				Memory:  append([]int(nil), base.Memory...),
				Counter: 0,
			}

			program.Memory[1] = noun
			program.Memory[2] = verb

			err = program.Run()
			if err != nil {
				panic(err)
			}

			if program.Memory[0] == 19690720 {
				fmt.Printf("Noun: %d, Verb: %d, (100 * noun + verb) = %d", noun, verb, 100 * noun + verb)
				return
			}
		}
	}
}