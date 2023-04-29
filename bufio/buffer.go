package bufio

import (
	"bufio"
	"encoding/csv"
	"os"
)

func OpenCsv(fileName string, useBuffer bool) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var csvReader *csv.Reader

	if useBuffer {
		reader := bufio.NewReaderSize(file, 4096*128)
		csvReader = csv.NewReader(reader)
	} else {
		csvReader = csv.NewReader(file)
	}

	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}
