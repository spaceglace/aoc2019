package main

import (
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
