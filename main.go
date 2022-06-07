/*
 * This file is part of csvrw
 * Copyright 2022 frankierosa
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published
 * by the Free Software Foundation, version 3 of the License, or (at your
 * option) any later version.
 *
 * This program is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
 * General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>.
 */

/* 
 * main.go:
 */
 
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/frankierosa/csvrw/source/asciiart"
	"github.com/frankierosa/csvrw/source/csvfm"
	"github.com/frankierosa/csvrw/source/csvreader"
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