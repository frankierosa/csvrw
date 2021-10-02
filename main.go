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
)

var (
    fileInfo *os.FileInfo
    err      error

	// testing proper
	//file = "data/sample.csv"
)

func main() {
	// Loading ascii art.
	asciiart.PrintAsciiArt()

	//Check csv file ext and passing into CSVreader. 
	//If file is csv extesion, it will invoked the func CSVreader to read the file. 
	//If the file is not csv extesion, will ask for the file again or exit the program.
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

	    // Getting file extesion.
	    fileExt := filepath.Ext(file)
		if fileExt == ".csv" || fileExt == ".CSV" {
			file, err := csvreader.CSVreader(file)
	        if err != nil {
		        log.Fatal(err)
	        }
			for i,v := range file {
				fmt.Println(i, v)
			}
			break
	    } else if fileExt != ".csv" || fileExt != ".CSV" {
			log.Printf("incorrect path or file extesion")
			continue
		}
	}
}