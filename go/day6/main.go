package main

import "fmt"

type Body struct {
	Name   string
	Orbits *Body
}

func findBodyIndex(bodies []Body, name string) int {
	for i, body := range bodies {
		if body.Name == name {
			return i
		}
	}
	return -1
}

func generateBodies(orbits [][]string) []Body {
	output := []Body{}

	// get unique names
	names := map[string]bool{}
	for _, orbit := range orbits {
		names[orbit[0]] = true
		names[orbit[1]] = true
	}

	for name := range names {
		output = append(output, Body{
			Name: name,
		})
	}

	for _, orbit := range orbits {
		left := findBodyIndex(output, orbit[0])
		right := findBodyIndex(output, orbit[1])

		output[right].Orbits = &output[left]
	}

	return output
}

func main() {
	input := readInput("input/input.txt")
	values := parseInput(input)

	bodies := generateBodies(values)
	count := 0

	for _, body := range bodies {
		parent := body.Orbits

		for parent != nil {
			count++
			parent = parent.Orbits
		}
	}

	fmt.Println(count)
}
