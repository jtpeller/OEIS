// ============================================================================
// = A000005.go																  =
// = 	Description: d(n) (also called tau(n) or sigma_0(n)), the number of   =
// = 		divisors of n. 		 							  				  =
// = 	Date: October 08, 2021												  =
// = 	Link: https://oeis.org/A000005										  =
// ============================================================================

package seq

import (
	util "OEIS/utils"
)

func A000005(len int64) ([]int64, int64) {
	a := make([]int64, 0)
	for i := int64(1); i < len; i++ {
		count := util.GetFactorCount(i)
		a = append(a, count)
	}
	return a, 1
}