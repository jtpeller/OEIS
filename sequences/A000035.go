// ============================================================================
// = A000035.go																  =
// = 	Description: Period 2: repeat [0, 1]; a(n) = n mod 2; parity of n.    =
// = 	Date: October 09, 2021												  =
// = 	Link: https://oeis.org/A000035										  =
// ============================================================================

package seq

// computes the parity of n (basically, n mod 2)
func A000035(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	for i := int64(0); i < seqlen; i++ {
		a[i] = i % 2
	}
	return a, 0
}