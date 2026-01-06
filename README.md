# Multiplication Table Generator

A simple command-line tool written in Go that generates multiplication tables for any given number.

## Features

- Generate multiplication tables (1-10) for any integer
- Input validation to prevent integer overflow
- Support for positive, negative, and zero values

## Usage

### Run the program

```bash
go run main.go
```

You'll be prompted to enter a number, and the program will display its multiplication table:

```
Welcome to the Multiplcation Table Generator!
Enter a number: 7

Multiplication table of 7 is:
--------------------------------
7 x 1 = 7
7 x 2 = 14
7 x 3 = 21
7 x 4 = 28
7 x 5 = 35
7 x 6 = 42
7 x 7 = 49
7 x 8 = 56
7 x 9 = 63
7 x 10 = 70
--------------------------------
```

### Build and run

```bash
go build -o multiplication-table
./multiplication-table
```

## Testing

Run all tests:

```bash
go test -v
```

Run tests with coverage:

```bash
go test -cover
```

Run benchmarks:

```bash
go test -bench=.
```

## Project Structure

```
multiplication-table/
├── go.mod          # Go module definition
├── main.go         # Main application code
├── main_test.go    # Test cases
└── README.md       # This file
```

## License

MIT

