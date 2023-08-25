package helper

import (
	"encoding/csv"
	"os"
)

func ReadCSVAll(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	rows, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

	return rows
}
