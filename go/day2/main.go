package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readInput(path string) string {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return string(file)
}

func parseInput(input string) []int {
	output := []int{}

	for _, str := range strings.Split(input, ",") {
		val, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		output = append(output, val)
	}

	return output
}

func add(input []int, offset int) {
	val := input[input[offset+1]] + input[input[offset+2]]
	input[input[offset+3]] = val
}

func multiply(input []int, offset int) {
	val := input[input[offset+1]] * input[input[offset+2]]
	input[input[offset+3]] = val
}

func runProgram(input []int) {
	for offset := 0; offset < len(input); offset += 4 {
		switch input[offset] {
		case 1:
			add(input, offset)
		case 2:
			multiply(input, offset)
		case 99:
			return
		default:
			fmt.Printf("Unknown opcode: %d (%v with offset %d)\n", input[offset], input, offset)
		}
	}
}

func main() {
	input := readInput("input/input.txt")
	values := parseInput(input)

	goal := 19690720

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			working := append(values[:0:0], values...)
			working[1] = noun
			working[2] = verb

			runProgram(working)

			if working[0] == goal {
				fmt.Printf("Answer:\n noun %d\n verb %d", noun, verb)
			}
		}
	}
}
