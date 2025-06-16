# csvfix

A Go program to read a CSV file, process each row with a customizable `fixit` function, and write the results to a new CSV file compatible with Excel.

## Usage

```
go run main.go -in input.csv -out output.csv
```

- **input.csv**: Input CSV file (comma-separated, quoted strings allowed)
- **output.csv**: Output CSV file (Excel-compatible)

## How it works
- Reads the input CSV file.
- Each row is passed to the `fixit` function (to be implemented by you).
- The modified rows are written to the output CSV file.

## Customization
Edit the `fixit` function in `main.go` to implement your own row processing logic.
