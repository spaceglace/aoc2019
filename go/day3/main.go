package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"sync"
)

type point struct {
	X int
	Y int
}

type collision struct {
	X        int
	Y        int
	Distance int
	Delay    int
}

func intAbs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func (p point) distance() int {
	return intAbs(p.X) + intAbs(p.Y)
}

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
		output = append(output, strings.Split(line, ","))
	}

	return output
}

func parseStep(input string) (string, int) {
	direction := input[0:1]
	distance, err := strconv.Atoi(input[1:])
	if err != nil {
		panic(err)
	}

	return direction, distance
}

func addStepToPath(path []point, direction string, distance int) []point {
	start := path[len(path)-1]

	modifier := map[string]point{
		"U": point{0, 1},
		"D": point{0, -1},
		"L": point{-1, 0},
		"R": point{1, 0},
	}

	for i := 1; i <= distance; i++ {
		path = append(path, point{
			X: start.X + (i * modifier[direction].X),
			Y: start.Y + (i * modifier[direction].Y),
		})
	}

	return path
}

func generatePath(input []string) []point {
	output := []point{point{0, 0}}

	for _, step := range input {
		direction, distance := parseStep(step)
		output = addStepToPath(output, direction, distance)
	}

	return output
}

func findCollisions(first, second []point) []collision {
	output := []collision{}

	for fi, fv := range first {
		for si, sv := range second {
			if fv.X == sv.X && fv.Y == sv.Y && (fv.X != 0 || fv.Y != 0) {
				output = append(output, collision{
					X:        fv.X,
					Y:        fv.Y,
					Distance: fv.distance(),
					Delay:    fi + si,
				})
				break
			}
		}
	}

	return output
}

func main() {
	input := readInput("input/input.txt")
	lines := parseInput(input)
	paths := [][]point{}

	buffer := make(chan []point, len(lines))
	var wg sync.WaitGroup

	for _, line := range lines {
		go func(line []string) {
			buffer <- generatePath(line)
			wg.Done()
		}(line)
		wg.Add(1)
	}

	wg.Wait()
	close(buffer)
	for path := range buffer {
		paths = append(paths, path)
	}

	for _, c := range findCollisions(paths[0], paths[1]) {
		fmt.Printf("%+v\n", c)
	}
}
