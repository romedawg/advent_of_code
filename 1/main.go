package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main(){

	var r io.Reader
	r, err := os.Open("input.txt")
	if err != nil{
		fmt.Printf("could not open file %s", err)
	}

	// Loop through this code
	data := make([]byte, 256)
	n, err := r.Read(data)
	if err !=nil{
		fmt.Printf( "could not read %s", err)
	}

	split_data := bytes.Split(data[:n], []byte(", "))
	fmt.Printf("input data is  is %s\n", split_data)
	currency_value := 0

	for _, d := range split_data {
		bb := string(d)
		cv, err := strconv.Atoi(strings.TrimSpace(bb))
		if err != nil{
			fmt.Printf( "cannot convert %s", err)
		}
		currency_value += cv

		//fmt.Printf("bb %s\n", bb)
		fmt.Printf("adding %d\n", cv)
		fmt.Printf("current value is %v\n", currency_value)
	}

	fmt.Printf("Currency value is %v\n", currency_value)
}
