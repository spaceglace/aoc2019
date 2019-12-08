package main

import (
	"io/ioutil"
)

func readInput(path string) string {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return string(file)
}
