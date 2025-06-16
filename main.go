package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

// fixit processes a row and returns the modified row. To be implemented later.
func fixit(row []string) []string {
	// TODO: implement your logic here
	return row
}

func main() {
	inputPath := flag.String("in", "", "Input CSV file path")
	outputPath := flag.String("out", "", "Output CSV file path")
	flag.Parse()

	if *inputPath == "" || *outputPath == "" {
		fmt.Println("Usage: csvfix -in input.csv -out output.csv")
		os.Exit(1)
	}

	inFile, err := os.Open(*inputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening input file: %v\n", err)
		os.Exit(1)
	}
	defer inFile.Close()

	reader := csv.NewReader(inFile)
	reader.Comma = ','
	reader.LazyQuotes = true

	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading CSV: %v\n", err)
		os.Exit(1)
	}

	outFile, err := os.Create(*outputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating output file: %v\n", err)
		os.Exit(1)
	}
	defer outFile.Close()

	writer := csv.NewWriter(outFile)
	writer.Comma = ','
	writer.UseCRLF = false // Excel on macOS/Linux is fine with LF

	for _, row := range rows {
		fixed := fixit(row)
		if err := writer.Write(fixed); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing row: %v\n", err)
		}
	}
	writer.Flush()

	if err := writer.Error(); err != nil {
		fmt.Fprintf(os.Stderr, "Error flushing CSV: %v\n", err)
		os.Exit(1)
	}
}
