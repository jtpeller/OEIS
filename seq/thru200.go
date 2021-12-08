package seq

import (
	"OEIS/utils"
	"math"
	"strconv"
)

const (
	LONG_A000101 = 17
	OVERFLOW_A000102 = 68
	OVERFLOW_A000108 = 36
	OVERFLOW_A000110 = 25
	OVERFLOW_A000111 = 21
	OVERFLOW_A000116 = 32
)

/**
 * A000101 computes the record gaps b/w primes (upper end)
 * Date: December 07, 2021
 * Link: https://oeis.org/A000101
 */
func A000101(seqlen int64) ([]int64, int64) {
	// warn user of long calculation times
	if seqlen > LONG_A000101 {
		utils.LongCalculationWarning("A000101", LONG_A000101)
	}

	// init
	getPrimeCount := func(n int64) int64 {
		return int64(1.5 * math.Exp(0.65 * float64(n)))
	}
	a := make([]int64, seqlen)
	primes, _ := utils.Primes(getPrimeCount(seqlen))
	a[0] = 3

	// loop
	gap := int64(0)
	prev := int64(1)
	pidx := 0
	for i := int64(1); i < seqlen; {
		gap = primes[pidx+1] - primes[pidx]
		if gap > prev {
			a[i] = primes[pidx + 1]
			i++
			prev = gap
		}
		pidx++
	}
	return a, 1
}

/**
 * A000102 computes a(n) such that a(n) is the # of compositions of n in which the
 *  maximal part is 3. Convoltuion of tribonacci & tetranacci
 * Date: December 07, 2021
 * Link: https://oeis.org/A000102
 */
func A000102(seqlen int64) ([]int64, int64) {
	if seqlen > OVERFLOW_A000102 {
		utils.OverflowError("A000102", OVERFLOW_A000102)
	}

	a := make([]int64, seqlen)
	a[4], a[5], a[6] = 1, 2, 5
	for i := int64(7); i < seqlen; i++ {
		a[i] = 2 * a[i-1] + a[i-2] - 2 * a[i-4] - 3 * a[i-5] - 2 * a[i-6] - a[i-7]
	}
	return a, 0
}

/**
 * A000108 computes the Catalan numbers, which are #s such that
 *  C(n) = binomial(2n,n)/(n+1) = (2n)!/(n!(n+1)!)
 * Date: December 07, 2021
 * Link: https://oeis.org/A000108
 */
func A000108(seqlen int64) ([]int64, int64) {
	if seqlen > OVERFLOW_A000108 {
		utils.OverflowError("A000108", OVERFLOW_A000108)
	}

	a := make([]int64, seqlen)
	a[0], a[1] = 1, 1
	for i := int64(2); i < seqlen; i++ {
		for j := int64(0); j < i; j++ {
			a[i] += a[j] * a[i - j - 1]
		}
	}
	return a, 0
}

/**
 * A000110 computes the Bell or exponential numbers.
 *  Or: # of ways to partition a set of n labelled elements
 * Note: this uses the triangle approach to compute:
 *	Triangle Approach. Bell numbers are the left and right sides of the triangle
 *		 1
 *		 1   2
 *		 2   3   5
 *		 5   7  10  15
 *		15  20  27  37  52
 * Date: December 07, 2021
 * Link: https://oeis.org/A000110
 */
func A000110(seqlen int64) ([]int64, int64) {
	if seqlen > OVERFLOW_A000110 {
		utils.OverflowError("A000110", OVERFLOW_A000110)
	}

	// init
	a := make([]int64, seqlen+1)	// the seq
	a[0] = 1
	old := make([]int64, seqlen)	// last row
	new := make([]int64, seqlen)	// new row
	old[0] = 1

	// compute each row & store into a
	row, col := int64(0), int64(0)
	for ; row < seqlen; row++ {
		col = 0

		// calculate new row
		for  ; col < row; col++ {
			new[col + 1] = new[col] + old[col]
		}

		// copy down
		if col > 0 {
			for i := int64(0); i < col + 1; i++ {
				old[i] = new[i]
				new[i] = 0		// erase new row
			}
		}

		// copy the last element
		new[0] = old[col]		// overwrite first elem
		a[row + 1] = old[col]		// copy last elem of old
	}
	return a, 0
}

/**
 * A000111 computes the Euler zigzag numbers (aka up/down numbers)
 * Date: December 07, 2021
 * Link: https://oeis.org/A000111
 */
func A000111(seqlen int64) ([]int64, int64) {
	if seqlen > OVERFLOW_A000111 {
		utils.OverflowError("A000111", OVERFLOW_A000111)
	}

	a := make([]int64, seqlen)
	for i := int64(1); i <= seqlen; i++ {
		a[i-1] = int64(math.Round(2.0 * math.Pow(2.0 / math.Pi, float64(i)) * float64(utils.Fact(i))))
	}
	return a, 0
}

/**
 * A000115 computes the denumerants (expansion of 1/((1-x)*(1-x^2)*(1-x^5)))
 *  Or: a(n) = round((n+4)^2/20)
 * Date: December 07, 2021
 * Link: https://oeis.org/A000115
 */
func A000115(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	for i := int64(0); i < seqlen; i++ {
		a[i] = int64(math.Round( math.Pow(float64(i+4), 2) / 20))
	}
	return a, 0
}

/**
 * A000116 computes the # of even sequences with period 2n. Also the bisection of A000013
 * Note: A000013 is currently incomplete
 * Date: December 07, 2021
 * Link: https://oeis.org/A000116
 */
func A000116(seqlen int64) ([]int64, int64) {
	if seqlen > OVERFLOW_A000116 {
		utils.OverflowError("A000116", OVERFLOW_A000116)
	}
	a := make([]int64, seqlen)
	euler := make([]int64, seqlen*5)

	// populate euler with the euler totient of i
	for i := int64(1); i < seqlen*5; i++ {
		euler[i - 1] = utils.EulerTotient(i)
	}

	// populate a
	a[0] = 1
	for i := int64(1); i < seqlen; i++ {
		var eidx int64	// euler index
		var b int64		// temp value

		// compute divisors of 2*i
		divisors := utils.Factors(2*i)
		fcount := len(divisors)

		// use these divisors to calculate a
		for j := 0; j < fcount; j++ {
			if (2 * i) % divisors[j] == 0 {
				eidx = 2 * divisors[j] - 1
				b = int64(math.Pow(2.0, float64(2 * i) / float64(divisors[j])))
				a[i] += euler[eidx] * b
			}
		}
		a[i] = a[i] / (4.0 * i)
	}
	return a, 0
}

/**
 * A000120 is the 1s counting seq. It is the # of 1s in the binary expansion of
 *  n (or, the binary weight of n).
 * Date: December 07, 2021
 * Link: https://oeis.org/A000120
 */
 func A000120(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	for i := int64(0); i < seqlen; i++ {
		binary := strconv.FormatInt(i, 2)	// convert to binary
		count := int64(0)
		for _, bit := range binary {
			if bit == '1' {
				count++
			}
		}
		a[i] = count
	}
	return a, 1
}
