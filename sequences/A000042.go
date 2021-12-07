// ============================================================================
// = A000042.go																  =
// = 	Description: Unary representation of natural numbers.	  	  		  =
// = 	Date: October 09, 2021												  =
// = 	Link: https://oeis.org/A000042										  =
// ============================================================================

package seq

import "math"

func A000042(seqlen int64) ([]int64, int64) {
	a := make([]int64, 0)
	for i := int64(1); i < seqlen; i++ {
		temp := int64(1)
		for j := int64(1); j < i; j++ {
			temp = temp + int64(math.Pow(10, float64(j)))
		}
		a = append(a, temp)
	}
	return a, 1
}