package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	_ "github.com/frankierosa/csvrw/asciiart"
	"github.com/frankierosa/csvrw/csvreader"
)

var (
    fileInfo *os.FileInfo
    err      error
	sortData []string
	mData = make(map[int]string)

	// testing proper
	file = "data/sample.csv"
)



func main() {
	// Loading ascii art.
	//asciiart.PrintAsciiArt()

	/*fmt.Print("Enter the csv file location: ")
	
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
	*/

	// Getting file extesion.
	fileExt := filepath.Ext(file)

	//Check if file is csv ext and passing into CSVreader. 
	//If file is csv extesion, it will invoked the func CSVreader to read the file. 
	//If the file is not csv extesion, will ask for the file again or exit the program.
	for {
		if fileExt == ".csv" || fileExt == ".CSV" {
			file, err := csvreader.CSVreader(file)
	        if err != nil {
		        log.Fatal(err)
	        }
			for _, r := range file[0][:len(file[0])] {
				
				// appending csv file header column to a new slice sortData
				// in this way we can sort by these column.
				sortData = append(sortData, r)
			}

			fmt.Println("Items list: ")
			for i, v := range sortData {
				i = i + 1
				fmt.Println(i, v)
				
				// Adding index and data into map "mData"
				mData[i] = v
			}
			fmt.Print("\nFilter by: ")

			// Reading user input
	        reader := bufio.NewReader(os.Stdin)

	        // Reading buffer from customer input and change it string.
	        userSort, err := reader.ReadString('\n')
	        if err != nil {
		        log.Fatal(err)
	        }

	       // removing any whitespace from ReadString.
	       userSort = strings.TrimSuffix(userSort, "\n")
	       userSort = strings.TrimSuffix(userSort, "\r")

		   // AtoI func that will cover string to int to look for the index form the mData
		   sToI, err := strconv.Atoi(userSort)
		   if err != nil {
			   log.Fatal(err)
		   }

		   // fetching data from csv file by user selection from the sort list.
			for _, v := range file[sToI][:] {
				fmt.Println(v)
			}

	    } else if fileExt != ".csv" || fileExt != ".CSV" {
			log.Printf("incorrect file extesion: %s", fileExt)
		} 
		break
	}
}