package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"path/filepath"

	"github.com/frankierosa/csvrw/asciiart"
	"github.com/frankierosa/csvrw/csvreader"
)


func main() {
	// Loading ascii art.
	asciiart.PrintAsciiArt()

	fmt.Print("Enter the csv file location: ")

	// Reading from user input
	reader := bufio.NewReader(os.Stdin)

	// Reading buffer from customer input and change it string.
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
<<<<<<< HEAD

	// removing any whitespace from ReadString.
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")	

	// Get file extesion.
	fileExt := filepath.Ext(input)

	//Check if file is csv ext and passing into CSVreader. 
	//If file is csv extesion, it will invoked the func CSVreader to read the file. 
	//If the file is not csv extesion, will ask for the file again or exit the program.
	for {
		if fileExt == ".csv" || fileExt == ".CSV" {
			f, err := csvreader.CSVreader(input)
	        if err != nil {
		        log.Fatal(err)
	        }
	       for _, v := range f {
		   fmt.Println(v)
	      }
	    }
		break
	} 
=======
	fmt.Println(f)
	for i, v := range f {
		//customer := v[1:]
		fmt.Println(i, v)
	}
>>>>>>> e8fbf6e (test2)
}