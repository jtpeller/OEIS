// ============================================================================
// = A007947.go																  =
// = 	Description: Largest squarefree number dividing n: the squarefree 	  =
// = 		kernel of n, rad(n), radical of n.	  	  		  				  =
// = 	Date: October 09, 2021												  =
// = 	Link: https://oeis.org/A007947										  =
// ============================================================================

package seq

import util "OEIS/utils"

func A007947(max int64) ([]int64, int64) {
	a := make([]int64, 0)
	for i := int64(0); i < max; i++ {
		// calculate the prime factorization of i
		pfact := util.PrimeFactorization(int64(i))

		// strip all non unique elements
		set := make(map[int64]bool)
		for j := int64(0); j < int64(len(pfact)); j++ {
			set[int64(pfact[j])] = true
		}

		// get the radical
		radical := int64(1)
		for key := range set {
			radical *= key
		}
		a = append(a, radical)
	}
	return a, 1
}