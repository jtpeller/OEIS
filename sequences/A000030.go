// ============================================================================
// = A000030.go																  =
// = 	Description: Initial digit of n 		 						  	  =
// = 	Date: October 09, 2021												  =
// = 	Link: https://oeis.org/A000030										  =
// ============================================================================

package seq

import (
	"OEIS/utils"
	"math"
)

// sequence of the first digit of n
func A000030(seqlen int64) ([]int64, int64) {
	a := make([]int64, 0)
	for i := int64(0); i < seqlen; i++ {
		a = append(a, GetFirstDigit(i))
	}
	return a, 0
}

// finds the first digit of the number
func GetFirstDigit(n int64) int64 {
	return int64(float64(n) / math.Pow(10, float64(utils.GetDigits(n) - 1)))
}