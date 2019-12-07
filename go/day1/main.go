package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func readInput(path string) []int {
	output := []int{}

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		output = append(output, val)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return output
}

func getFuelFromMass(mass int) int {
	// Integer division already truncates in go
	return (mass / 3) - 2
}

func getRecursiveFuel(mass int) int {
	extra := getFuelFromMass(mass)

	if extra <= 0 {
		return 0
	}

	return extra + getRecursiveFuel(extra)
}

func main() {
	totalFuel := 0
	for _, mass := range readInput("input/input.txt") {
		initial := getFuelFromMass(mass)
		added := getRecursiveFuel(initial)
		totalFuel += initial + added
	}

	log.Printf("Total fuel: %d", totalFuel)
}
