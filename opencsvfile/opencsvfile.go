package opencsvfile

import (
	"io/ioutil"
	"log"
)

// OpFile func will open the file and read
// will get filenem string and return []byte, error
func OpFile(filename string) ([]byte, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return file, err
}