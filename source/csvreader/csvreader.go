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
 * csvreader.go:
 */
 
 package csvreader

import (
	"bytes"
	"encoding/csv"
	"io/ioutil"
	"log"
)

// CSVReader will open a csv file and return all the content.
func CSVReader(filename string) ([][]string, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	// Convert the file from []byte to io.Reader
	cToReader := bytes.NewReader(file)
	if err != nil {
		log.Fatal(err)
	}
	// Readering from cToReader.
	r := csv.NewReader(cToReader)
	if err != nil {
		log.Fatal(err)
	}
	// Reading all the records from r and return back.
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return records, err
}