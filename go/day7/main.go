package main

import "fmt"

func runAmplifiers(program, phases []int) int {
	result1 := run(program, []int{phases[0], 0})
	result2 := run(program, []int{phases[1], result1})
	result3 := run(program, []int{phases[2], result2})
	result4 := run(program, []int{phases[3], result3})
	result5 := run(program, []int{phases[4], result4})
	return result5
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

	for _, phases := range generatePhasePermutations([]int{0, 1, 2, 3, 4}) {
		current := runAmplifiers(program, phases)
		if current > maxThruster {
			maxThruster = current
			maxThrusterConfig = phases
		}
	}

	fmt.Printf("Max Thruster Output: %d\n", maxThrusterConfig)
	fmt.Printf("Max Phase Layout: %v\n", maxThruster)
}
