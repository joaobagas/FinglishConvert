package src

import (
	"encoding/csv"
	"os"
)

func ReadCsv(filePath string) map[string]string {
	file, err := os.Open(filePath)
	if err != nil {
		return make(map[string]string)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return make(map[string]string)
	}

	ret := make(map[string]string)
	for _, record := range records {
		ret[record[0]] = record[1]
	}

	return ret
}
