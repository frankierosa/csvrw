package main

import (
	"fmt"
	"log"

	"github.com/frankierosa/csvrw/csvreader"
)

var file = "data/sample.csv"
var data = make(map[int][]string)

 type Person struct {
	 ID int
	 FirstName, LastName, CompanyName, Address, City, County, State, ZipCode, Phone, Email string

 }

 func (p Person) customerData(c map[int]string){
	 
 }

func main() {
	f, err := csvreader.CSVreader(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(f)
	for i, v := range f {
		append(slice []Type, elems ...Type)
		fmt.Println(i, v)
	}
}