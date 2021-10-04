package csvfm

import (
    "fmt"
    "os"
    "strings"
    "text/tabwriter"
)

var (
    columns []string
)



//CSVFormat display the content separe for each csv columns
// example form https://networkbit.ch/golang-column-print/
func CSVFormat(data [][]string) {
    
    // initialize tabwrite
    w := new(tabwriter.Writer)

    // minwidth, tabwidth, padding, padchar, flags
    w.Init(os.Stdout, 8, 8, 0, 't', 0)

    defer w.Flush()

    // get columns
    for _, d := range data[0][:] {
        columns = append(columns, d)
    }

    for _, c := range columns {
        fmt.Println(strings.Trim(fmt.Sprint(c), "[]"))
    }
}