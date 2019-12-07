package main

import (
	"reflect"
	"testing"
)

func TestLoadImmediateParam(t *testing.T) {
	type goal struct {
		Program  []int
		Offset   int
		Param    int
		Expected int
	}
	goals := []goal{
		goal{[]int{1002, 4, 3, 4, 33}, 0, 2, 3},
	}

	for _, g := range goals {
		actual := loadImmediateParam(g.Program, g.Offset, g.Param)
		if actual != g.Expected {
			t.Errorf("%+v received %d", g, actual)
		}
	}
}

func TestLoadPositionParam(t *testing.T) {
	type goal struct {
		Program  []int
		Offset   int
		Param    int
		Expected int
	}
	goals := []goal{
		goal{[]int{1002, 4, 3, 4, 33}, 0, 1, 33},
		goal{[]int{1002, 4, 3, 4, 33}, 0, 3, 33},
	}

	for _, g := range goals {
		actual := loadPositionParam(g.Program, g.Offset, g.Param)
		if actual != g.Expected {
			t.Errorf("%+v received %d", g, actual)
		}
	}
}

func TestLoadParam(t *testing.T) {
	type goal struct {
		Program  []int
		Offset   int
		Param    int
		Mode     mode
		Panic    bool
		Expected int
	}
	goals := []goal{
		goal{[]int{1002, 4, 3, 4, 33}, 0, 1, position, false, 33},
		goal{[]int{1002, 4, 3, 4, 33}, 0, 2, immediate, false, 3},
		goal{[]int{1002, 4, 3, 4, 33}, 0, 3, position, false, 33},
		goal{[]int{1002, 4, 3, 4, 33}, 0, 3, 8, true, 0},
	}

	for _, g := range goals {
		if g.Panic {
			func() {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("%+v did not panic, but should have", g)
					}
				}()
				loadParam(g.Program, g.Offset, g.Param, g.Mode)
			}()
		} else {
			actual := loadParam(g.Program, g.Offset, g.Param, g.Mode)
			if actual != g.Expected {
				t.Errorf("%+v received %d", g, actual)
			}
		}
	}
}

func TestWriteToProgram(t *testing.T) {
	type goal struct {
		Program  []int
		Offset   int
		Param    int
		Value    int
		Expected []int
	}
	goals := []goal{
		goal{[]int{1002, 4, 3, 4, 33}, 0, 1, 99, []int{1002, 4, 3, 4, 99}},
	}

	for _, g := range goals {
		program := append(g.Program[:0:0], g.Program...)
		writeToProgram(program, g.Offset, g.Param, g.Value)
		if !reflect.DeepEqual(program, g.Expected) {
			t.Errorf("%+v received %v", g, program)
		}
	}
}

func TestOpAdd(t *testing.T) {
	type goal struct {
		Program  []int
		Offset   int
		Param1   mode
		Param2   mode
		Expected []int
	}
	goals := []goal{
		goal{[]int{1101, 100, -1, 4, 0}, 0, immediate, immediate, []int{1101, 100, -1, 4, 99}},
	}

	for _, g := range goals {
		program := append(g.Program[:0:0], g.Program...)
		opAdd(program, g.Offset, g.Param1, g.Param2)

		if !reflect.DeepEqual(program, g.Expected) {
			t.Errorf("%+v received %v", g, program)
		}
	}
}

func TestOpMultiply(t *testing.T) {
	type goal struct {
		Program  []int
		Offset   int
		Param1   mode
		Param2   mode
		Expected []int
	}
	goals := []goal{
		goal{[]int{1102, 4, 3, 4, 33}, 0, position, immediate, []int{1102, 4, 3, 4, 99}},
	}

	for _, g := range goals {
		program := append(g.Program[:0:0], g.Program...)
		opMultiply(program, g.Offset, g.Param1, g.Param2)

		if !reflect.DeepEqual(program, g.Expected) {
			t.Errorf("%+v received %v", g, program)
		}
	}
}
