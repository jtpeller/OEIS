// ============================================================================
// = A000012.go																  =
// = 	Description: the ones sequence			 							  =
// = 	Date: October 08, 2021												  =
// = 	Link: https://oeis.org/A000012										  =
// ============================================================================

package seq

// all 1s
func A000012(seqlen int64) ([]int64, int64) {
	a := make([]int64, 0)
	for i := int64(0); i < seqlen; i++ {
		a = append(a, 1)
	}
	return a, 0
}