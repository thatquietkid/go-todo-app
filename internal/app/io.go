package app

import (
	"encoding/csv"
	"fmt"
	"os"
)

// WriteCSVFile writes the records back to the given CSV file
func WriteCSVFile(filename string, records [][]string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("could not create file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.WriteAll(records); err != nil {
		return fmt.Errorf("error writing records to CSV: %w", err)
	}

	return nil
}
