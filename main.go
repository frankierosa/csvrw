package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/frankierosa/csvrw/source/asciiart"
	"github.com/frankierosa/csvrw/source/csvfm"
	"github.com/frankierosa/csvrw/source/csvreader"
	"github.com/frankierosa/csvrw/source/progressbar"

)

var (
    fileInfo *os.FileInfo
    err      error
	data [][]string
)


func main() {
	// Loading ascii art.
	asciiart.PrintAsciiArt()

	for {
	    // Reading from user input
		fmt.Print("CSV file location: ")
		reader := bufio.NewReader(os.Stdin)

		// Progress Bar will be trigger if the file is a csv.
		var bar Bar
		bar.NewOption(0, 100)
		for i := 0; i <= 100; i++ {
			time.Sleep(100 * time.Millisecond)
			bar.Play(int64(i))
		}
		bar.Finish()

	    // Reading a buffer from customer input and change it string.
	    file, err := reader.ReadString('\n')
	    if err != nil {
		    log.Fatal(err)
	    }
	    // Removing any whitespace from ReadString.
	    file = strings.TrimSuffix(file, "\n")
	    file = strings.TrimSuffix(file, "\r")

	    // Checking for file extesion.
	    fileExt := filepath.Ext(file)

		// Validate file ext before passing into CSVReader func.
		if fileExt == ".csv" || fileExt == ".CSV" {
			fmt.Println()

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