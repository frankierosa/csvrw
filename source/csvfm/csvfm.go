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
    // Inserting csv first row into slice columns in Upper Case.
	for _, v := range data[0][:] {
		c := strings.ToUpper(v)
		columns = append(columns, c)
	}
    // Appeding columns into newData slice
    for i := 0; i < 1; i++ {
        n := strings.Join(columns, " | ")
        newData = append(newData, n)
    }
    // creating a empty slice of string to separe column with the data.
    newData = append(newData, " ")

    // Columnize data content into slice newData
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