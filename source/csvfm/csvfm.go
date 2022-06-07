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
 * csvfm.go:
 */
 
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