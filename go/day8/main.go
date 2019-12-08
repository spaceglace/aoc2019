package main

import (
	"fmt"
	"strings"
)

func verifyImage(input [][]string, width, height int) int {
	fewestZero := width*height + 1
	result := 0

	for _, layer := range input {
		// store a count of pixels seen
		pixels := map[string]int{
			"0": 0,
			"1": 0,
			"2": 0,
		}

		// iterate over layer and store pixel counts
		for _, pixel := range layer {
			pixels[pixel]++
		}

		// verification is based on layer with fewest zeroes
		if pixels["0"] < fewestZero {
			result = pixels["1"] * pixels["2"]
			fewestZero = pixels["0"]
		}
	}

	return result
}

func main() {
	input := readInput("./input/input.txt")
	width := 25
	height := 6

	output := [][]string{}
	for i := 0; i < len(input); i += (width * height) {
		output = append(output, strings.Split(input[i:i+width*height], ""))
	}

	fmt.Println("Verification Code:", verifyImage(output, width, height))
}
