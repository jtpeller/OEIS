// ============================================================================
// = generator.go
// = 	Description		Generator functions like primes.
// = 	Date			December 07, 2021
// ============================================================================

package utils

import (
	"fmt"
	"math"
)

// ########################## GENERATOR FUNCTIONS #############################
// ### given a number, it will generate a sequence with some quality up to that
// ### number. things like primes, evens, odds, etc.

// given a sequence, it will return the bisection of that sequence
func Bisection(seq []int64) []int64 {
	a := make([]int64, 0)
	for i := 0; i < len(seq); i += 2 {
		a = append(a, seq[i])
	}
	return a
}

func BisectionBig(seq []*bint) []*bint {
	a := iSlice(0)
	for i := 0; i < len(seq); i += 2 {
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

// Computes ALL factors (divisors) of num
func Factors(num int64) []int64 {
	factors := make([]int64, 0)
	for i := int64(1); i <= num; i++ {
		if num%i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}

// Computes ALL factors (divisors) of num.
// Includes num in the result if includeNum is true.
func FactorsBig(num *bint, includeNum bool) []*bint {
	factors := iSlice(0)
	for i := inew(1); lt(i, num); inc(i) {
		if equals(mod(num, i), zero()) { // n % i == 0
			factors = append(factors, add(zero(), i))
		}
	}
	if includeNum {
		factors = append(factors, num)
	}
	return factors
}

// generates the Gamma function output, which is just factorials shifted over by 1 idx
func Gamma(seqlen int64) ([]*bint, int64) {
	a := iSlice(seqlen + 1)
	for i := int64(0); i < seqlen; i++ {
		a[i+1] = fact(inew(i))
	}
	return a, 1 // Gamma "starts" at 1
}

// initializes a slice with a set of nums
func InitIslice(seqlen int64, init []int64) []int64 {
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

// initializes a slice of *bint
func InitBslice(seqlen int64, init []*bint) []*bint {
	a := iSlice(seqlen)
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

// Generates a(n) = n^e
func Exponents(seqlen int64, e *bint) []*bint {
	a := iSlice(seqlen)
	for i := int64(0); i < seqlen; i++ {
		a[i] = pow(inew(i), e)
	}
	return a
}

// Generates a "nacci" sequence, where it adds the kth previous values
// e.g., Fibonacci generated by Nacci(seqlen, 2, true), or hexanacci by (seqlen, 6, false)
func Nacci(seqlen int64, k int64, firstIsZero bool) []*bint {
	// error checking
	if k <= 0 {
		return nil
	}

	a := iSlice(seqlen)
	for n := int64(0); n < seqlen; n++ {
		if firstIsZero && n == 0 {
			a[n] = zero()
		} else if n < k {
			a[n] = inew(1)
		} else {
			sum := zero()
			for i := k; i > 0; i-- {
				sum = add(sum, a[n-i])
			}
			a[n] = sum
		}
	}
	return a
}

// generates a(n) = e^n
func Powers(seqlen int64, e *bint) []*bint {
	a := iSlice(seqlen)
	for i := int64(0); i < seqlen; i++ {
		a[i] = pow(e, inew(i))
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
func PrimesBig(seqlen int64) []*bint {
	primes := iSlice(0)
	num := zero()
	for i := int64(0); i < seqlen; {
		if IsBigPrime(num) {
			primes = append(primes, num)
			i++
		}
		num = add(num, inew(1))
	}
	return primes
}

// generates a sequence calculating the # of positive integers <= 2^n
// of the form px^2 + qy^2
func Repr(seqlen, p, q, init int64) []*bint {
	a := iSlice(seqlen)
	a[0] = inew(init)
	for n := int64(1); n < seqlen; n++ {
		nf := float64(n)
		count := a[n-1]
		for k := int64(math.Pow(2, nf-1) + 1); k <= int64(math.Pow(2, nf)); k++ {
			if IsRepr(k, p, q) {
				count = inew(0).Add(count, inew(1))
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

// performs Shift(), but with []*bint
func ShiftBigSliceRight(a []*bint, amt int) []*bint {
	out := iSlice(int64(len(a) + amt))
	for i, v := range a {
		out[i+amt] = v
	}
	return out
}

// shifts slice left. less Shift(), but with []*bint
func ShiftBigSliceLeft(a []*bint, amt int) []*bint {
	newlen := int64(len(a) - amt)
	out := iSlice(newlen)

	for i := amt; i < len(a); i++ {
		out[i-amt] = a[i]
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

// a(n) = Sum_{j=k..n-k} (-1)^j*n!/(k!*j!)
func Recontres(seqlen, k int64) []*bint {
	a := iSlice(seqlen)
	kb := inew(k)

	for n := int64(0); n < seqlen; n++ {
		nb := inew(n)
		sum := fzero()
		for j := int64(0); j <= n; j++ {
			jb := inew(j)
			numer := mulall(pow(inew(-1), inew(j-k)), nCr(jb, kb), fact(nb))
			sum = fadd(sum, fdiv(itof(numer), itof(fact(jb))))
		}
		fmt.Println(n, sum)
		a[n] = round(sum)
	}
	return a
}
