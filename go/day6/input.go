package main

import (
	"io/ioutil"
	"strings"
)

func readInput(path string) string {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return string(file)
}

func parseInput(input string) [][]string {
	output := [][]string{}

	for _, line := range strings.Split(input, "\n") {
		output = append(output, strings.SplitN(line, ")", 2))
	}

	return output
}
