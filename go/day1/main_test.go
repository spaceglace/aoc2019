package main

import "testing"

type goal struct {
	Input  int
	Output int
}

func TestGetFuelFromMass(t *testing.T) {
	goals := []goal{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}

	for _, g := range goals {
		actual := getFuelFromMass(g.Input)
		if actual != g.Output {
			t.Errorf(
				"Input was %d, expected %d but got %d",
				g.Input,
				g.Output,
				actual,
			)
		}
	}
}

func TestGetRecursiveFuel(t *testing.T) {
	goals := []goal{
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}

	for _, g := range goals {
		actual := getRecursiveFuel(g.Input)
		if actual != g.Output {
			t.Errorf(
				"Input was %d, expected %d but got %d",
				g.Input,
				g.Output,
				actual,
			)
		}
	}
}
