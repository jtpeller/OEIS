// ============================================================================
// = check.go																  =
// = 	Description: Useful checker functions like isPrime.					  =
// = 	Date: December 07, 2021												  =
// ============================================================================

package utils

import (
	"math"
	"math/big"
)

// ############################### CHECKERS ###################################
// ### this section checks if a number has a specific property

// IsPrime returns true if num is prime. False otherwise.
func IsPrime(num int64) bool {
	return big.NewInt(num).ProbablyPrime(20)
}

// IsBigPrime returns true if num is prime. False otherwise.
func IsBigPrime(num *big.Int) bool {
	return num.ProbablyPrime(20)
}

// Checks if the given number n is a prime power of k
func IsPrimePower(n int64, k int64) bool {
	nf := float64(n)
	kf := float64(k)
	return (math.Log(nf) / math.Log(kf) == math.Floor(math.Log(nf) / math.Log(kf)))
}

// IsSquare returns true if the num is a perfect square. False otherwise.
func IsSquare(n int64) bool {
	nf := float64(n)
	sr := math.Sqrt(nf)
	return ((sr - math.Floor(sr)) == 0)
}