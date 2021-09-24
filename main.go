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
	for _, v := range f {
		fmt.Println(v)
	}
}
