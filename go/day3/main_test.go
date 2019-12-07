package main

import (
	"reflect"
	"testing"
)

func TestReadInput(t *testing.T) {
	type goal struct {
		Input    string
		Expected string
	}
	goals := []goal{
		{
			Input:    "input/example1.txt",
			Expected: "R8,U5,L5,D3\nU7,R6,D4,L4",
		},
		{
			Input:    "input/example2.txt",
			Expected: "R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83",
		},
	}

	for _, g := range goals {
		actual := readInput(g.Input)
		if g.Expected != actual {
			t.Errorf(
				"Mismatch in %s. Expected:\n%s\nActual:\n%s",
				g.Input,
				g.Expected,
				actual,
			)
		}
	}
}

func TestParseInput(t *testing.T) {
	type goal struct {
		Input    string
		Expected [][]string
	}
	goals := []goal{
		{
			Input: "R8,U5,L5,D3\nU7,R6,D4,L4",
			Expected: [][]string{
				[]string{"R8", "U5", "L5", "D3"},
				[]string{"U7", "R6", "D4", "L4"},
			},
		},
		{
			Input: "R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83",
			Expected: [][]string{
				[]string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"},
				[]string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"},
			},
		},
	}

	for _, g := range goals {
		actual := parseInput(g.Input)

		if !reflect.DeepEqual(g.Expected, actual) {
			t.Errorf(
				"Mismatch for input:\n%s\nExpected:\n%v\nActual:\n%v",
				g.Input,
				g.Expected,
				actual,
			)
		}
	}
}

func TestParseStep(t *testing.T) {
	type goal struct {
		Input     string
		Direction string
		Distance  int
	}
	goals := []goal{
		{"R8", "R", 8},
		{"U5", "U", 5},
		{"L12", "L", 12},
		{"D49", "D", 49},
	}

	for _, g := range goals {
		actualDirection, actualDistance := parseStep(g.Input)

		if actualDirection != g.Direction || actualDistance != g.Distance {
			t.Errorf(
				"%s Expected direction %s and distance %d, but got %s and %d",
				g.Input,
				g.Direction,
				g.Distance,
				actualDirection,
				actualDistance,
			)
		}
	}
}

func TestAddStepToPath(t *testing.T) {
	type goal struct {
		Direction string
		Distance  int
		Path      []point
		Expected  []point
	}
	goals := []goal{
		{
			Direction: "R",
			Distance:  3,
			Path:      []point{point{0, 0}},
			Expected: []point{
				point{0, 0}, point{1, 0}, point{2, 0}, point{3, 0},
			},
		},
		{
			Direction: "L",
			Distance:  3,
			Path:      []point{point{0, 0}},
			Expected: []point{
				point{0, 0}, point{-1, 0}, point{-2, 0}, point{-3, 0},
			},
		},
		{
			Direction: "U",
			Distance:  3,
			Path:      []point{point{5, 5}},
			Expected: []point{
				point{5, 5}, point{5, 6}, point{5, 7}, point{5, 8},
			},
		},
		{
			Direction: "D",
			Distance:  3,
			Path:      []point{point{-80, 50}},
			Expected: []point{
				point{-80, 50}, point{-80, 49}, point{-80, 48}, point{-80, 47},
			},
		},
	}

	for _, g := range goals {
		path := append(g.Path[:0:0], g.Path...)
		path = addStepToPath(path, g.Direction, g.Distance)

		if !reflect.DeepEqual(g.Expected, path) {
			t.Errorf(
				"Mismatch\nPath: %v\nDirection: %s  Distance %d\nExpected: %v\nActual: %v",
				g.Path,
				g.Direction,
				g.Distance,
				g.Expected,
				path,
			)
		}
	}
}
