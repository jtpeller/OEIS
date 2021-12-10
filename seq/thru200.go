package seq

import (
	"OEIS/utils"
	"math"
	"math/big"
	"strconv"
)

const (
	LONG_A000101 = 17
	OVERFLOW_A000102 = 68
	OVERFLOW_A000108 = 36
	OVERFLOW_A000110 = 25
	OVERFLOW_A000111 = 21
	OVERFLOW_A000116 = 32
	OVERFLOW_A000117 = 31
	OVERFLOW_A000126 = 89
	OVERFLOW_A000133 = 7
	OVERFLOW_A000138 = 21
)

/**
 * A000101 computes the record gaps b/w primes (upper end)
 * Date: December 07, 2021	Confirmed working: December 09, 2021
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
	primes := utils.Primes(getPrimeCount(seqlen))
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
 * Date: December 07, 2021	Confirmed working: December 09, 2021
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
 * Date: December 07, 2021	Confirmed working: December 09, 2021
 * Link: https://oeis.org/A000108
 */
func A000108(seqlen int64) ([]*big.Int, int64) {
	if seqlen > OVERFLOW_A000108 {
		utils.BigIntWarning("A000108", OVERFLOW_A000108)
	}

	a := utils.CreateSlice(seqlen)
	a[0], a[1] = big.NewInt(1), big.NewInt(1)
	for i := int64(2); i < seqlen; i++ {
		for j := int64(0); j < i; j++ {
			temp := big.NewInt(0)
			temp.Mul(a[j], a[i - j - 1])
			a[i].Add(a[i], temp)
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
 * Date: December 07, 2021	Confirmed working: December 09, 2021
 * Link: https://oeis.org/A000110
 */
func A000110(seqlen int64) ([]*big.Int, int64) {
	if seqlen > OVERFLOW_A000110 {
		utils.BigIntWarning("A000110", OVERFLOW_A000110)
	}

	// init
	a := utils.CreateSlice(seqlen+1)	// the seq
	a[0] = big.NewInt(1)
	old := utils.CreateSlice(seqlen)	// last row
	new := utils.CreateSlice(seqlen)	// new row
	old[0] = big.NewInt(1)

	// compute each row & store into a
	row, col := int64(0), int64(0)
	for ; row < seqlen; row++ {
		col = 0

		// calculate new row
		for  ; col < row; col++ {
			new[col + 1].Add(new[col], old[col])
		}

		// copy down
		if col > 0 {
			for i := int64(0); i < col + 1; i++ {
				old[i] = new[i]
				new[i] = big.NewInt(0)		// erase new row
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
 * Date: December 07, 2021	Confirmed working: December 09, 2021
 * Link: https://oeis.org/A000111
 */
func A000111(seqlen int64) ([]*big.Int, int64) {
	// warn the user about inaccuracies
	utils.BigFloatWarning("A000111")

	a := utils.CreateSlice(seqlen)
	for i := int64(1); i <= seqlen; i++ {
		temp := utils.Factorial(big.NewInt(i))
		ifact := BigIntToBigFloat(temp)
		frac := DivFloat(NewFloat(2), NewFloat(math.Pi))
		pow := PowFloat(frac, i)			// (2/pi)^i
		prod := MulFloat(NewFloat(2), pow)	// 2 * (2/pi)^i
		fin := MulFloat(prod, ifact)		// 2 * (2/pi)^i * i!
		a[i-1] = RoundFloat(fin)
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
 * A000117 computes the # of even sequences w/ period 2n
 *  Also, the bisection of A000011
 * Date: December 09, 2021
 * Link: https://oeis.org/A000117
 */
func A000117(seqlen int64) ([]int64, int64) {
	// this uses the bisection method. generates double the number of 
	// a000011 terms, and "deletes" every other term
	// because of this, it is necessary to limit the number of terms to 1/2 that of A000011's overflow limit
	if seqlen > OVERFLOW_A000117 {
		utils.OverflowError("A000117", OVERFLOW_A000117)
	}
	a := make([]int64, seqlen)
	a11, _ := A000011(seqlen*2)
	for i := int64(0); i < seqlen; i++ {
		a[i] = a11[2*i]
	}
	return a, 0
}

/**
 * A000118 computes the # of ways of writing n as a sum of 4 squares
 *  Also theta series of lattice Z^4
 * Date: December 09, 2021
 * Link: https://oeis.org/A000118
 */
func A000118(seqlen int64) ([]int64, int64) {
	// generate sigma
	sigma := make([]int64, seqlen)
	for i := int64(1); i < seqlen; i++ {
		divisors := utils.Factors(i)
		for j := 0; j < len(divisors); j++ {
			sigma[i-1] += divisors[j]		// compute sum
		}
	}

	// now generate A000118
	a := make([]int64, seqlen)
	a[0] = 1
	b := int64(0)
	for i := int64(1); i < seqlen; i++ {
		if i % 4 != 0 {
			b = 0
		} else {
			b = 32 * sigma[(i/4) - 1]
		}
		a[i] = 8 * sigma[i-1] - b
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

/**
 * A000123 computes the # of binary partitions: # of partitions of 2n into powers of 2
 * Date: December 09, 2021
 * Link: https://oeis.org/A000123
 */
func A000123(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	a[0] = 1
	for i := int64(1); i < seqlen; i++ {
		a[i] = a[i/2] + a[i - 1]
	}
	return a, 0
}

/**
 * A000124 computes the central polygonal #s (or, the Lazy Caterer's sequence)
 * n(n+1)/2 + 1, or the maximal # of pieces formed when slicing a pancake w/ n cuts
 * Date: December 09, 2021
 * Link: https://oeis.org/A000124
 */
func A000124(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	for i := int64(0); i < seqlen; i++ {
		a[i] = (i * (i + 1)) / 2 + 1
	}

	return a, 0
}

/**
 * A000125 computes the cake #s: the maximal # of pieces resulting from n planar
 * cuts through a cube (or cake)
 * C(n+1,3)+n+1
 * Date: December 09, 2021
 * Link: https://oeis.org/A000125
 */
func A000125(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	for i := int64(0); i < seqlen; i++ {
		a[i] = (int64(math.Pow(float64(i), 3)) + 5*i + 6) / 6
	}
	return a, 0
}

/**
 * A000126 computes a nonlinear binomial sum
 * Date: December 09, 2021
 * Link: https://oeis.org/A000126
 */
func A000126(seqlen int64) ([]int64, int64) {
	if seqlen > OVERFLOW_A000126 {
		utils.OverflowError("A000126", OVERFLOW_A000126)
	}
	a := make([]int64, seqlen)
	a[0], a[1], a[2] = 1, 2, 4

	for i := int64(3); i < seqlen; i++ {
		a[i] = 2 * a[i-1] - a[i-3] + 1
	}
	return a, 1
}

/**
 * A000127 computes the maximal # of regions obtained by joining n points around
 *  a circle by straight lines. Also # of regions in 4-space formed by n-1 hyperplanes
 * Date: December 09, 2021
 * Link: https://oeis.org/A000127
 */
func A000127(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	for i := int64(1); i <= seqlen; i++ {
		j := float64(i)		// conversion to float64 from int64
		a[i-1] = (int64(math.Pow(j, 4)) - 6 * int64(math.Pow(j, 3)) + 23 * int64(math.Pow(j, 2)) - 18 * i + 24) / 24 
	}
	return a, 1
}

/**
 * A000128 computes yet another nonlinear binomial sum
 * Date: December 09, 2021
 * Link: https://oeis.org/A000128
 */
func A000128(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	Fib, _ := A000045(seqlen+5)
	F := utils.ToIntSlice(Fib)
	for i := int64(1); i <= seqlen; i++ {
		a[i-1] = F[i + 4] - i * (i + 1) / 2 - 3
	}
	return a, 1
}

/**
 * A000129 computes the Pell #s: a[0] = 0; a[1] = 1; for n>1, a[n] = 2*a[n-1]+a[n-2]
 * Date: December 09, 2021
 * Link: https://oeis.org/A000129
 */
func A000129(seqlen int64) ([]int64, int64) {
	if seqlen <= 2 {
		utils.TooSmallError("A000129", seqlen)
	}

	a := make([]int64, seqlen)
	a[0], a[1] = 0, 1
	for i := int64(2); i < seqlen; i++ {
		a[i] = 2 * a[i-1] + a[i-2]
	}
	return a, 0
}

/**
 * A000133 computes the # of Boolean functions of n variables
 * Date: December 09, 2021
 * Link: https://oeis.org/A000133
 */
func A000133(seqlen int64) ([]int64, int64) {
	if seqlen > OVERFLOW_A000133 {
		utils.OverflowError("A000133", OVERFLOW_A000133)
	}

	a := make([]int64, seqlen)
	for i := int64(0); i < seqlen; i++ {
		// a(n) = (2^(2^n) + (2^n-1)*2^(2^(n-1)+1))/2^(n+1)
		j := float64(i)
		twoN := math.Pow(2, j)
		a[i] = int64((math.Pow(2, twoN) + (twoN-1) * math.Pow(2, math.Pow(2, j-1) + 1)) / math.Pow(2, j+1))
	}
	return a, 1
}

/**
 * A000138 computes the expansion of e.g.f. exp(-x^4/4)/(1-x).
 * Date: December 09, 2021
 * Link: https://oeis.org/A000138
 */
func A000138(seqlen int64) ([]int64, int64) {
	if seqlen > OVERFLOW_A000138 {
		utils.OverflowError("A000138", OVERFLOW_A000138)
	}

	a := make([]int64, seqlen)
	sum := float64(0)
	for n := int64(0); n < seqlen; n++ {
		sum = 0;		// reset sum

		// a(n) = n! * sum i=0 ... [n/4]( (-1)^i /(i! * 4^i))
		for i := int64(0); i <= int64(math.Floor(float64(n) / 4.0)); i++ {
			j := float64(i)
			powI1 := math.Pow(-1, j);			// (-1)^i
			iFact := float64(utils.Fact(i+1));	// i!
			powI2 := math.Pow(4, j);			// i! * 4^i
			sum += (powI1 / (iFact * powI2));	// (-1)^i /(i! * 4^i)
		}
		a[n] = int64(float64(utils.Fact(n+1)) * sum)
	}
	return a, 0
}

