// ============================================================================
// = A000008.go																  =
// = 	Description: Number of ways of making change for n cents using coins  =
// = 		of 1, 2, 5, 10 cents. 			 							  	  =
// = 	Date: October 08, 2021												  =
// = 	Link: https://oeis.org/A000008										  =
// ============================================================================

package seq

func A000008(seqlen int64) ([]int64, int64) {
	denoms := []int64{1, 2, 5, 10}
	a := make([]int64, 0)
	coins := int64(len(denoms))

	for i := int64(0); i < seqlen; i++ {
		a = append(a, makeChange(coins, i, denoms))
	}
	return a, 0
}

func makeChange(len int64, val int64, denom []int64) int64 {
	if val < 0 {
		return 0
	} else if val == 0 {
		return 1			// 1 way to make 0 change
	} else if len <= 0 && val >= 1 {
		return 0			// combo doesn't work
	}
	return makeChange(len - 1, val, denom) + makeChange(len, val - denom[len - 1], denom)
}
