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
	"math"
	"math/big"
	"os"
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

// ############################ ERROR CHECKING ##############################
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

func PrintSequence(title string, a []int64, startidx int64) {
	// ensure startidx isn't negative
	if startidx < 0 {
		panic("startidx cannot be negative")
	}

	// convert uint 
	if title != "" {
		fmt.Println(title)
	}
	fmt.Println("n\ta(n)")
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d\t%d\n", startidx, a[i])
		startidx++
	}
}

// ############################### CHECKERS ###################################
// ### this section checks if a number has a specific property

// Checks if the given number is prime
func IsPrime(num int64) bool {
	return big.NewInt(num).ProbablyPrime(0)
}

// Checks if the given number n is a prime power of k
func IsPrimePower(n int64, k int64) bool {
	nf := float64(n)
	kf := float64(k)
	return (math.Log(nf) / math.Log(kf) == math.Floor(math.Log(nf) / math.Log(kf)))
}

// ########################## CALCULATION SECTION #############################
// ### this section calculates some property of a number (factor count,
// ### prime factorization, etc.)

// Compute the number of digits of the given number
func GetDigits(n int64) int64 {
	count := int64(0)
	for n != 0 {
		count++
		n /= 10
	}
	return count
}

// Calculate the integer square root of a number
func Isqrt(num int64) (int64, error) {
	return int64(math.Floor(math.Sqrt(float64(num)))), nil
}

// Calculate the number of factors of num
func GetFactorCount(num int64) int64 {
	// calculate factor count
	count := int64(0)
	for i := int64(1); i <= num; i++ {
		if num % i == 0 {
			count++
		}
	}
	return count
}

// Computes the factors of num
func Factors(num int64) []int64 {
	factors := make([]int64, 0)
	for i := int64(1); i <= num; i++ {
		if num % i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}

// Calculates the prime factorization of num
func PrimeFactorization(num int64) []int64 {
	// initialize array
	primefact := make([]int64, 0)

	// the number of 2s
	for num % 2 == 0 {
		primefact = append(primefact, 2)
		num /= 2
	}

	// num is now odd. check 3s and beyond
	for i := int64(3); i * i <= num; i += 2 {
		// get factors until n is zero
		for num % i == 0 {
			primefact = append(primefact, i)
			num /= i
		}
	}

	// what if num is prime?
	if num > 2 {
		primefact = append(primefact, num)
	}

	return primefact
}

// given two numbers, it will calculate the greatest common divisor
func GreatestCommonDivisor(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return GreatestCommonDivisor(b, a % b)
}

// given a number num, it will compute Euler's Totient of the number
func EulerTotient(num int64) int64 {
	val := int64(0)
	for i := int64(0); i < num; i++ {
		if GreatestCommonDivisor(i, num) == 1 {
			val++
		}
	}
	return val
}

func PrimeFloor(arr []int64, n int64) int64 {
	// lazy case
	if n == 1 {
		return n
	}

	// loop thru k (counter for exponent)
	foo := int64(0)
	found := false
	for k := int64(2); k < n; k++ {
		// check if k is a valid prime power
		for i := int64(0); i <= k; i++ {
			if arr[i] == k {
				found = true
				break
			}
			found = false
		}

		// if k is a valid prime power, perform the algo
		if found {
			foo = n
			if IsPrimePower(n, k) {
				return foo
			}
		} else {
			continue
		}
	}
	// leaving the loop means it isn't a prime power of k > 1
	// check for prime power k == 1
	if IsPrime(n) {
		return n
	} else {
		return PrimeFloor(arr, n + 1)
	}
}


// ########################## GENERATOR FUNCTIONS #############################
// ### given a number, it will generate a sequence with some quality up to that
// ### number. things like primes, evens, odds, etc.

// generates the sequence of primes; count = num
func Primes(seqlen int64) ([]int64, error) {
	// check valid num
	if seqlen <= 0 {
		return nil, errors.New("utils error in Primes(): num cannot be negative")
	}

	primes := make([]int64, 0)
	num := int64(0)
	for i := int64(0); i < seqlen; i++ {
		if IsPrime(int64(num)) {
			primes = append(primes, num)
		}
		num++
	}
	return primes, nil
}

