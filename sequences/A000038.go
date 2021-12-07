// ============================================================================
// = A000038.go																  =
// = 	Description: Twice A000007. 										  =
// = 	Date: October 09, 2021												  =
// = 	Link: https://oeis.org/A000038										  =
// ============================================================================

package seq

// Instead of making the sequence in A000007, this just makes a different array
// with the first value set at 2
func A000038(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	a[0] = 2
	return a, 0
}