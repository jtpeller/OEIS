# OEIS

## Overview

The ongoing quest to program every sequence in the OEIS database (in Golang)

## Content

- `sequences` -- The folder containing the seq package, which contains all programmed sequences
- `utils` -- Contains any and all utility functions that are very common (say, a PrintSequence function). Also includes any common calculations or generator functions for common sequences (such as primes or the factors of a number).
- `go.mod` -- Handles the OEIS module
- `main.go` -- The file containing main
- `README.md` -- The file you're reading right now

## Notes

Each of the sequence functions (those functions starting with `A...`) will return:

- The integer sequence that is produced. Type: `[]int64 || *big.Int`
- The offset (aka starting position or starting index). Type: `int64`

My strategy is not completing 100% of every sequence in order, but rather program as many of the OEIS sequences as possible. There's ~350 *thousand* sequences so my goal is to just get as many programmed as possible.

## Usage

Run the program with `go run main.go` and some options. For example:

```sh
go run main.go -seq A000045 -seqlen 50 -time
```

Use `go run main.go -h` or `go run main.go --help` for more information.

Options:

- `-seq` -- Give the sequence ID (A000002 for example)
- `-seqlen` -- Give the number of elements to generate. There may be limits on some of the sequences due to overflow or warnings due to rounding inaccuracies or lengthy computations.
