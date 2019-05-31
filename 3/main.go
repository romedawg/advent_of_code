package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os"
	"strconv"
)

// Define the base of the carpet
type orderForm struct {
	orderNumber int
	pattern string
	border string
	portrait string
}

func input(inputFile string)[]string{
	file, err := os.Open(inputFile)
	if err != nil{
		errors.Wrap(err, "failed to open file")
	}

	read := bufio.NewReader(file)

	//lineData := make([]byte, 0, 256)
	var rawData []string

	defer file.Close()
	for {
		line, _, err := read.ReadLine()

		if err == io.EOF{
			break
		}
		rawData = append(rawData, string(line))
	}
	return rawData
}

func main(){

	data := input("input.txt")

	order := make(map[int]orderForm)

	for n, x := range data{
		fmt.Printf("%s\n", string(x[1:2]))

		orderNum, err :=  strconv.Atoi(string(x[1:2]))
		if err !=nil {
			errors.Wrap(err, "failed")
		}
		pattern := string(x[3:4])
		border :=  string(x[5:8])
		portrait := string(x[10:])

		order[n] = orderForm{
			 orderNum,
			 pattern,
			 border,
			 portrait,
		}
	}
	fmt.Printf("%d\n",order[0])
}
