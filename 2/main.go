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

	inputData := make([]byte, 6)
	dataCount, err := file.Read(inputData)

	return string(inputData[:dataCount])
}

// Takes an input and returns a map object
func inputMap(input string)(map[int]string){

	inputMap := make(map[int]string, len(input))

	for n, uv := range input {
		inputMap[n] = string(uv)
	}
	return inputMap
}

// Takes a map object
func countUnique(inputMap map[int]string){

	key := make([]int, len(inputMap))

	for index := range inputMap {
		key = append(key, index)
	}

	//var uniqueValues []string
	//var duplicateValues map[string]int
	//
	//for k, v := range inputMap{
	//	if v ==
	//}


	// Now match input value w/ all vaules.  do the keys matchup. if not unique pile
	// else
}

func main(){
	// evaluate...
	// how many unique intances ? - Done
	// How to match up duplicates
	// Do a count of duplicates.

	dd := inputMap(inputText("input.txt"))
	for _, n := range dd{
		fmt.Printf("number of unique values: %s\n", string(n))
	}
}