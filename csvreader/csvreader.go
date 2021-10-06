package csvreader

import (
	"bytes"
	"encoding/csv"
	"io/ioutil"
	"log"
)

// CSVReader open a csv file and return all the content.
func CSVReader(filename string) ([][]string, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	// cover file []byte to io.Reader
	cToReader := bytes.NewReader(file)
	if err != nil {
		log.Fatal(err)
	}
	// Readering from cToReader.
	r := csv.NewReader(cToReader)
	if err != nil {
		log.Fatal(err)
	}
	// reading all the records from r.
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return records, err
}