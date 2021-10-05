package csvfm

import (
	"fmt"
	"strings"

	"github.com/ryanuber/columnize"
)

var (
	columns []string
    newData []string
)

func CSVFormat(data [][]string) {

	// Inserting csv columns with "|"
	for _, v := range data[0][:] {
		c := strings.ToUpper(v)
		columns = append(columns, c)
	}
	/*
	
	for _, v := range columns {
		for i := 0; i < 1; i++ {
		c := strings.Join(v, " | ")
		newData = append(newData, c)
		}
	}
	*/

    // Columnize content into slice newData
    // For more detail go to https://github.com/ryanuber/columnize
    for _, d := range data[1:] {
        for i := 0; i < 1; i++ {
        n := strings.Join(d, " | ")
        newData = append(newData, n)
        }
    }


    displayData := columnize.SimpleFormat(newData)
    fmt.Println(displayData)   

}
