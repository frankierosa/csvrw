package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

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

// renderbar func for Progress bar
func renderbar(count, total int) {
    barwidth := 30
    done := int(float64(barwidth) * float64(count) / float64(total))

    fmt.Printf("Progress: \x1b[33m%3d%%\x1b[0m ", count*100/total)
    fmt.Printf("[%s%s]",
        strings.Repeat("=", done),
        strings.Repeat("-", barwidth-done))
}

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
			fmt.Println()
			total := 50
            for i := 1; i <= total; i++ {
				//<ESC> means ASCII "escape" character, 0x1B
				fmt.Print("\x1b7")   // Save the cursor position, save the cursor and Attrs < ESC > 7
                fmt.Print("\x1b[2k") // Clear the content erase line of the current line < ESC > [2K]
                renderbar(i, total)
                time.Sleep(50 * time.Millisecond)
                fmt.Print("\x1b8") // Recovery cursor position recovery cursor and Atrs < ESC > 8
            }
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