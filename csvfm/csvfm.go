package csvfm

import (
	"fmt"
	"os"
	"text/tabwriter"
)



//CSVFormat display the content separe for each csv columns
// example form https://networkbit.ch/golang-column-print/
func CSVFormat(data [][]string) {
	
	// initialize tabwrite
	w := new(tabwriter.Writer)

	// minwidth, tabwidth, padding, padchar, flags
	w.Init(os.Stdout, 8, 8, 0, 't', 0)

	defer w.Flush()

	// read the data
	for _, d := range data {
		for _, v := range d {
			fmt.Println(v)
		}
	}
}