package main

import (
	"bytes"
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

	numberOfSlices := bytes.IndexByte(inputData, 0)
	fmt.Printf("number of slices: %d\n", numberOfSlices)

	return string(inputData[:dataCount])
}

func inputMap(input string)map[int]string{

	inputMap := make(map[int]string, len(input))

	for _, i := range input{
		fmt.Printf("i: %s\n", string(i))
		for n, uv := range input {
			inputMap[n] = string(uv)
		}
	}
	return inputMap
}

func countUnique(inputMap map[int]string){

	// return unique count, doubles count
}

func main(){
	// evaluate...
	// how many unique intances ? - Done
	// How to match up duplicates
	// Do a count of duplicates.
	data := inputText("input.txt")

	dd := inputMap(data)
	for _, n := range dd{
		fmt.Printf("number of unique values: %s\n", string(n))
	}
}