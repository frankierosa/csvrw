package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/frankierosa/csvrw/asciiart"
	"github.com/frankierosa/csvrw/csvreader"
	"github.com/frankierosa/csvrw/csvfm"
)

var (
    fileInfo *os.FileInfo
    err      error
	data [][]string

	// testing proper
	//file = "data/sample.csv"
)

func main() {
	// Loading ascii art.
	asciiart.PrintAsciiArt()

	for {
		fmt.Print("Enter the csv file location: ")
	
	    // Reading from user input
	    reader := bufio.NewReader(os.Stdin)

	    // Reading buffer from customer input and change it string.
	    file, err := reader.ReadString('\n')
	    if err != nil {
		    log.Fatal(err)
	    }

	     // removing any whitespace from ReadString.
	    file = strings.TrimSuffix(file, "\n")
	    file = strings.TrimSuffix(file, "\r")	

	    // Checking file extesion.
	    fileExt := filepath.Ext(file)
		if fileExt == ".csv" || fileExt == ".CSV" {
			
			// Pasing the file into CSVreader func to collect all the records.
			file, err := csvreader.CSVReader(file)

	        if err != nil {
		        log.Fatal(err)
	        }

			// creating appending into slice data.
			for _, v := range file {
				data = append(data, v)
			}
			break
	    } else if fileExt != ".csv" || fileExt != ".CSV" {
			log.Printf("incorrect path or file extesion")
			continue
		} 
	}
}