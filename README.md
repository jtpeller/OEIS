# OEIS

The ongoing quest to program every sequence in the OEIS database (in Golang)

## Content

- sequences -- The folder containing the seq package, which contains all programmed sequences
- utils -- Contains any and all utility functions that are very common (say, a PrintSequence function). Also includes any common calculations or generator functions for common sequences (such as primes or the factors of a number).
- go.mod -- Handles the OEIS module
- main.go -- The file containing main
- README.md -- The file you're reading right now

## Notes

Each of the sequence functions (those functions starting with `A...`) will return:
 - The integer sequence that is produced. Type: `[]int64 || *big.Int`
 - The offset (aka starting position or starting index). Type: `int64`

The various helper functions are left as public in the event that it will eventually be
needed elsewhere (which is common in OEIS)

## Usage

Run the program with `go run main.go` and some options

Use `go run main.go -h` or `go run main.go --help` for more information.

Options:
 - `-seq` -- Give the sequence ID (A000002 for example)
 - `-seqlen` -- Give the number of elements to generate. There will be limits on some of the sequences due to overflow.