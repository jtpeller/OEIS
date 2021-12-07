// ============================================================================
// = A000006.go																  =
// = 	Description: Integer part of square root of n-th prime.  		 	  =
// = 	Date: October 08, 2021												  =
// = 	Link: https://oeis.org/A000006										  =
// ============================================================================

package seq

import (
	util "OEIS/utils"
)

// isqrt(nth_prime)
func A000006(seqlen int64) ([]int64, int64) {
	primes, err := util.Primes(seqlen)
	util.CheckError(err)
	a := isqrtarray(primes)
	return a, 1
}

func isqrtarray(arr []int64) []int64 {
	a := make([]int64, 0)
	for i := 0; i < len(arr); i++ {
		val, err := util.Isqrt(arr[i])
		util.HandleError(err)
		a = append(a, int64(val))
	}
	return a
}