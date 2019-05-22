package main

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os"
	"strconv"
)

func main(){

	var r io.Reader
	r, err := os.Open("input.txt")
	if err != nil{
		fmt.Printf("could not open file %s", err)
	}

	data := make([]byte, 256)
	_, err = r.Read(data)
	if err !=nil{
		errors.Wrapf(err, "could not read %s", r)
	}

	split_data := bytes.Split(data, []byte(", "))
	fmt.Printf("input data is  is %s\n", split_data)
	currency_value := 0

	for _, d := range split_data {
		bb := string(d)
		cv, err := strconv.Atoi(bb)
		if err != nil{
			errors.Wrapf(err, "cannot convert")
		}
		currency_value += cv

		fmt.Printf("adding %d\n", cv)
		fmt.Printf("current value is %v\n", currency_value)
	}

	fmt.Printf("Currency value is %v\n", currency_value)
}
