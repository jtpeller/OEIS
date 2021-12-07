// ============================================================================
// = A000002.go																  =
// = 	Description: Kolakoski sequence: a(n) is length of n-th run; a(1) = 1 =
// = 		sequence consists just of 1's and 2's. 							  =
// = 	Date: October 08, 2021												  =
// = 	Link: https://oeis.org/A000002										  =
// ============================================================================

package seq


func A000002(len int64, numcount int64) ([]int64, int64) {
	return Kolakoski(len, numcount), 1
}

func Kolakoski(len int64, numcount int64) []int64 {
	// declarations n stuff
	a := make([]int64, len)
	nums := make([]int64, numcount)
	for i := int64(1); i <= numcount; i++ {
		nums[i - 1] = i
	}

	// special cases
	a[0] = nums[0]
	for i := int64(0); i < nums[0]; i++ {
		a[i+1] = nums[1]
	}

	// compute other values
	numidx := int64(1)
	count := int64(1)
	for i, j := 1 + nums[0], a[count]; i < len; i, j = i+1, j-1 {
		// checks for when to reset counters/indexes
		if j <= 0 {
			i--
			count++
			j = a[count] + 1
			numidx++
			continue
		}
		if numidx >= numcount {
			numidx = 0
		}

		// assign the value
		a[i - 1] = nums[numidx]
	}
	return a
}
