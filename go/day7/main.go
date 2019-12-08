package main

import (
	"fmt"
	"sync"
)

func runAmplifiers(program, phases []int) int {
	var wg sync.WaitGroup
	wg.Add(5)

	// form the communication links
	fmt.Printf("Starting phase sequence %v:\t", phases)
	AtoB := make(chan int)
	BtoC := make(chan int)
	CtoD := make(chan int)
	DtoE := make(chan int)
	EtoA := make(chan int, 2)

	// Clone the program so each amp gets their own code
	go run(append(program[:0:0], program...), EtoA, AtoB, "A", &wg)
	go run(append(program[:0:0], program...), AtoB, BtoC, "B", &wg)
	go run(append(program[:0:0], program...), BtoC, CtoD, "C", &wg)
	go run(append(program[:0:0], program...), CtoD, DtoE, "D", &wg)
	go run(append(program[:0:0], program...), DtoE, EtoA, "E", &wg)

	// seed each machine with their phase
	EtoA <- phases[0]
	AtoB <- phases[1]
	BtoC <- phases[2]
	CtoD <- phases[3]
	DtoE <- phases[4]

	// seed machine A with the initial value
	EtoA <- 0

	// wait for all the machines to halt
	wg.Wait()

	result := <-EtoA
	fmt.Println("...completed: ", result)
	return result
}

// I'll fully admit I don't know what's going on here
// https://yourbasic.org/golang/generate-permutation-slice-string/
func generatePhasePermutations(a []int) [][]int {
	output := [][]int{}
	var p func(a []int, i int)
	p = func(a []int, i int) {
		if i > len(a) {
			clone := append(a[:0:0], a...)
			output = append(output, clone)
			return
		}
		p(a, i+1)
		for j := i + 1; j < len(a); j++ {
			a[i], a[j] = a[j], a[i]
			p(a, i+1)
			a[i], a[j] = a[j], a[i]
		}
	}
	p(a, 0)
	return output
}

func main() {
	input := readInput("input/input.txt")
	program := parseInput(input)

	maxThruster := 0
	var maxThrusterConfig []int

	for _, phases := range generatePhasePermutations([]int{5, 6, 7, 8, 9}) {
		current := runAmplifiers(program, phases)
		if current > maxThruster {
			maxThruster = current
			maxThrusterConfig = phases
		}
	}

	fmt.Printf("Max Thruster Output: %d\n", maxThrusterConfig)
	fmt.Printf("Max Phase Layout: %v\n", maxThruster)
}
