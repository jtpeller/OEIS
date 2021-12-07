// ============================================================================
// = A000040.go																  =
// = 	Description: The prime numbers 										  =
// = 	Date: October 09, 2021												  =
// = 	Link: https://oeis.org/A000040										  =
// ============================================================================

package seq

import (
	util "OEIS/utils"
)

// Computes prime numbers using Golang's built-in (see utils.go)
func A000040(seqlen int64) ([]int64, int64) {
	a := make([]int64, 0)
	for i := int64(0); i < seqlen; i++ {
		if util.IsPrime(i) {
			a = append(a, i)
		}
	}
	return a, 1
}