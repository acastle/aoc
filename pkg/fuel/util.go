package fuel

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func ReadModulesFromFile(path string) ([]Module, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	output := []Module{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mass, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, fmt.Errorf("read floats from file: %w", err)
		}

		output = append(output, Module{mass})
	}

	return output, nil
}