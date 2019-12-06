package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Node struct {
	Name string
	Orbiting *Node
}

func ReadOrbits(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	paths := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		paths = append(paths, scanner.Text())
	}

	return paths, nil
}

func ParseOrbit(orbit string) (string, string) {
	names := strings.Split(orbit, ")")
	return names[0], names[1]
}

func Orbits(orbits []string) map[string]*Node {
	index := map[string]*Node{}
	index["COM"] = &Node {
		Name: "COM",
	}

	for _, o := range orbits {
		parentName, childName := ParseOrbit(o)
		parent, exists := index[parentName]
		if !exists {
			parent = &Node{ Name: parentName }
		}

		child, exists := index[childName]
		if !exists {
			child = &Node{
				Name: childName,
			}
		}

		child.Orbiting = parent
		index[parent.Name] = parent
		index[child.Name] = child
	}

	return index
}

func Depth(n *Node) int {
	if n.Orbiting == nil {
		return 0
	}

	return 1 + Depth(n.Orbiting)
}

func CommonAncestor(n1 *Node, n2 *Node) *Node {
	ancestors := map[string]*Node{}
	for n1.Orbiting != nil {
		ancestors[n1.Orbiting.Name] = n1.Orbiting
		n1 = n1.Orbiting
	}

	for n2.Orbiting != nil {
		_, exists := ancestors[n2.Orbiting.Name]
		if exists {
			return n2.Orbiting
		}

		n2 = n2.Orbiting
	}

	return nil
}

func DistanceToNode(n1 *Node, n2 *Node) int{
	count := 0
	for n1.Orbiting != nil {
		if n1.Orbiting.Name == n2.Name {
			break
		}

		n1 = n1.Orbiting
		count++
	}

	return count
}

func main() {
	o, err := ReadOrbits("inputs/06_01/input")
	if err != nil {
		panic(err)
	}

	orbits := Orbits(o)
	total := 0
	for _,v := range orbits {
		total += Depth(v)
	}

	fmt.Println("Problem 1: ", total)

	you := orbits["YOU"]
	san := orbits["SAN"]
	common := CommonAncestor(you, san)
	sanToCommon := DistanceToNode(san, common)
	youToCommon := DistanceToNode(you, common)
	distance := sanToCommon + youToCommon

	fmt.Println("Problem 2: ", distance)
}