// ============================================================================
// = thru400.go
// = 	Description		OEIS sequences from A000301-A000400
// = 	Note			Not all sequences in this range have been programmed
// = 	Date 			2025.01.26
// ============================================================================

package seq

import (
	"OEIS/utils"
	"math"
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
		a[n] = pow(inew(2), fib[n])
	}

	return a, 0
}

/**
 * A000302 Powers of 4: a(n) = 4^n.
 * Date: 2025.01.26
 * Link: https://oeis.org/A000302
 */
 func A000302(seqlen int64) ([]*big.Int, int64) {
	a := utils.CreateSlice(seqlen)

	// compute a
	for n := int64(0); n < seqlen; n++ {
		a[n] = pow(inew(4), inew(n))
	}

	return a, 0
}

/**
 * A000304 a(n) = a(n-1)*a(n-2)
 * Date: 2025.01.27
 * Link: https://oeis.org/A000304
 */
 func A000304(seqlen int64) ([]*big.Int, int64) {
	a := utils.CreateSlice(seqlen)
	a[0] = inew(2)
	a[1] = inew(3)

	// compute a
	for n := int64(2); n < seqlen; n++ {
		a[n] = mul(a[n-1], a[n-2])
	}

	return a[:], 0
}

/**
 * A000308 a(n) = a(n-1)*a(n-2)
 * Date: 2025.01.27
 * Link: https://oeis.org/A000308
 */
 func A000308(seqlen int64) ([]*big.Int, int64) {
	a := utils.CreateSlice(seqlen)
	a[0] = inew(1)
	a[1] = inew(2)
	a[2] = inew(3)

	// compute a
	for n := int64(3); n < seqlen; n++ {
		a[n] = mul(mul(a[n-1], a[n-2]), a[n-3])
	}

	return a, 1
}

/**
 * A000309 Number of rooted planar bridgeless cubic maps with 2n nodes.
 * Date: 2025.01.27
 * Link: https://oeis.org/A000309
 */
 func A000309(seqlen int64) ([]*big.Int, int64) {
	a := utils.CreateSlice(seqlen)
	a139, _ := A000139(seqlen)
	a[0] = inew(1)
	
	// compute a
	for n := int64(1); n < seqlen; n++ {
		a[n] = mul(pow(inew(2), inew(n-1)), a139[n])
	}

	return a, 0
}

/**
 * A000312 a(n) = n^n; number of labeled mappings from n points to themselves (endofunctions)
 * Date: 2025.01.27
 * Link: https://oeis.org/A000312
 */
 func A000312(seqlen int64) ([]*big.Int, int64) {
	a := utils.CreateSlice(seqlen)
	
	// compute a
	for n := int64(0); n < seqlen; n++ {
		a[n] = pow(inew(n), inew(n))
	}

	return a, 0
}

/**
 * A000313 Number of permutations of length n with 3 consecutive ascending pairs. 
 * Date: 2025.01.27
 * Link: https://oeis.org/A000313
 */
 func A000313(seqlen int64) ([]*big.Int, int64) {
	a := utils.CreateSlice(seqlen)
	f, _ := A000142(seqlen+2)
	
	// compute a
	for n := int64(0); n < seqlen; n++ {
		left := fdiv(fmul(tofloat(inew(n)), tofloat(f[n+1])), fnew(6))
		sum := fzero()
		for k := int64(0); k <= n; k++ {
			p := fpow(fnew(-1), k)
			foo := fdiv(p, tofloat(f[k]))
			sum = fadd(sum, foo)
		}
		a[n] = round(fmul(left, sum))
	}

	return a, 1
}

/**
 * A000317 Number of permutations of length n with 3 consecutive ascending pairs. 
 * Date: 2025.01.27
 * Link: https://oeis.org/A000317
 */
 func A000317(seqlen int64) ([]*big.Int, int64) {
	a := utils.CreateSlice(seqlen+1)
	a[0] = inew(1)
	a[1] = inew(2)
	
	// compute a
	for n := int64(1); n < seqlen; n++ {
		a[n+1] = add(sub(pow(a[n], inew(2)), mul(a[n], a[n-1])), pow(a[n-1], inew(2)))
	}

	return a, 1
}

/**
 * A000318 Number of permutations of length n with 3 consecutive ascending pairs. 
 * Date: 2025.01.27
 * Link: https://oeis.org/A000318
 */
 func A000318(seqlen int64) ([]*big.Int, int64) {
	a := utils.CreateSlice(seqlen)
	a182, _ := A000182(seqlen)
	
	// compute a
	for n := int64(1); n <= seqlen; n++ {
		// -1 due to indexing starting at zero!
		a[n-1] = mul(pow(inew(2), inew(4*n-2)), a182[n-1])
	}

	return a, 1
}

/**
 * A000319 a(n) = floor(b(n)), where b(n) = tan(b(n-1)), b(0)=1. 
 * Date: 2025.01.27
 * Link: https://oeis.org/A000319
 */
 func A000319(seqlen int64) ([]int64, int64) {
	utils.PrintWarning("Due to the implementation of tan() in Go, this sequence is inaccurate for n > 14. More precision is necessary, but tan is not available for arbitrary precision floats.")

	// first generate b
	b := make([]float64, seqlen+1)
	b[0] = 1
	for n := int64(1); n <= seqlen; n++ {
		b[n] = math.Tan(b[n-1])
	}

	// then compute a
	a := make([]int64, seqlen)
	for n := int64(0); n < seqlen; n++ {
		a[n] = int64(math.Floor(b[n]))
	}

	return a, 1
}

