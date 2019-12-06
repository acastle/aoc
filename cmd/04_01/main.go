package main

import "fmt"

func main() {
	low := 265275
	high := 781584

	part1 := []int{}
	part2 := []int{}
	for i := low; i<=high; i++ {
		if ValidPass(i) {
			part1 = append(part1, i)
		}

		if ValidPass2(i) {
			part2 = append(part2, i)
		}
	}

	fmt.Printf("Part1 valid combinations: %d\n", len(part1))
	fmt.Printf("Part2 valid combinations: %d\n", len(part2))
}

func SplitCharacters(input int) []int {
	chars := []int{}
	for input > 0{
		char := input % 10
		input = (input - char)/10
		chars = append([]int{char}, chars...)
	}

	return chars
}

func Doubles(input []int) bool {
	for i,v := range input {
		if i > 0 && input[i-1] == v {
			return true
		}
	}

	return false
}

func ContainsStandaloneDouble(input []int) bool {
	for i,v := range input {
		left := i - 2 < 0 || input[i - 2] != v
		right := i + 1 >= len(input) || input[i + 1] != v
		double := i - 1 >=0 && input[i - 1] == v
		if left && right && double{
			return true
		}
	}

	return false
}

func Increasing(input []int) bool {
	for i,v := range input {
		if i > 0 && input[i-1] > v {
			return false
		}
	}

	return true
}

func ValidPass(input int) bool {
	split := SplitCharacters(input)
	return Doubles(split) && Increasing(split)
}

func ValidPass2(input int) bool {
	split := SplitCharacters(input)
	return ContainsStandaloneDouble(split) && Increasing(split)
}