// ============================================================================
// = thru400.go
// = 	Description		OEIS sequences from A000301-A000400
// = 	Note			Not all sequences in this range have been programmed
// = 	Date 			2025.01.26
// ============================================================================

package seq

import (
	"OEIS/utils"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

/**
 * A000301 a(n) = a(n-1)*a(n-2) with a(0) = 1, a(1) = 2; also a(n) = 2^Fibonacci(n)
 * Date		2025.01.26
 * Link		https://oeis.org/A000301
 */
 func A000301(seqlen int64) ([]*bint, int64) {
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
 func A000302(seqlen int64) ([]*bint, int64) {
	a := utils.Powers(seqlen, inew(4))
	return a, 0
}

/**
 * A000304 a(n) = a(n-1)*a(n-2)
 * Date		2025.01.27
 * Link		https://oeis.org/A000304
 */
 func A000304(seqlen int64) ([]*bint, int64) {
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
 func A000308(seqlen int64) ([]*bint, int64) {
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
 func A000309(seqlen int64) ([]*bint, int64) {
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
 func A000312(seqlen int64) ([]*bint, int64) {
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
 func A000313(seqlen int64) ([]*bint, int64) {
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
 func A000317(seqlen int64) ([]*bint, int64) {
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
 func A000318(seqlen int64) ([]*bint, int64) {
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
 func A000321(seqlen int64) ([]*bint, int64) {
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
 func A000322(seqlen int64) ([]*bint, int64) {
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
 func A000324(seqlen int64) ([]*bint, int64) {
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
 func A000325(seqlen int64) ([]*bint, int64) {
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
 func A000332(seqlen int64) ([]*bint, int64) {
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
 func A000336(seqlen int64) ([]*bint, int64) {
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
 func A000337(seqlen int64) ([]*bint, int64) {
	a := iSlice(seqlen)

	for n := int64(0); n < seqlen; n++ {
		twon := pow(inew(2), inew(n))
		n1 := sub(inew(n), inew(1))
		a[n] = add(mul(n1, twon), inew(1))
	}

	return a, 0
}

/**
 * A000339: Number of partitions into non-integral powers. 
 * Date		2025.02.09
 * Link		https://oeis.org/A000339
 */
 func A000339(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	offset := int64(2)

	for n := offset; n < seqlen+offset; n++ {
		fn := float64(n)
		nsqr := int64(math.Pow(fn, 2))
		sum := int64(0)
		for x1 := int64(1); x1 <= nsqr; x1++ {
			x1sqrt := math.Sqrt(float64(x1))
			x2 := int64(math.Pow(fn - x1sqrt, 2.0))
			if (x2 >= x1) {
				sum += int64(x2-x1+1)
			}
		}
		a[n-offset] = sum
	}

	return a, offset
}

/**
 * A000340: a(0)=1, a(n) = 3*a(n-1) + n + 1 
 * Date		2025.02.09
 * Link		https://oeis.org/A000340
 */
 func A000340(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	a[0] = 1

	for n := int64(1); n < seqlen; n++ {
		a[n] = 3 * a[n-1] + n + 1
	}

	return a, 0
}

/**
 * A000344: a(n) = 5*binomial(2n, n-2)/(n+3). 
 * Date		2025.02.09
 * Link		https://oeis.org/A000344
 */
 func A000344(seqlen int64) ([]*bint, int64) {
	a := iSlice(seqlen)
	offset := int64(2)

	for n := offset; n < seqlen+offset; n++ {
		nb := inew(n)
		two := inew(2)
		binom := nCr(mul(two, nb), sub(nb, two))
		a[n-offset] = div(mul(inew(5), binom), add(nb, inew(3)))
	}

	return a, offset
}

/**
 * A000346: a(n) = 2^(2*n+1) - binomial(2*n+1, n+1).
 * Date		2025.02.09
 * Link		https://oeis.org/A000346
 */
 func A000346(seqlen int64) ([]*bint, int64) {
	a := iSlice(seqlen)

	for n := int64(0); n < seqlen; n++ {
		nb := inew(n)
		twonplus1 := add(mul(inew(2), nb), inew(1))
		a[n] = sub(pow(inew(2), twonplus1), nCr(twonplus1, add(nb, inew(1))))
	}

	return a, 0
}

/**
 * A000350: Numbers m such that Fibonacci(m) ends with m. 
 * Date		2025.02.09
 * Link		https://oeis.org/A000350
 */
 func A000350(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	fib, _ := A000045(seqlen*seqlen)

	n := int64(0)
	for m := int64(0); n < seqlen; m++ {
		if m >= int64(len(fib)) {
			fib = append(fib, add(fib[m-2], fib[m-1]))
		}
		str := fib[m].String()
		m_str := strconv.FormatInt(m, 10)

		// only save to a[n] if the fib(m) ends with m
		if strings.HasSuffix(str, m_str) {
			a[n] = m
			n++
		}
	}

	return a, 1
}

/**
 * A000351: Powers of 5: a(n) = 5^n. 
 * Date		2025.02.09
 * Link		https://oeis.org/A000351
 */
 func A000351(seqlen int64) ([]*bint, int64) {
	a := utils.Powers(seqlen, inew(5))
	return a, 0
}

/**
 * A000352: One half of the number of permutations of [n] such 
	that the differences have three runs with the same signs.
 * Date		2025.02.09
 * Link		https://oeis.org/A000352
 */
 func A000352(seqlen int64) ([]*bint, int64) {
	a := iSlice(seqlen)
	offset := int64(4)

	for n := offset; n < seqlen+offset; n++ {
		nb := inew(n)
		three_n := pow(inew(3), nb)					// 3^n
		fourpow := mul(inew(4), pow(inew(2), nb))	// 4*2^n
		twon := mul(inew(2), nb)					// 2*n
		a[n-offset] = div(add(suball(three_n, fourpow, twon), inew(11)), inew(4))
	}

	return a, offset
}

/**
 * A000353: Primes p == 7, 19, 23 (mod 40) such that (p-1)/2 is also prime. 
 * Date		2025.02.09
 * Link		https://oeis.org/A000353
 */
 func A000353(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	primes := []int64{7, 19, 23}
	const MODVAL = 40

	n := int64(0)
	for p := int64(1); n < seqlen; p++ {
		if utils.IsPrime(p) && slices.Contains(primes, p % int64(MODVAL)) {
			if utils.IsPrime((p-1)/2) {
				a[n] = p
				n++
			}
		}
	}

	return a, 1
}

/**
 * A000354: Expansion of e.g.f. exp(-x)/(1-2*x).
 * Date		2025.02.09
 * Link		https://oeis.org/A000354
 */
 func A000354(seqlen int64) ([]*bint, int64) {
	a := iSlice(seqlen)

	for n := int64(0); n < seqlen; n++ {
		nb := inew(n)
		sum := inew(0)
		for k := int64(0); k <= n; k++ {
			kb := inew(k)
			pow1 := pow(inew(-1), add(nb, kb))	// (-1)^(n+k)
			binom := nCr(nb, kb)
			kfact := fact(kb)
			twok := pow(inew(2), kb)
			sum = add(sum, mulall(pow1, binom, kfact, twok))
		}
		a[n] = sum
	}

	return a, 0
}

/**
 * A000355: Primes p == 3, 9, 11 (mod 20) such that 2p+1 is also prime
 * Date		2025.02.09
 * Link		https://oeis.org/A000355
 */
 func A000355(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	primes := []int64{3, 9, 11}
	const MODVAL = 20

	n := int64(0)
	for p := int64(1); n < seqlen; p++ {
		if utils.IsPrime(p) && slices.Contains(primes, p % int64(MODVAL)) {
			if utils.IsPrime(2*p+1) {
				a[n] = p
				n++
			}
		}
	}

	return a, 1
}

/**
 * A000356: Number of rooted cubic maps with 2n nodes and a 
 *		distinguished Hamiltonian cycle: (2n)!(2n+1)! / (n!^2*(n+1)!(n+2)!)
 * Date		2025.02.09
 * Link		https://oeis.org/A000356
 */
 func A000356(seqlen int64) ([]*bint, int64) {
	a := iSlice(seqlen)
	offset := int64(1)

	for n := offset; n <= seqlen; n++ {
		twofact := fact(inew(2*n))
		twofact1 := fact(add(inew(2*n), inew(1)))
		nfactsqr := pow(fact(inew(n)), inew(2))
		nfact1 := fact(inew(n+1))
		nfact2 := fact(inew(n+2))
		numer := mul(twofact, twofact1)
		denom := mulall(nfactsqr, nfact1, nfact2)
		a[n-offset] = div(numer, denom)
	}

	return a, offset
}

/**
 * A000358: Number of binary necklaces of length n with no subsequence 00, 
 *		excluding the necklace "0". 
 * Date		2025.02.09
 * Link		https://oeis.org/A000358
 */
 func A000358(seqlen int64) ([]*bint, int64) {
	a := iSlice(seqlen)
	offset := int64(1)
	fib, _ := A000045(seqlen+offset*2)

	for n := offset; n <= seqlen; n++ {
		nb := inew(n)
		sum := inew(0)
		for d := offset; d <= n; d++ {
			if n % d == 0 {
				tot := utils.EulerTotientBig(inew(n/d))
				fibs := add(fib[d-1], fib[d+1])
				sum = add(sum, mul(tot, fibs))
			}
		}
		a[n-offset] = div(sum, nb)
	}

	return a, offset
}

/**
 * A000363: Number of permutations of [n] with exactly 2 increasing 
 *		runs of length at least 2. 
 * Date		2025.02.09
 * Link		https://oeis.org/A000363
 */
 func A000363(seqlen int64) ([]*bint, int64) {
	a := iSlice(seqlen)
	offset := int64(4)

	for n := offset; n < seqlen+offset; n++ {
		nb := inew(n)
		fiven := pow(inew(5), nb)
		twon1 := sub(mul(inew(2), nb), inew(1))
		three_n := pow(inew(3), nb)
		twonsqr := mul(inew(2), pow(nb, inew(2)))
		twon := mul(inew(2), nb)
		numer := suball(add(sub(fiven, mul(twon1, three_n)), twonsqr), twon, inew(2))
		a[n-offset] = div(numer, inew(16))
	}

	return a, offset
}

/**
 * A000371: a(n) = Sum_{k=0..n} (-1)^(n-k)*binomial(n,k)*2^(2^k).
 * Date		2025.02.09
 * Link		https://oeis.org/A000371
 */
 func A000371(seqlen int64) ([]*bint, int64) {
	a := iSlice(seqlen)
	offset := int64(0)

	for n := offset; n < seqlen+offset; n++ {
		nb := inew(n)
		sum := zero()
		for k := int64(0); k <= n; k++ {
			kb := inew(k)
			t1 := pow(inew(-1), inew(n-k))
			t2 := nCr(nb, kb)
			t3 := pow(inew(2), pow(inew(2), kb))
			sum = add(sum, mulall(t1, t2, t3))
		}
		a[n] = sum
	}

	return a, offset
}

/**
 * A000381: a(n) = Sum_{k=0..n} (-1)^(n-k)*binomial(n,k)*2^(2^k).
 * Date		2025.02.09
 * Link		https://oeis.org/A000381
 */
 func A000381(seqlen int64) ([]*bint, int64) {
	a1611, _ := A001611(seqlen+2)
	a := utils.ShiftBigSliceLeft(a1611, 2)
	return a, 0
}

/**
 * A000383: Hexanacci numbers with a(0) = ... = a(5) = 1. 
 * Date		2025.02.09
 * Link		https://oeis.org/A000383
 */
 func A000383(seqlen int64) ([]*bint, int64) {
	a := utils.Nacci(seqlen, 6, false)
	return a, 0
}

/**
 * A000384: Hexagonal numbers: a(n) = n*(2*n-1). 
 * Date		2025.02.09
 * Link		https://oeis.org/A000384
 */
 func A000384(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	for n := int64(0); n < seqlen; n++ {
		a[n] = n*(2*n-1)
	}
	return a, 0
}

/**
 * A000385: Convolution of A000203 with itself.
 * Date		2025.02.09
 * Link		https://oeis.org/A000385
 */
 func A000385(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	a203, offset203 := A000203(seqlen+1)
	offset := int64(1)
	for n := offset; n < seqlen+offset; n++ {
		sum := int64(0)
		for k := int64(1); k <= n; k++ {
			sum += a203[k-offset203] * a203[n-k+1-offset203]
		}
		a[n-offset] = sum
	}
	return a, offset
}

/**
 * A000387: Rencontres numbers: number of permutations of [n] with 
 *		exactly two fixed points. 
 * Date		2025.02.09
 * Link		https://oeis.org/A000387
 */
 func A000387(seqlen int64) ([]*bint, int64) {
	a := utils.Recontres(seqlen, 2)
	return a, 0
}

/**
 * A000389: Rencontres numbers: number of permutations of [n] with 
 *		exactly two fixed points. 
 * Date		2025.02.09
 * Link		https://oeis.org/A000389
 */
 func A000389(seqlen int64) ([]*bint, int64) {
	a := iSlice(seqlen)
	big5 := inew(5)
	for n := int64(0); n < seqlen; n++ {
		a[n] = nCr(inew(n), big5)
	}
	return a, 0
}

/**
 * A000392: Stirling numbers of second kind S(n,3). 
 * Date		2025.02.09
 * Link		https://oeis.org/A000392
 */
 func A000392(seqlen int64) ([]*bint, int64) {
	a := iSlice(seqlen)
	for n := int64(0); n < seqlen; n++ {
		a[n] = utils.Stirling2(n, 3)
	}
	return a, 0
}

/**
 * A000396: Perfect numbers k: k is equal to the sum of the proper divisors of k. 
 * Date		2025.02.09
 * Link		https://oeis.org/A000396
 */
 func A000396(seqlen int64) ([]*bint, int64) {
	utils.LongCalculationWarning("A000396")
	offset := int64(1)
	a := iSlice(seqlen)
	n := int64(0)
	for k := offset; n < seqlen; k++ {
		kb := inew(k)
		divs := utils.FactorsBig(kb, false)
		sumdivs := utils.SumBig(divs)
		if equals(sumdivs, kb) {
			a[n] = kb
			fmt.Println(a)
			n++
		}
	}
	return a, offset
}

/**
 * A000399: Unsigned Stirling numbers of first kind s(n,3). 
 * Date		2025.02.09
 * Link		https://oeis.org/A000399
 */
 func A000399(seqlen int64) ([]*bint, int64) {
	utils.AccuracyWarning("A000399")
	a := iSlice(seqlen)
	offset := int64(3)

	for n := offset; n < seqlen+offset; n++ {
		a[n-offset] = utils.Stirling1(n, offset)
	}

	return a, offset
}

/**
 * A000400: Unsigned Stirling numbers of first kind s(n,3). 
 * Date		2025.02.09
 * Link		https://oeis.org/A000399
 */
 func A000400(seqlen int64) ([]*bint, int64) {
	a := utils.Powers(seqlen, inew(6))
	return a, 0
}
