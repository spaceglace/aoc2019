package main

import (
	"fmt"
	"strconv"
)

func isSixDigits(input string) bool {
	return len(input) == 6
}

func doesNotDecrease(input string) bool {
	for i := 0; i < len(input)-1; i++ {
		if input[i+1] < input[i] {
			return false
		}
	}
	return true
}

func containsDouble(input string) bool {
	for i := 0; i < len(input)-1; i++ {
		if input[i+1] == input[i] {
			return true
		}
	}
	return false
}

func main() {
	count := 0
	for i := 152085; i <= 670283; i++ {
		str := strconv.Itoa(i)
		if isSixDigits(str) && doesNotDecrease(str) && containsDouble(str) {
			count++
		}
	}
	fmt.Println(count)
}
