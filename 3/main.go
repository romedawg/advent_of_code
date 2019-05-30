package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os"
	"strings"
)

// Define the base of the carpet
//struct{
//default_text "."
//pattern = "#"
//width = 11
//heght = 11
//}

func input(inputFile string)[]string{
	file, err := os.Open(inputFile)
	if err != nil{
		errors.Wrap(err, "failed to open file")
	}

	read := bufio.NewReader(file)

	//lineData := make([]byte, 0, 256)
	var data []string

	defer file.Close()
	for {
		line, _, err := read.ReadLine()

		if err == io.EOF{
			break
		}
		data = append(data, string(line))
	}

	return data
}
func returnCarpet(order string){

	dd := strings.Split(string(order), " ")

	for _,v := range dd {
		fmt.Printf("%s", v)
	}
}

func main(){

	data := input("input.txt")

	for _, x := range data{
		dd := strings.Split(x, " ")

		fmt.Printf("%s\n", dd[0])
	}
}
