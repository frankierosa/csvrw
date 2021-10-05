package csvfm

import (
	"fmt"
	"strings"

	"github.com/ryanuber/columnize"
)

var (
    columns []string
)

//CSVFormat display the content separe for each csv columns
// example form https://networkbit.ch/golang-column-print/
func CSVFormat(data [][]string) {

    // Inserting csv columns with "|"
     for i := 0; i < 1; i++ {
        c := strings.Join(data[0][:], " | ")
        columns = append(columns, c)
    }

    // Columnize Columns from slice columns
    // For more detail go to https://github.com/ryanuber/columnize
    result := columnize.SimpleFormat(columns)
    fmt.Println(result)
}