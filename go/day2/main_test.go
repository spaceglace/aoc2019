package main

import "testing"

func isIntSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func TestParseInput(t *testing.T) {
	type goal struct {
		Input    string
		Expected []int
	}
	goals := []goal{
		{"1,0,0,0,99", []int{1, 0, 0, 0, 99}},
		{"2,3,0,3,99", []int{2, 3, 0, 3, 99}},
	}

	for _, g := range goals {
		actual := parseInput(g.Input)
		if !isIntSliceEqual(g.Expected, actual) {
			t.Errorf("Input %s expected %v but got %v", g.Input, g.Expected, actual)
		}
	}
}

func TestAdd(t *testing.T) {
	type goal struct {
		Input    []int
		Offset   int
		Expected []int
	}
	goals := []goal{
		{[]int{1, 0, 0, 0, 99}, 0, []int{2, 0, 0, 0, 99}},
	}

	for _, g := range goals {
		actual := append(g.Input[:0:0], g.Input...)
		add(actual, g.Offset)
		if !isIntSliceEqual(g.Expected, actual) {
			t.Errorf(
				"Input %v with offset %d expected %v but got %v",
				g.Input,
				g.Offset,
				g.Expected,
				actual,
			)
		}
	}
}

func TestMultiply(t *testing.T) {
	type goal struct {
		Input    []int
		Offset   int
		Expected []int
	}
	goals := []goal{
		{[]int{2, 3, 0, 3, 99}, 0, []int{2, 3, 0, 6, 99}},
		{[]int{2, 4, 4, 5, 99, 0}, 0, []int{2, 4, 4, 5, 99, 9801}},
	}

	for _, g := range goals {
		actual := append(g.Input[:0:0], g.Input...)
		multiply(actual, g.Offset)
		if !isIntSliceEqual(g.Expected, actual) {
			t.Errorf(
				"Input %v with offset %d expected %v but got %v",
				g.Input,
				g.Offset,
				g.Expected,
				actual,
			)
		}
	}
}

func TestRunProgram(t *testing.T) {
	type goal struct {
		Input    []int
		Expected []int
	}
	goals := []goal{
		{[]int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}},
		{[]int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}},
		{[]int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}},
		{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}

	for _, g := range goals {
		actual := append(g.Input[:0:0], g.Input...)
		runProgram(actual)
		if !isIntSliceEqual(g.Expected, actual) {
			t.Errorf(
				"Input %v expected %v but got %v",
				g.Input,
				g.Expected,
				actual,
			)
		}
	}
}
