package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/frankierosa/csvrw/asciiart"
	"github.com/frankierosa/csvrw/csvfm"
	"github.com/frankierosa/csvrw/csvreader"
)

var (
    fileInfo *os.FileInfo
    err      error
	data [][]string

	// testing proper.
	file = "data/sample.csv"
)

func main() {
	// Loading ascii art.
	asciiart.PrintAsciiArt()

	for {
		fmt.Print("Enter the csv file location: ")
		
	    // Reading from user input
		reader := bufio.NewReader(os.Stdin)

	    // Reading a buffer from customer input and change it string.
	    file, err := reader.ReadString('\n')
	    if err != nil {
		    log.Fatal(err)
	    }
	    // removing any whitespace from ReadString.
	    file = strings.TrimSuffix(file, "\n")
	    file = strings.TrimSuffix(file, "\r")

	    // Checking for file extesion.
	    fileExt := filepath.Ext(file)

		// Validate file ext before passing into CSVReader func.
		if fileExt == ".csv" || fileExt == ".CSV" {
			file, err := csvreader.CSVReader(file)
	        if err != nil {
		        log.Fatal(err)
	        }
			// Appending content into slice data.
			for _, v := range file {
				data = append(data, v)
			}
			break

	    } else {
			log.Printf("incorrect path or file extesion")
			continue
		} 
	}
	// passing the data into CSVFormat Func.
	csvfm.CSVFormat(data)
}