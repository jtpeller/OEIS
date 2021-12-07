// ============================================================================
// = A000010.go																  =
// = 	Description: a(n) is the number of partitions of n (the partition 	  =
// = 		numbers).   			 							  	  		  =
// = 	Date: October 08, 2021												  =
// = 	Link: https://oeis.org/A000010										  =
// ============================================================================

package seq

import (
	"OEIS/utils"
)

func A000010(seqlen int64) ([]int64, int64) {
	a := make([]int64, 0)
	for i := int64(1); i < seqlen; i++ {
		a = append(a, utils.EulerTotient(i))
	}
	return a, 1
}