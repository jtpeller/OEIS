# OEIS

The ongoing quest to program every sequence in the OEIS database (in Golang)

## Content

- sequences -- The folder containing the seq package, which contains all programmed sequences
- utils -- Contains any and all utility functions that are very common (say, a PrintSequence function)
- go.mod -- Handles the OEIS module
- main.go -- The file containing main
- README.md -- The file you're reading right now

## Notes

Each of the sequence functions (those functions starting with `A...`) will return:
 - The integer sequence (specifically `[]int64`)
 - The offset (aka starting position or starting index)

The various helper functions are left as public in the event that it will eventually be
needed elsewhere (which is common in OEIS)

## Usage
