package main

import (
	"fmt"
	"log"

	"github.com/frankierosa/csvrw/csvreader"
)

var file = "data/customer.csv"

func main() {
	f, err := opencsvfile.OpFile(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(f))
}