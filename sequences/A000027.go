// ============================================================================
// = A000027.go																  =
// = 	Description: The positive integers. Also called the natural numbers,  =
// = 		the whole numbers or the counting numbers, but these terms are	  =
// = 		ambiguous. 		 						  						  =
// = 	Date: October 09, 2021												  =
// = 	Link: https://oeis.org/A000027										  =
// ============================================================================

package seq

func A000027(seqlen int64) ([]int64, int64) {
	a := make([]int64, 0)
	for i := int64(0); i < seqlen + 1; i++ {
		a = append(a, i+1)
	}
	return a, 1
}