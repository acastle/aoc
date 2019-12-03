package main

import (
	"acastle.dev/aoc/pkg/nav"
	"fmt"
)

func main() {
	paths,err := nav.ParsePathsFromFile("inputs/03_01/input")
	if err != nil {
		panic(err)
	}

	a := paths[0]
	b := paths[1]
	lowest := nav.MinStepsToIntersection(a, b)

	fmt.Println(lowest)
}