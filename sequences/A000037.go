// ============================================================================
// = A000037.go																  =
// = 	Description: Numbers that are not squares (or, the nonsquares). 	  =
// = 	Date: October 09, 2021												  =
// = 	Link: https://oeis.org/A000037										  =
// ============================================================================

package seq

import "math"

// computes nonsquares
func A000037(seqlen int64) ([]int64, int64) {
	a := make([]int64, 0)
	for i := int64(1); i < seqlen; i++ {
		a = append(a, i + int64(math.Floor(0.5 + math.Sqrt(float64(i)))))
	}
	return a, 1
}