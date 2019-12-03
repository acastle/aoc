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

	points := nav.IntersectPaths(paths[0], paths[1])
	lowest := nav.ManhattanDistance(points[0])
	point := points[0]
	for _, p := range points[1:] {
		dist := nav.ManhattanDistance(p)
		if dist < lowest {
			point = p
			lowest = dist
		}
	}

	fmt.Println(points)
	fmt.Println(point)
	fmt.Println(lowest)
}