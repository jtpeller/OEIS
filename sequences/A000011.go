// ============================================================================
// = A000011.go																  =
// = 	Description: Number of n-bead necklaces (turning over is allowed) 	  =
// = 		where complements are equivalent. 		 						  =
// = 	Date: October 08, 2021												  =
// = 	Link: https://oeis.org/A000011										  =
// ============================================================================

package seq

import (
	"OEIS/utils"
	"math"
)

// NOTE: this implementation is slightly inaccurate due to rounding errors
func A000011(seqlen int64) ([]int64, int64) {
	// generate euler phi
	eulerlen := seqlen * 2
	euler := make([]int64, 0)
	for i := int64(0); i < eulerlen; i++ {
		euler = append(euler, utils.EulerTotient(i))
	}

	// generate even sequence
	even := make([]int64, seqlen)
	for i := int64(0); i < seqlen; i++ {
		if i == 0 {
			even[i] = 1
		} else {
			divisors := utils.Factors(int64(i))
			euleridx := int64(0)

			// use the divisors to calculate the sequence
			factorcount := len(divisors)
			for j := 0; j < factorcount; j++ {
				euleridx = 2 * divisors[j] - 1
				b := math.Pow(2, float64(i) / float64(divisors[j]))
				even[i] += int64(float64(euler[euleridx]) * b)
			}
			even[i] = even[i] / (2.0 * i)
			foo := math.Pow(2, float64(i / 2))
			even[i] += int64(foo)
			even[i] = even[i] / 2
		}
	}

	return even, 0
}