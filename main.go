package main

import (
	"github.com/thatquietkid/go-todo-app/cmd"
	"github.com/thatquietkid/go-todo-app/internal/app"
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func readCSVFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func parseCSV(data []byte) (*csv.Reader, error) {
	reader := csv.NewReader(bytes.NewReader(data))
	return reader, nil
}

func processCSV(reader *csv.Reader) ([][]string, error) {
	var records [][]string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	return records, nil
}

func createCSVWriter(filename string) (*csv.Writer, *os.File, error) {
    f, err := os.Create(filename)
    if err != nil {
        return nil, nil, err
    }
    writer := csv.NewWriter(f)
    return writer, f, nil
}

func main() {
	// Read CSV file
	// Ensure the file exists and is readable
	data, err := readCSVFile("tasks.csv")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Parse CSV data
	reader, err := parseCSV(data)
	if err != nil {
		fmt.Println("Error creating CSV reader:", err)
		return
	}

	// Process CSV data
	records, err := processCSV(reader)
	if err != nil {
		fmt.Println("Error processing CSV data:", err)
		return
	}

	// Create app state
	state := &app.AppState{
		Records: records,
	}

	// Inject app state into context
	ctx := context.WithValue(context.Background(), "appState", state)

	// Set context on the root Cobra command
	cmd.RootCmd.SetContext(ctx)

	// Execute CLI
	if err := cmd.Execute(); err != nil {
		fmt.Println("Command execution error:", err)
		os.Exit(1)
	}
}
