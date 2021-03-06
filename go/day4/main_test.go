package main

import "testing"

func TestIsSixDigits(t *testing.T) {
	type goal struct {
		Input    string
		Expected bool
	}
	goals := []goal{
		goal{"111111", true},
		goal{"223450", true},
		goal{"123789", true},
		goal{"", false},
		goal{"12", false},
		goal{"1234567890", false},
	}

	for _, g := range goals {
		actual := isSixDigits(g.Input)
		if actual != g.Expected {
			t.Errorf(
				"%s expected %t but got %t",
				g.Input,
				g.Expected,
				actual,
			)
		}
	}
}

func TestDoesNotDecrease(t *testing.T) {
	type goal struct {
		Input    string
		Expected bool
	}
	goals := []goal{
		goal{"111111", true},
		goal{"223450", false},
		goal{"123789", true},
	}

	for _, g := range goals {
		actual := doesNotDecrease(g.Input)
		if actual != g.Expected {
			t.Errorf(
				"%s expected %t but got %t",
				g.Input,
				g.Expected,
				actual,
			)
		}
	}
}

func TestContainsIsolatedDouble(t *testing.T) {
	type goal struct {
		Input    string
		Expected bool
	}
	goals := []goal{
		goal{"111111", false},
		goal{"223450", true},
		goal{"123789", false},
		goal{"112233", true},
		goal{"123444", false},
		goal{"111122", true},
	}

	for _, g := range goals {
		actual := containsIsolatedDouble(g.Input)
		if actual != g.Expected {
			t.Errorf(
				"%s expected %t but got %t",
				g.Input,
				g.Expected,
				actual,
			)
		}
	}
}
