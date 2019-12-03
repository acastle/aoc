package main

import (
	"acastle.dev/aoc/pkg/fuel"
	"fmt"
)

func main() {
	modules, err := fuel.ReadModulesFromFile("inputs/01_01/input")

	if err != nil {
	  panic(err)
  }

	fmt.Printf("Total Fuel: %f", TotalFuel(modules))
}

func TotalFuel(module []fuel.Module) float64 {
	var sum float64
	for _, module := range module {
		sum += module.Fuel()
	}

	return sum
}