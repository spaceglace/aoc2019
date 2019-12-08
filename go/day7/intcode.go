package main

import (
	"fmt"
	"strconv"
)

type mode int

const (
	position mode = iota
	immediate
)

func loadImmediateParam(program []int, offset, param int) int {
	return program[offset+param]
}

func loadPositionParam(program []int, offset, param int) int {
	return program[program[offset+param]]
}

func loadParam(program []int, offset, param int, m mode) int {
	switch m {
	case position:
		return loadPositionParam(program, offset, param)
	case immediate:
		return loadImmediateParam(program, offset, param)
	default:
		panic(fmt.Errorf("Unknown parameter mode: %s", string(m)))
	}
}

func writeToProgram(program []int, offset, param, value int) {
	program[program[offset+param]] = value
}

func opAdd(program []int, offset int, param1mode, param2mode mode) int {
	param1 := loadParam(program, offset, 1, param1mode)
	param2 := loadParam(program, offset, 2, param2mode)
	result := param1 + param2
	writeToProgram(program, offset, 3, result)
	return 4
}

func opMultiply(program []int, offset int, param1mode, param2mode mode) int {
	param1 := loadParam(program, offset, 1, param1mode)
	param2 := loadParam(program, offset, 2, param2mode)
	result := param1 * param2
	writeToProgram(program, offset, 3, result)
	return 4
}

func opInput(program []int, offset, value int, param1mode mode) int {
	writeToProgram(program, offset, 1, value)
	return 2
}

func opOutput(program []int, offset int, param1mode mode) int {
	param1 := loadParam(program, offset, 1, param1mode)
	return param1
}

func opJumpIfTrue(program []int, offset int, param1mode, param2mode mode) int {
	param1 := loadParam(program, offset, 1, param1mode)
	param2 := loadParam(program, offset, 2, param2mode)

	if param1 != 0 {
		return param2 - offset
	}
	return 3
}

func opJumpIfFalse(program []int, offset int, param1mode, param2mode mode) int {
	param1 := loadParam(program, offset, 1, param1mode)
	param2 := loadParam(program, offset, 2, param2mode)

	if param1 == 0 {
		return param2 - offset
	}
	return 3
}

func opLessThan(program []int, offset int, param1mode, param2mode mode) int {
	param1 := loadParam(program, offset, 1, param1mode)
	param2 := loadParam(program, offset, 2, param2mode)

	if param1 < param2 {
		writeToProgram(program, offset, 3, 1)
	} else {
		writeToProgram(program, offset, 3, 0)
	}
	return 4
}

func opEquals(program []int, offset int, param1mode, param2mode mode) int {
	param1 := loadParam(program, offset, 1, param1mode)
	param2 := loadParam(program, offset, 2, param2mode)

	if param1 == param2 {
		writeToProgram(program, offset, 3, 1)
	} else {
		writeToProgram(program, offset, 3, 0)
	}
	return 4
}

func runOpcode(program, inputs []int, offset int, inputOffset *int) (int, bool) {
	opcode := "0000000000" + strconv.Itoa(program[offset])

	params := []mode{}
	for i := len(opcode) - 3; i >= 0; i-- {
		switch string(opcode[i]) {
		case "0": // position mode
			params = append(params, position)
		case "1": // immediate mode
			params = append(params, immediate)
		default:
			panic(fmt.Errorf("Unknown parameter mode: %s", string(opcode[i])))
		}
	}

	switch opcode[len(opcode)-2:] {
	case "01": // addition
		return opAdd(program, offset, params[0], params[1]), false
	case "02": // multiplication
		return opMultiply(program, offset, params[0], params[1]), false
	case "03": // input
		val := opInput(program, offset, inputs[*inputOffset], params[0])
		*inputOffset++
		return val, false
	case "04": // output
		return opOutput(program, offset, params[0]), true
	case "05": // jump-if-true
		return opJumpIfTrue(program, offset, params[0], params[1]), false
	case "06": // jump-if-false
		return opJumpIfFalse(program, offset, params[0], params[1]), false
	case "07": // less than
		return opLessThan(program, offset, params[0], params[1]), false
	case "08": // equals
		return opEquals(program, offset, params[0], params[1]), false
	case "99": // halt
		fmt.Println("\nProgram Halted.")
		return 0, false
	default:
		panic(fmt.Errorf("Unknown opcode %s", opcode))
	}
}

func run(program, inputs []int) int {
	inputOffset := 0
	for offset := 0; offset < len(program); {
		space, output := runOpcode(program, inputs, offset, &inputOffset)

		if output {
			return space
		}

		if space == 0 {
			break
		}

		offset += space
	}
	fmt.Println("Program Finished.")
	return 0
}
