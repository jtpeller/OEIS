// ============================================================================
// = otherseq.go															  =
// = 	Description: Any sequences that don't yet have a file				  =
// = 	Date: October 08, 2021												  =
// = 	Last Update: December 12, 2021										  =
// ============================================================================

package seq

import (
	"OEIS/utils"
	"fmt"
	"math"
)

/**
 * A001065 computes the sum of proper divisors (or aliquot parts)
 *  of n: sum of divisors of n that are less than n.
 * Date		December 15, 2021
 * Link		https://oeis.org/A001065
 */
func A001065(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	for n := int64(1); n <= seqlen; n++ {
		f := utils.Factors(n)
		a[n-1] = utils.Sum(f[:len(f)-1])
	}
	return a, 1
}

/**
 * A001223 computes the prime gaps: differences b/w consecutive primes
 * Date		December 15, 2021
 * Link		https://oeis.org/A001223
 */
func A001223(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	a40, _ := A000040(seqlen + 1)
	for n := int64(0); n < seqlen; n++ {
		a[n] = a40[n+1] - a40[n]
	}
	return a, 1
}

/**
 * A001611 a(n) = Fibonacci(n) + 1.
 * Date		December 15, 2021
 * Link		https://oeis.org/A001622
 */
func A001611(seqlen int64) ([]*bint, int64) {
	a := iSlice(seqlen)
	fib, _ := A000045(seqlen)
	for n := int64(0); n < seqlen; n++ {
		a[n] = add(fib[n], inew(1))
	}

	return a, 1
}

/**
 * A001622 computes the decimal expansion of the golden ratio phi (or tau)
 *  = (1+sqrt(5))/2
 * Date		December 15, 2021
 * Link		https://oeis.org/A001622
 */
func A001622(seqlen int64) ([]int64, int64) {
	if seqlen > 1 {
		utils.PrintWarning("note: A001622 computes the decimal expansion of the golden ratio.\nThis will not be a sequence, just the value itself")
	}
	fmt.Println(fdiv(fadd(fnew(1), fsqrt(fnew(5))), fnew(2)))
	return nil, 1
}

/**
 * A001840 computes the expansion of x /((1 - x)^2 * (1 - x^3)).
 * Date		December 16, 2021
 * Link		https://oeis.org/A001840
 */
func A001840(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	for n := int64(0); n < seqlen; n++ {
		a[n] = int64(math.Floor(float64((n + 1)) * float64(n+2) / 6.0))
	}
	return a, 0
}

/**
 * A002061 computes the central polygonal #s: a(n) = n^2 - n + 1
 * Date		December 16, 2021
 * Link		https://oeis.org/A002061
 */
func A002061(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	for n := int64(0); n < seqlen; n++ {
		a[n] = int64(math.Pow(float64(n), 2.0)) - n + 1
	}
	return a, 0
}

/**
 * A002386 computes the record gaps b/w primes (lower bound)
 * Date		December 16, 2021
 * Link		https://oeis.org/A002386
 */
func A002386(seqlen int64) ([]int64, int64) {
	utils.LongCalculationWarning("A002386")

	// init
	getPrimeCount := func(n int64) int64 {
		return int64(1.5 * math.Exp(0.65*float64(n)))
	}
	a := make([]int64, seqlen)
	primes := utils.Primes(getPrimeCount(seqlen))
	a[0] = 2

	// loop
	gap := int64(0)
	prev := int64(1)
	pidx := 0
	for i := int64(1); i < seqlen; {
		gap = primes[pidx+1] - primes[pidx]
		if gap > prev {
			a[i] = primes[pidx]
			i++
			prev = gap
		}
		pidx++
	}
	return a, 1
}

/**
 * A003048 computes a[n+1]=n*a[n] - (-1)^n
 * Date		December 10, 2021	Confirmed working: December 10, 2021
 * Link		https://oeis.org/A003048
 */
func A003048(seqlen int64) ([]*bint, int64) {
	a := iSlice(seqlen)
	a[0] = inew(1)
	for i := int64(1); i < seqlen; i++ {
		a[i] = sub(mul(inew(i), a[i-1]), pow(inew(-1), inew(i)))
	}
	return a, 0
}

/**
 * A007947 computes the largest squarefree number dividing n: the
 *  squarefree kernel of n, rad(n), radical of n.
 * Date		December 16, 2021
 * Link		https://oeis.org/A007947
 */
func A007947(max int64) ([]int64, int64) {
	a := make([]int64, 0)
	for i := int64(0); i < max; i++ {
		// calculate the prime factorization of i
		pfact := utils.PrimeFactorization(i)

		// strip all non unique elements
		set := make(map[int64]bool)
		for j := int64(0); j < int64(len(pfact)); j++ {
			set[int64(pfact[j])] = true
		}

		// get the radical
		radical := int64(1)
		for key := range set {
			radical *= key
		}
		a = append(a, radical)
	}
	return a, 1
}

/**
 * A011848 computes a(n) = floor(binomial(n,2)/2)
 * Date		December 16, 2021
 * Link		https://oeis.org/A011848
 */
func A011848(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	for n := int64(2); n < seqlen; n++ {
		a[n] = int64(math.Floor(float64(utils.Binomial(n, 2)) / 2.0))
	}
	return a, 0
}

/**
 * A011858 computes a(n) = floor(n*(n-1)/5)
 * Date		December 16, 2021
 * Link		https://oeis.org/A011858
 */
func A011858(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	for n := int64(0); n < seqlen; n++ {
		a[n] = int64(math.Floor(float64(n*(n-1)) / 5.0))
	}
	return a, 0
}

/**
 * A027641 computes the numerator of Bernoulli number B_n
 * Date		December 12, 2021	Confirmed working: December 12, 2021
 * Link		https://oeis.org/A027641
 */
func A027641(seqlen int64) ([]*bint, int64) {
	a := iSlice(seqlen)
	for i := int64(0); i < seqlen; i++ {
		a[i] = utils.Bernoulli(i).Num()
	}
	return a, 0
}

/**
 * A027642 computes the denominator of Bernoulli number B_n
 * Date		December 12, 2021	Confirmed working: December
 * Link		https://oeis.org/A027642
 */
func A027642(seqlen int64) ([]*bint, int64) {
	a := iSlice(seqlen)
	for i := int64(0); i < seqlen; i++ {
		a[i] = utils.Bernoulli(i).Denom()
	}
	return a, 0
}

/**
 * A032346 essentially shifts 1 place right under inverse binomial transform.
 *  essentially the same as A000108, except row starts at 1 instead of 0
 * Date		December 07, 2021
 * Link		https://oeis.org/A032346
 */
func A032346(seqlen int64) ([]int64, int64) {
	// init
	a := make([]int64, seqlen) // the seq
	a[0] = 1
	old := make([]int64, seqlen) // last row
	new := make([]int64, seqlen) // new row
	old[0] = 1

	// compute each row & store into a
	row, col := int64(1), int64(0)
	for ; row < seqlen; row++ {
		col = 0

		// calculate new row
		for ; col < row; col++ {
			new[col+1] = new[col] + old[col]
		}

		// copy down
		if col > 0 {
			for i := int64(0); i < col+1; i++ {
				old[i] = new[i]
				new[i] = 0 // erase new row
			}
		}

		// copy the last element
		new[0] = old[col] // overwrite first elem
		a[row] = old[col] // copy last elem of old
	}
	return a, 0
}

/**
 * A038040 computes a(n) = n*d(n), where d(n) = number of divisors of n (A000005).
 * Date		December 16, 2021
 * Link		https://oeis.org/A038040
 */
func A038040(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	d, _ := A000005(seqlen)
	for n := int64(0); n < seqlen; n++ {
		a[n] = (n + 1) * d[n]
	}
	return a, 1
}

/**
 * A052614 computes E.g.f. 1/((1-x)(1-x^4)).
 * Date		December 16, 2021
 * Link		https://oeis.org/A052614
 */
func A052614(seqlen int64) ([]*bint, int64) {
	a := iSlice(seqlen)
	for n := int64(0); n < seqlen; n++ {
		nf := float64(n)
		sum := fnew(0)
		for k := float64(0); k <= nf/4.0; k++ {
			sum = fadd(sum, fnew(math.Exp(-1.0/4.0)))
		}
		a[n] = floor(fmul(itof(fact(inew(n))), sum))
	}
	return a, 0
}

/**
 * A088218 computes the total number of leaves in all rooted ordered trees with n edges.
 * Date		December 07, 2021
 * Link		https://oeis.org/A088218
 */
func A088218(seqlen int64) ([]*bint, int64) {
	a := iSlice(seqlen)
	for i := int64(0); i < seqlen; i++ {
		a[i] = div(fact(inew(2*i)), mul(fact(inew(i)), fact(inew(i+1))))
	}
	return a, 0
}

/**
 * A128422 computes projective plane crossing number of K_{4,n}.
 * Date		December 16, 2021
 * Link		https://oeis.org/A128422
 */
func A128422(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	for n := int64(0); n < seqlen; n++ {
		a[n] = n * (n - 1) / 3
	}
	return a, 1
}

/**
 * A132269 computes Product{k>=0, 1+floor(n/2^k)}.
 * Date		December 16, 2021
 * Link		https://oeis.org/A132269
 */
func A132269(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	a[0] = 1
	for n := int64(1); n < seqlen; n++ {
		for k := int64(0); k <= n; k++ {
			a[n] += a[n/2]
		}
	}
	return a, 1
}

/**
 * A164514 computes 1 followed by numbers that are not squares
 * Date		December 16, 2021
 * Link		https://oeis.org/A164514
 */
func A164514(seqlen int64) ([]int64, int64) {
	a37, _ := A000037(seqlen)
	a := utils.Shift(a37, 1)
	a[0] = 1
	return a, 1
}

/**
 * A168014 computes the sum of all parts of all partitions of n
 *  into equal parts that do not contain 1 as a part.
 * Date		December 16, 2021
 * Link		https://oeis.org/A168014
 */
func A168014(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	a5, _ := A000005(seqlen)
	for n := int64(1); n < seqlen; n++ {
		a[n] = n * (a5[n-1] - 1)
	}
	return a, 0
}
