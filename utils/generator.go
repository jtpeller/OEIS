// ============================================================================
// = generator.go															  =
// = 	Description: Useful generator functions like primes.				  =
// = 	Date: December 07, 2021												  =
// ============================================================================

package utils

import "math/big"

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
