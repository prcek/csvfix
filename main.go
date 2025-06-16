package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

// fixit processes a row and returns the modified row.
func fixit(row []string) []string {
	if len(row) == 17 {
		// row[3] = integer part, row[4] = decimal part
		intPart := row[3]
		decPart := row[4]
		floatStr := intPart + "." + decPart
		row[3] = floatStr
		// Remove row[4]
		row = append(row[:4], row[5:]...)
	}

	if len(row) != 16 {
		fmt.Fprintf(os.Stderr, "Row length is not 16: %d\n", len(row))
		return row // Return the row as is if it doesn't match expected length
	}
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
