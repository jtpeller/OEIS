// ============================================================================
// = A000032.go																  =
// = 	Description: Lucas numbers beginning at 2: L(n) = L(n-1) + L(n-2), 	  =
// = 		L(0) = 2, L(1) = 1.  		 						 		 	  =
// = 	Date: October 09, 2021												  =
// = 	Link: https://oeis.org/A000032										  =
// ============================================================================

package seq

// computes the Lucas numbers
func A000032(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	a[0] = 2		// a(0)=2
	a[1] = 1		// a(1)=1	
	for i := int64(2); i < seqlen; i++ {
		a[i] = a[i - 2] + a[i - 1]
	}
	return a, 0
}