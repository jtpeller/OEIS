// ============================================================================
// = thru400.go																  =
// = 	Description		All OEIS sequences from A000301-A000400				  =
// = 	Note			Not all sequences in this range have been programmed  =
// = 	Date: 			2025.01.26											  =
// = 	Last Update		2025.01.26											  =
// ============================================================================

package seq

import (
	"OEIS/utils"
	"math/big"
)

const (
	LONG_A000301 = 10
)

/**
 * A000301 a(n) = a(n-1)*a(n-2) with a(0) = 1, a(1) = 2; also a(n) = 2^Fibonacci(n)
 * Date: 2025.01.26
 * Link: https://oeis.org/A000301
 */
 func A000301(seqlen int64) ([]*big.Int, int64) {
	if seqlen > LONG_A000301 {
		utils.LongCalculationWarning("A000205", LONG_A000301)
	}

	fib, _ := A000045(seqlen)
	a := utils.CreateSlice(seqlen)

	// compute a
	for n := int64(0); n < seqlen; n++ {
		a[n] = pow(big.NewInt(2), fib[n])
	}

	return a, 0
}