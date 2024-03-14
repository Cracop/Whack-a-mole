package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

func exportar(data [][]float64, path string, filename string) {

	// Create the CSV file in the specified directory

	// Create the directory if it doesn't exist
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		panic(err)
	}
	file, err := os.Create(path + filename + ".csv")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write data to the CSV file
	for _, row := range data {
		stringRow := make([]string, len(row))
		for i, v := range row {
			stringRow[i] = strconv.FormatFloat(v, 'f', -1, 64) // Convert float64 to string
		}
		err := writer.Write(stringRow)
		if err != nil {
			panic(err)
		}
	}
}
