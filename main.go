package main

import (
	"fmt"
	"log"

	"github.com/frankierosa/csvrw/opencsvfile"
)

var file = "data/customer.csv"

func main() {
	f, err := opencsvfile.OpFile(file)
	if err != nil {
		log.Fatal(err)
	}
}