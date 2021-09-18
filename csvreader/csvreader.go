package csvreader

import (
	"bytes"
	"encoding/csv"
	"io/ioutil"
	"log"
)

// CSVreader func will open a csv file and return slice []byte if the contents.
func CSVreader(filename string) ([][]string, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
<<<<<<< HEAD
	
	// cover file []byte to io.Reader
	cToReader := bytes.NewReader(file)
	if err != nil {
		log.Fatal(err)
	}

	// Reader the cToReader "*Reader" return.
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
=======
	return file, err
}
>>>>>>> 4cab8b6 (test)
