package main

import (
	"fmt"
	"log"

	"github.com/frankierosa/csvrw/csvreader"
)

var file = "data/sample.csv"

func main() {
	f, err := csvreader.CSVreader(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(f)
	for i, v := range f {
		fmt.Println(i, v)
	}
}