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

func findTotalOrbit(bodies []Body) int {
	count := 0

	for _, body := range bodies {
		parent := body.Orbits

		for parent != nil {
			count++
			parent = parent.Orbits
		}
	}

	return count
}

func findTransferOrbits(bodies []Body) {
	santa := findBodyIndex(bodies, "SAN")
	you := findBodyIndex(bodies, "YOU")

	santaParent := &bodies[santa]
	youParent := &bodies[you]

	// track how far both paths have travelled
	santaCount := -1
	youCount := -1

	// loop until we find a common ancestor
	for youParent.Name != santaParent.Name {
		// move up to the next planet on santa's chain
		santaParent = santaParent.Orbits
		santaCount++

		// reset our position
		youCount = -1
		youParent = &bodies[you]

		// iterate through all our orbits to COM
		for youParent.Orbits != nil {
			youParent = youParent.Orbits
			youCount++
			// check if we found a common ancestor
			if youParent.Name == santaParent.Name {
				break
			}
		}
	}

	fmt.Printf("Santa's distance to common ancestor: %d\n", santaCount)
	fmt.Printf("Your distance to common ancestor: %d\n", youCount)
	fmt.Printf("Combined transfer orbits: %d\n", santaCount+youCount)
}

func main() {
	input := readInput("input/input.txt")
	values := parseInput(input)
	bodies := generateBodies(values)

	findTransferOrbits(bodies)
}
