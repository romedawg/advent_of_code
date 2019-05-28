package main

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
)

func inputText(input string)(string){

	file, err := os.Open(input)
	if err != nil {
		errors.Wrapf(err, "could not open file")
	}

	inputData := make([]byte, 10)
	dataCount, err := file.Read(inputData)

	return string(inputData[:dataCount])
}

// Takes an input and returns a map object - no longer a map, TODO move this to the input string func
func inputMap(input string)([]string){

	inputMap := make([]string, len(input))

	for n, uv := range input {
		inputMap[n] = string(uv)
	}
	return inputMap
}

// Takes a map object
func countUnique(inputSlice []string)(map[string]int){

	key := make([]int, len(inputSlice))
	letterCount := make(map[string]int, len(inputSlice))

	for index, value := range inputSlice {
		key = append(key, index)
		letterCount[value] = 1
	}

	for k, l := range inputSlice {
		for _, vv := range inputSlice[k+1:] {
			if inputSlice[k] == string(vv) {
				letterCount[l] = letterCount[l] + 1
			}
		}
	}
	return letterCount
}

func main(){

	// Contains a slice of strings, whatever the input is.
	dd := inputMap(inputText("input.txt"))

	valueCount := countUnique(dd)
	fmt.Printf("letter %s\n", valueCount)
}