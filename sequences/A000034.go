// ============================================================================
// = A000034.go																  =
// = 	Description: Period 2: repeat [1, 2]; a(n) = 1 + (n mod 2).   		  =
// = 	Date: October 09, 2021												  =
// = 	Link: https://oeis.org/A000034										  =
// ============================================================================

package seq

// a(n) = 1 + (n mod 2), or 1 + A000035(n)
func A000034(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	for i := int64(0); i < seqlen; i++ {
		a[i] = i % 2 + 1
	}
	return a, 0
}