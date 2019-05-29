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

// Takes an input and returns a slice of strings
func inputMap(input string)([]string){

	inputMap := make([]string, len(input))

	for n, uv := range input {
		inputMap[n] = string(uv)
	}
	return inputMap
}

// Takes a slice of strings and returns a map[letter(string)]count(int)
func countUnique(inputSlice []string)(map[string]int){

	//key := make(map[int]string, len(inputSlice))
	letterCount := make(map[string]int, len(inputSlice))

	for _, value := range inputSlice {
		//key[index] = value
		letterCount[value] = 1
	}

	for k, l := range inputSlice {
		for _, vv := range inputSlice[k+1:] {
			if inputSlice[k] == string(vv) {
				letterCount[l] += 1
			}
		}
	}
	return letterCount
}

func main(){

	letters := inputMap(inputText("input.txt"))

	letterCount := countUnique(letters)
	fmt.Printf("letter %s\n", letterCount)

	for letter, count := range letterCount{
		fmt.Printf("letter %s count %d\n", letter, count)
	}

}