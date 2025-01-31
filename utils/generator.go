// ============================================================================
// = generator.go															  =
// = 	Description		Generator functions like primes.					  =
// = 	Date			December 07, 2021									  =
// ============================================================================

package utils

import (
	"math"
	"math/big"

	gb "github.com/jtpeller/gobig"
)

// ########################## GENERATOR FUNCTIONS #############################
// ### given a number, it will generate a sequence with some quality up to that
// ### number. things like primes, evens, odds, etc.

// given a sequence, it will return the bisection of that sequence
func Bisection(seq []int64) []int64 {
	a := make([]int64, 0)
	for i := 0; i < len(seq); i+=2 {
		a = append(a, seq[i])
	}
	return a
}

func BigBisection(seq []*big.Int) []*big.Int {
	a := CreateSlice(0)
	for i := 0; i < len(seq); i+=2 {
		a = append(a, seq[i])
	}
	return a
}

// counts the digits of a given number
func countDigits(num int64) int64 {
	a := num
	count := int64(0)
	for a != 0 {
		a /= 10
		count++
	}
	return count
}

// grab the digit at a specific place. Helper for Digits()
func digit(num, idx int64) int64 {
    r := num % int64(math.Pow(10, float64(idx)))
    return r / int64(math.Pow(10, float64(idx-1)))
}

// separates a number into its digits
func Digits(num int64) []int64 {
	a := make([]int64, 0)
	ndigits := countDigits(num)
	for i := int64(1); i <= ndigits; i++ {
		a = append(a, digit(num, i))
	}
	return a
}

// Computes the factors (divisors) of num
func Factors(num int64) []int64 {
	factors := make([]int64, 0)
	for i := int64(1); i <= num; i++ {
		if num % i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}

// generates the Gamma function output, which is just factorials shifted over by 1 idx
func Gamma(seqlen int64) ([]*big.Int, int64) {
	a := CreateSlice(seqlen+1)
	for i := int64(0); i < seqlen; i++ {
		a[i+1] = Fact(big.NewInt(i))
	}
	return a, 1		// Gamma "starts" at 1
}

// initializes a slice with a set of nums 
func InitIslice (seqlen int64, init []int64) []int64 {
	a := make([]int64, seqlen)
	bound := int64(len(init))
	if seqlen < bound {
		bound = seqlen
	}

	// init
	for n := int64(0); n < bound; n++ {
		a[n] = init[n]
	}
	return a
}

// initializes a slice of *big.Int
func InitBslice(seqlen int64, init []*big.Int) []*big.Int {
	a := CreateSlice(seqlen)
	bound := int64(len(init))
	if seqlen < bound {
		bound = seqlen
	}

	// init
	for n := int64(0); n < bound; n++ {
		a[n] = init[n]
	}
	return a
}

// Calculates the isqrt of an array
func Isqrtarray(arr []int64) []int64 {
	a := make([]int64, len(arr))
	for i := 0; i < len(arr); i++ {
		val := Isqrt(arr[i])
		a[i] = val
	}
	return a
}

// Computes the multiples (i.e. given n and seqlen, it will generate 
// the multiples of n)
func Multiples(n int64, seqlen int64) []int64 {
	a := make([]int64, seqlen)
	for i := int64(0); i < n; i++ {
		a[i] = n * i
	}
	return a
}

// powers
func Power(seqlen int64, e *big.Int) []*big.Int {
	a := CreateSlice(seqlen)
	for i := int64(0); i < seqlen; i++ {
		a[i] = gb.Pow(gb.New(i), e)
	}
	return a
}

// generates the sequence of primes; count = num
func Primes(seqlen int64) []int64 {
	primes := make([]int64, 0)
	num := int64(0)
	for i := int64(0); i < seqlen; {
		if IsPrime(num) {
			primes = append(primes, num)
			i++
		}
		num++
	}
	return primes
}

// performs Primes(), but with big.Int instead
func PrimesBig(seqlen int64) []*big.Int {
	primes := CreateSlice(0)
	num := gb.New(0)
	for i := int64(0); i < seqlen; {
		if IsBigPrime(num) {
			primes = append(primes, num)
			i++
		}
		num = gb.Add(num, gb.New(1))
	}
	return primes
}

// generates a sequence calculating the # of positive integers <= 2^n 
// of the form px^2 + qy^2 
func Repr(seqlen, p, q, init int64) []*big.Int {
	a := CreateSlice(seqlen)
	a[0] = big.NewInt(init)
	for n := int64(1); n < seqlen; n++ {
		nf := float64(n)
		count := a[n-1]
		for k := int64(math.Pow(2, nf-1) + 1); k <= int64(math.Pow(2, nf)); k++ {
			if IsRepr(k, p, q) {
				count = big.NewInt(0).Add(count, big.NewInt(1))
			}
		}
		a[n] = count
	}
	return a
}

// shifts an array over by some amount
func Shift(a []int64, amt int) []int64 {
	out := make([]int64, len(a)+amt)
	for i, v := range a {
		out[i+amt] = v
	}
	return out
}

// performs Shift(), but with []*big.Int
func ShiftBig(a []*big.Int, amt int) []*big.Int {
	out := CreateSlice(int64(len(a)+amt))
	for i, v := range a {
		out[i+amt] = v
	}
	return out
}

// calculates the sum of squares of the digits of num
func SumSquares(num int64) int64 {
	sum := int64(0)
	digits := Digits(num)
	for _, d := range digits {
		sum += int64(math.Pow(float64(d), 2))
	}
	return sum
}
