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
 * Date		2025.01.26
 * Link		https://oeis.org/A000301
 */
 func A000301(seqlen int64) ([]*big.Int, int64) {
	fib, _ := A000045(seqlen)
	a := iSlice(seqlen)

	// compute a
	for n := int64(0); n < seqlen; n++ {
		a[n] = pow(inew(2), fib[n])
	}

	return a, 0
}

/**
 * A000302 Powers of 4: a(n) = 4^n.
 * Date		2025.01.26
 * Link		https://oeis.org/A000302
 */
 func A000302(seqlen int64) ([]*big.Int, int64) {
	a := iSlice(seqlen)

	// compute a
	for n := int64(0); n < seqlen; n++ {
		a[n] = pow(inew(4), inew(n))
	}

	return a, 0
}

/**
 * A000304 a(n) = a(n-1)*a(n-2)
 * Date		2025.01.27
 * Link		https://oeis.org/A000304
 */
 func A000304(seqlen int64) ([]*big.Int, int64) {
	a := iSlice(seqlen)
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
 * Date		2025.01.27
 * Link		https://oeis.org/A000308
 */
 func A000308(seqlen int64) ([]*big.Int, int64) {
	a := iSlice(seqlen)
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
 * Date		2025.01.27
 * Link		https://oeis.org/A000309
 */
 func A000309(seqlen int64) ([]*big.Int, int64) {
	a := iSlice(seqlen)
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
 * Date		2025.01.27
 * Link		https://oeis.org/A000312
 */
 func A000312(seqlen int64) ([]*big.Int, int64) {
	a := iSlice(seqlen)
	
	// compute a
	for n := int64(0); n < seqlen; n++ {
		a[n] = pow(inew(n), inew(n))
	}

	return a, 0
}

/**
 * A000313 Number of permutations of length n with 3 consecutive ascending pairs. 
 * Date		2025.01.27
 * Link		https://oeis.org/A000313
 */
 func A000313(seqlen int64) ([]*big.Int, int64) {
	a := iSlice(seqlen)
	f, _ := A000142(seqlen+2)
	
	// compute a
	for n := int64(0); n < seqlen; n++ {
		left := fdiv(fmul(itof(inew(n)), itof(f[n+1])), fnew(6))
		sum := fzero()
		for k := int64(0); k <= n; k++ {
			p := fpow(fnew(-1), k)
			foo := fdiv(p, itof(f[k]))
			sum = fadd(sum, foo)
		}
		a[n] = round(fmul(left, sum))
	}

	return a, 1
}

/**
 * A000317 Number of permutations of length n with 3 consecutive ascending pairs. 
 * Date		2025.01.27
 * Link		https://oeis.org/A000317
 */
 func A000317(seqlen int64) ([]*big.Int, int64) {
	a := iSlice(seqlen+1)
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
 * Date		2025.01.27
 * Link		https://oeis.org/A000318
 */
 func A000318(seqlen int64) ([]*big.Int, int64) {
	a := iSlice(seqlen)
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
 * Date		2025.01.27
 * Link		https://oeis.org/A000319
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

/**
 * A000321 H_n(-1/2), where H_n(x) is Hermite polynomial of degree n.
 * Date		2025.01.30
 * Link		https://oeis.org/A000321
 */
 func A000321(seqlen int64) ([]*big.Int, int64) {
	// init
	a := iSlice(seqlen)
	a[0] = inew(1)
	a[1] = inew(-1)

	// compute a
	for n := int64(2); n < seqlen; n++ {
		t1 := sub(zero(), a[n-1])
		t2:= mul( mul(inew(2), inew(n-1)), a[n-2])
		a[n] = sub(t1, t2)
	}

	return a, 0
}

/**
 * A000322 Pentanacci numbers: a(n) = a(n-1) + a(n-2) + a(n-3) + a(n-4) + a(n-5) with a(0) = a(1) = a(2) = a(3) = a(4) = 1.
 * Date		2025.01.30
 * Link		https://oeis.org/A000322
 */
 func A000322(seqlen int64) ([]*big.Int, int64) {
	// init
	a := iSlice(seqlen)
	a[0] = inew(1)
	a[1] = inew(1)
	a[2] = inew(1)
	a[3] = inew(1)
	a[4] = inew(1)

	// compute a
	for n := int64(5); n < seqlen; n++ {
		// TODO: convert to addall
		a[n] = add(add(add(add(a[n-1], a[n-2]), a[n-3]), a[n-4]), a[n-5])
	}

	return a, 0
}

/**
 * A000324 A nonlinear recurrence: a(0) = 1, a(1) = 5, a(n) = a(n-1)^2 - 4*a(n-1) + 4 for n>1.
 * Date		2025.01.30
 * Link		https://oeis.org/A000324
 */
 func A000324(seqlen int64) ([]*big.Int, int64) {
	// init
	a := iSlice(seqlen)
	a[0] = inew(1)
	a[1] = inew(5)

	// compute a
	for n := int64(2); n < seqlen; n++ {
		a[n] = add(sub(pow(a[n-1], inew(2)), mul(inew(4), a[n-1])), inew(4))
	}

	return a, 0
}

/**
 * A000325 a(n) = 2^n - n
 * Date		2025.02.08
 * Link		https://oeis.org/A000325
 */
 func A000325(seqlen int64) ([]*big.Int, int64) {
	a := iSlice(seqlen)

	for n := int64(0); n < seqlen; n++ {
		a[n] = sub(pow(inew(2), inew(n)), inew(n))
	}

	return a, 0
}

/**
 * A000326 Pentagonal numbers: a(n) = n*(3*n-1)/2
 * Date		2025.02.08
 * Link		https://oeis.org/A000326
 */
 func A000326(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)

	for n := int64(0); n < seqlen; n++ {
		a[n] = n*(3*n-1)/2
	}

	return a, 0
}

/**
 * A000327 Number of partitions into non-integral powers. 
 * Date		2025.02.08
 * Link		https://oeis.org/A000327
 */
 func A000327(seqlen int64) ([]int64, int64) {
	offset := int64(3)
	a := make([]int64, seqlen+offset)
	a148, a148_off := A000148(seqlen+offset)

	for n := int64(3); n < seqlen+offset; n++ {
		a[n-offset] = a148[n-a148_off] - int64(math.Floor( math.Pow( float64(n)/2.0, 3.0/2.0 ) ))
	}

	return a, 0
}

/**
 * A000328: Number of points of norm <= n^2 in square lattice.
 * Date		2025.02.08
 * Link		https://oeis.org/A000328
 */
 func A000328(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)

	for n := int64(0); n < seqlen; n++ {
		sum := int64(0)

		nsqr := math.Pow(float64(n), 2)
		for j := int64(0); j <= int64(nsqr) / 4; j++ {
			fj := float64(j)

			f1 := int64(nsqr / (4.0*fj + 1.0))
			f2 := int64(nsqr / (4.0*fj + 3.0))

			sum += f1 - f2
		}

		a[n] = 1 + 4 * sum
	}

	return a, 0
}

/**
 * A000329: Nearest integer to b(n), where b(n) = tan(b(n-1)), b(0) = 1. 
 * Date		2025.02.08
 * Link		https://oeis.org/A000329
 */
 func A000329(seqlen int64) ([]int64, int64) {
	utils.AccuracyWarning("A000329")

	a := make([]int64, seqlen)
	b := make([]float64, seqlen)
	a[0], b[0] = 1, 1

	for n := int64(1); n < seqlen; n++ {
		b[n] = math.Tan(b[n-1])
		a[n] = int64(math.Round(b[n]))
	}

	return a, 0
}

/**
 * A000330: Square pyramidal numbers: a(n) = 0^2 + 1^2 + 2^2 + ... + n^2 = n*(n+1)*(2*n+1)/6. 
 * Date		2025.02.08
 * Link		https://oeis.org/A000330
 */
 func A000330(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)

	for n := int64(0); n < seqlen; n++ {
		a[n] = n * (n + 1) * (2*n + 1) / 6
	}

	return a, 0
}

/**
 * A000332: Binomial coefficient binomial(n,4) = n*(n-1)*(n-2)*(n-3)/24. 
 * Date		2025.02.08
 * Link		https://oeis.org/A000332
 */
 func A000332(seqlen int64) ([]*big.Int, int64) {
	a := iSlice(seqlen)

	for n := int64(0); n < seqlen; n++ {
		a[n] = nCr(inew(n), inew(4))
	}

	return a, 0
}

/**
 * A000336: a(n) = a(n-1)*a(n-2)*a(n-3)*a(n-4); for n < 5, a(n) = n. 
 * Date		2025.02.08
 * Link		https://oeis.org/A000336
 */
 func A000336(seqlen int64) ([]*big.Int, int64) {
	a := iSlice(seqlen)

	for n := int64(0); n < seqlen; n++ {
		if n < 5 {
			a[n] = inew(n)
		} else {
			a[n] = mulall(a[n-1], a[n-2], a[n-3], a[n-4])
		}
	}

	return a, 0
}

/**
 * A000337: a(n) = a(n-1)*a(n-2)*a(n-3)*a(n-4); for n < 5, a(n) = n. 
 * Date		2025.02.08
 * Link		https://oeis.org/A000337
 */
 func A000337(seqlen int64) ([]*big.Int, int64) {
	a := iSlice(seqlen)

	for n := int64(0); n < seqlen; n++ {
		twon := pow(inew(2), inew(n))
		n1 := sub(inew(n), inew(1))
		a[n] = add(mul(n1, twon), inew(1))
	}

	return a, 0
}
