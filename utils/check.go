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
	return big.NewInt(num).ProbablyPrime(0)
}

// Checks if the given number n is a prime power of k
func IsPrimePower(n int64, k int64) bool {
	nf := float64(n)
	kf := float64(k)
	return (math.Log(nf) / math.Log(kf) == math.Floor(math.Log(nf) / math.Log(kf)))
}