// ============================================================================
// = utils.go																  =
// = 	Description: Useful utility functions like reverse, divisors,		  =
// = 		greatest common divisor, etc.		 							  =
// = 	Notes: OEIS often handles a lot of big numbers, so int64 is the data  =
// = 		type i used (unless something needed to be a decimal)			  =
// = 	Date: October 08, 2021												  =
// ============================================================================

package utils

import (
	"errors"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

// ############################ CONSTANTS ##############################
// ### this section holds all constants needed.
const (
	black = "\u001b[30m"
	red = "\u001b[31m"
	yellow = "\u001b[33m"
	green = "\u001b[32m"
	blue = "\u001b[34m"
	reset = "\u001b[0m"
)

// ############################ ERRORS #################################
// ### this section handles error checking, printing, etc.

// checks and error and panics. Used primarily for debugging
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

// handles an error in a pretty way for the user.
func HandleError(e error) {
	if e != nil {
		PrintError(e.Error())
		os.Exit(1)
	}
}

// used to issue an error about a non-positive sequence length
func PositiveError(seqname string) {
	msg := "error in sequence " + seqname + ": seqlen must be positive"
	HandleError(errors.New(msg))
}

// used to issue an error about sequence lengths that will cause overflow
func OverflowError(seqname string, seqlen int64) {
	slen := strconv.FormatInt(seqlen, 10)
	msg := "error in sequence " + seqname + ": a seqlen > " + slen + " will result in overflow. Exiting..."
	HandleError(errors.New(msg))
}

// used to issue an error about a seqlen that is too small for the sequence
// say the sequence is Fibonacci, and the user gives a seqlen of 1. hard coding of sequence numbers
// will cause golang to issue an error when doing a[1] = 1, for instance
func TooSmallError(seqname string, seqlen int64) {
	slen := strconv.FormatInt(seqlen, 10)
	msg := "error in sequence " + seqname + ": a seqlen < " + slen + " is too small. Exiting..."
	HandleError(errors.New(msg))
}

// ############################ WARNINGS ###########################
// used to issue a warning about sequence lengths that will take a long time to compute
// a "long time" is typically more than 5 seconds (which isn't that long, but users are impatient)
func LongCalculationWarning(seqname string, seqlen int64) {
	slen := strconv.FormatInt(seqlen, 10)
	msg := "Warning: Sequence " + seqname + " with a seqlen > " + slen + " will take time to compute"
	PrintWarning(msg)
}

// used to issue a warning about a sequence that uses big.Int fields instead of int64
func BigIntWarning(seqid string, seqlen int64) {
	slen := strconv.FormatInt(seqlen, 10)
	msg := "Warning: Sequence " + seqid + " has big.Int fields instead of int64. This *MAY* lead to accuracy issues, especially for seqlen > " + slen + "."
	PrintWarning(msg)
}

// used to issue a warning about the accuracy of a sequence using big.Float fields
func AccuracyWarning(seqid string) {
	msg := "Warning: Calculation for sequence " + seqid + " relies on decimal rounding, which will probably lead to some accuracy issues"
	PrintWarning(msg)
}

// ############################ PRINTING FUNCTIONS #########################
// ### this section contains all printing functions

func PrintDebug(msg string) {
	fmt.Println(blue + msg + reset)
}

func PrintInfo(msg string) {
	fmt.Println(green + msg + reset)
}

func PrintWarning(msg string) {
	fmt.Println(yellow + msg + reset)
}

func PrintError(msg string) {
	fmt.Println(red + msg + reset)
}

func PrintSequence(seqid string, a []int64, startidx int64) {
	if a == nil {
		return
	}
	if seqid != "" {
		PrintInfo("~~~~~ SEQUENCE " + seqid + " ~~~~~")
	}
	fmt.Println("n\ta(n)")
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d\t%d\n", startidx, a[i])
		startidx++
	}
}

func PrintBigSequence(seqid string, a []*big.Int, startidx int64) {
	if a == nil {
		return
	}
	if seqid != "" {
		PrintInfo("~~~~~ SEQUENCE " + seqid + " ~~~~~")
	}
	fmt.Println("n\ta(n)")
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d\t%d\n", startidx, a[i])
		startidx++
	}
}
