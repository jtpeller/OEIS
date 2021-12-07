// ============================================================================
// = A000120.go																  =
// = 	Description: 1's-counting sequence: number of 1's in binary expansion =
// = 			of n (or the binary weight of n).   						  =
// = 	Date: December 07, 2021												  =
// = 	Link: https://oeis.org/A000120										  =
// ============================================================================

package seq

import (
	"strconv"
)

func A000120(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	for i := int64(0); i < seqlen; i++ {
		binary := strconv.FormatInt(i, 2)	// convert to binary
		count := int64(0)
		for _, bit := range binary {
			if bit == '1' {
				count++
			}
		}
		a[i] = count
	}
	return a, 1
}
