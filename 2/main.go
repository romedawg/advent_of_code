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

// Takes an input and returns a map object
func inputMap(input string)(map[int]string){

	inputMap := make(map[int]string, len(input))

	for n, uv := range input {
		inputMap[n] = string(uv)
	}
	return inputMap
}

// Takes a map object
func countUnique(inputMap map[int]string)(map[string]int){

	key := make([]int, len(inputMap))
	valueCount := make(map[string]int, len(inputMap))

	for index, v := range inputMap {
		key = append(key, index)
		valueCount[v] = 0
	}

	for k, v := range inputMap{
		for _, vv := range inputMap[k+1:-1] {
			if v == string(vv) {
				valueCount[v] = +1
			}
		}
	}
	return valueCount
}

func main(){

	dd := inputMap(inputText("input.txt"))
	for _, n := range dd{
		fmt.Printf("number of unique values: %s\n", string(n))
	}

	fmt.Printf("%s", countUnique(dd))
}