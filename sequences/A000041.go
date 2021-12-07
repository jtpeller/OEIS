// ============================================================================
// = A000041.go																  =
// = 	Description: a(n) is the number of partitions of n (the partition 	  =
// = 		numbers).   			 							  	  		  =
// = 	Date: October 08, 2021												  =
// = 	Link: https://oeis.org/A000041										  =
// ============================================================================

package seq

// having a value greater than 50 will take some time
func A000041(seqlen int64) ([]int64, int64) {
	a := make([]int64, 0)
	for i := int64(0); i < seqlen; i++ {
		a = append(a, CountParts(i))
	}
	return a, 1
}

// counts the partitions of a given integer n
func CountParts(n int64) int64 {
	if n == 0 {
		return 1
	}

	// initializations
	parts := make([]int64, n)
	k := 0
	parts[k] = n

	// loop to generate partitions
	partcount := int64(0)
	for {
		// update partition count
		partcount++

		// find rightmost non-one value
		remval := int64(0)			// holds how much val can be changed
		for k >= 0 && parts[k] == 1 {
			remval += parts[k]
			k--
		}

		// if k < 0, all vals = 1; no more partitions
		if k < 0 {
			return partcount
		}
		
		// update values
		parts[k]--
		remval++

		// resort array. modify remval based on sort
		for remval > parts[k] {
			parts[k+1] = parts[k]
			remval -= parts[k]
			k++
		}

		// copy remval to next position & increment k
		parts[k+1] = remval
		k++
	}
}
