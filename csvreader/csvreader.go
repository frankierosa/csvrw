package csvreader

import (
	"io/ioutil"
	"log"
)

// CSVreader func will read a csv file and print the content.
func CSVreader(filename string) ([]byte, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return file, err
}
