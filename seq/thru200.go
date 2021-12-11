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
	LONG_A000158 = 37
	LONG_A000160 = 18
	LONG_A000174 = 50
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
	utils.AccuracyWarning("A000111")

	a := utils.CreateSlice(seqlen)
	for i := int64(1); i <= seqlen; i++ {
		temp := utils.Factorial(big.NewInt(i))
		ifact := ToBigFloat(temp)
		frac := DivFloat(NewFloat(2), NewFloat(math.Pi))
		pow := PowFloat(frac, i)			// (2/pi)^i
		prod := MulFloat(NewFloat(2), pow)	// 2 * (2/pi)^i
		fin := MulFloat(prod, ifact)		// 2 * (2/pi)^i * i!
		a[i-1] = Round(fin)
	}
	return a, 0
}

/**
 * A000115 computes the denumerants (expansion of 1/((1-x)*(1-x^2)*(1-x^5)))
 *  Or: a(n) = round((n+4)^2/20)
 * Date: December 07, 2021	Confirmed working: December 10, 2021
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
 * Date: December 07, 2021
 * Link: https://oeis.org/A000116
 */
func A000116(seqlen int64) ([]*big.Int, int64) {
	// warn the user about inaccuracies
	utils.AccuracyWarning("A000111 (which computes the bisection of A000013)")

	a13, _ := A000013(seqlen*2)
	a := utils.BigBisection(a13)
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

/**
 * A000139 computes a(n) = 2*(3*n)!/((2*n+1)!*((n+1)!)). 
 * Date: December 10, 2021	Confirmed working: December 10, 2021
 * Link: https://oeis.org/A000139
 */
func A000139(seqlen int64) ([]*big.Int, int64) {
	// a(n) = 2(3n)!/((2n+1)!*(n+1)!)
	a := utils.CreateSlice(seqlen)
	for n := int64(0); n < seqlen; n++ {
		nplus1 := utils.Factorial(big.NewInt(n+1))		// (n+1)!
		twonplus1 := utils.Factorial(big.NewInt(2*n+1))	// (2n+1)!
		threen := utils.Factorial(big.NewInt(3*n))		// (3n)!
		numer := Mul(New(2), threen)					// 2(3n)!
		denom := Mul(twonplus1, nplus1)
		a[n] = Floor(DivFloat(ToBigFloat(numer), ToBigFloat(denom)))
	}
	return a, 0
}

/**
 * A000142 generates the sequence of factorials
 * Date: December 10, 2021	Confirmed working: December 10, 2021
 * Link: https://oeis.org/A000142
 */
func A000142(seqlen int64) ([]*big.Int, int64) {
	a := utils.CreateSlice(seqlen)
	a[0] = utils.Factorial(New(0))
	for i := int64(1); i < seqlen; i++ {
		a[i] = Mul(a[i-1], New(i))
	}
	return a, 0
}

/**
 * A000149 computes the seq where a(n) = floor(e^n)
 * Date: December 10, 2021	Confirmed working: December 10, 2021
 * Link: https://oeis.org/A000149
 */
func A000149(seqlen int64) ([]*big.Int, int64) {
	a := utils.CreateSlice(seqlen)
	for i := int64(0); i < seqlen; i++ {
		a[i] = Floor(PowFloat(NewFloat(math.E), i))
	}
	return a, 0
}

/**
 * A000153 computes a(n) = n*a(n-1) + (n-2)*a(n-2), a(0)=0, a(1)=1
 * Date: December 10, 2021	Confirmed working: December 10, 2021
 * Link: https://oeis.org/A000153
 */
func A000153(seqlen int64) ([]*big.Int, int64) {
	a := utils.CreateSlice(seqlen)
	a[0] = New(0)
	a[1] = New(1)

	for i := int64(2); i < seqlen; i++ {
		j := New(i)
		a[i] = Add(Mul(j, a[i-1]), Mul(Sub(j, New(2)), a[i-2]))
	}
	return a, 0
}

/**
 * A000158 computes the # of partitions into non-integral powers
 * Date: December 10, 2021	Confirmed working: December 10, 2021
 * Link: https://oeis.org/A000158
 */
func A000158(seqlen int64) ([]int64, int64) {
	if seqlen > LONG_A000158 {
		utils.LongCalculationWarning("A000158", LONG_A000158)
	}
	check := func(x1, x2, x3, n int64) bool {
		twothirds := 2.0/3.0
		return int64(math.Ceil(math.Pow(float64(x1), twothirds) +
				math.Pow(float64(x2), twothirds) +
				math.Pow(float64(x3), twothirds))) <= n
	}
	a := make([]int64, seqlen)
	for i := int64(0); i < seqlen; i++ {
		// count the # of solutions to the inequality x1^(2/3)+x2^(2/3)+x3^(2/3)<=n
		count := int64(0)
		for x3 := int64(1); x3 <= int64(math.Ceil(math.Pow(float64(i+3), 1.5))); x3++ {
			for x2 := int64(1); x2 <= x3; x2++ {
				for x1 := int64(1); x1 <= x2; x1++ {
					if check(x1, x2, x3, i+3) {
						count++
					}
				}
			}
		}
		a[i] = count
	}
	return a, 3
}

/**
 * A000160 is similar to A000158, except there are 4 terms
 *  i.e. this counts the # of solutions to the inequality x1^(2/3)+x2^(2/3)+x3^(2/3)+x4^(2/3)<=n
 * Date: December 10, 2021	Confirmed working: December 10, 2021
 * Link: https://oeis.org/A000160
 */
func A000160(seqlen int64) ([]int64, int64) {
	if seqlen > LONG_A000160 {
		utils.LongCalculationWarning("A000160", LONG_A000160)
	}
	check := func(x1, x2, x3, x4, n int64) bool {
		twothirds := 2.0/3.0
		return int64(math.Ceil(math.Pow(float64(x1), twothirds) +
				math.Pow(float64(x2), twothirds) +
				math.Pow(float64(x3), twothirds) +
				math.Pow(float64(x4), twothirds))) <= n
	}
	a := make([]int64, seqlen)
	for i := int64(0); i < seqlen; i++ {
		// count the # of solutions to the inequality x1^(2/3)+x2^(2/3)+x3^(2/3)<=n
		count := int64(0)
		for x4 := int64(1); x4 <= int64(math.Ceil(math.Pow(float64(i+4), 1.5))); x4++ {
			for x3 := int64(1); x3 <= x4; x3++ {
				for x2 := int64(1); x2 <= x3; x2++ {
					for x1 := int64(1); x1 <= x2; x1++ {
						if check(x1, x2, x3, x4, i+4) {
							count++
						}
					}
				}
			}
		}
		a[i] = count
	}
	return a, 4
}

/**
 * A000161 computes the # of partitions of n into 2 squares
 * Date: December 10, 2021	Confirmed working: December 10, 2021
 * Link: https://oeis.org/A000161
 */
func A000161(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	for n := int64(0); n < seqlen; n++ {
		count := int64(0)
		for i := int64(0); i <= n; i++ {
			fi := float64(i)
			for j := int64(0); j <= i; j++ {
				fj := float64(j)
				if int64(math.Pow(fi, 2) + math.Pow(fj,2)) == n {
					count++
				}
			}
		}
		a[n] = count
	}
	return a, 0
}

/**
 * A000164 computes the # of partitions of n into 3 squares (allowing part zero)
 * Date: December 10, 2021	Confirmed working: December 10, 2021
 * Link: https://oeis.org/A000164
 */
func A000164(seqlen int64) ([]int64, int64) {
	getCounts := func(n int64) int64 {	// computes the # of partitions of n
		count := int64(0)
		for x := 0; ; x++ {
			xf := float64(x)
			xsq := int64(math.Pow(xf, 2))
			if 3*xsq > n {
				return count
			}
			for y := x; ; y++ {
				yf := float64(y)
				ysq := int64(math.Pow(yf, 2))
				if xsq + 2*ysq > n {
					break
				}
				z2 := n - xsq - ysq;
				if utils.IsSquare(z2) {
					z := math.Sqrt(float64(z2))
					if z >= yf {
						count++
					}
				}
			}
		}
	}

	// generate sequence
	a := make([]int64, seqlen)
	for i := int64(0); i < seqlen; i++ {
		a[i] = getCounts(i)
	}
	return a, 0
}

/**
 * A000165 computes the double factorial of even #s; (2n)!! = 2^n*n!
 * Date: December 10, 2021	Confirmed working: December 10, 2021
 * Link: https://oeis.org/A000165
 */
func A000165(seqlen int64) ([]*big.Int, int64) {
	a := utils.CreateSlice(seqlen)
	for i := int64(0); i < seqlen; i++ {
		a[i] = Mul(Pow(New(2), New(i)), utils.Factorial(New(i)))
	}
	return a, 0
}

/**
 * A000166 computes the subfactorial or rencontres #s (or derangements)
 *  # of permutations of n elements w/ no fixed points
 * Date: December 10, 2021	Confirmed working: December 10, 2021
 * Link: https://oeis.org/A000166
 */
func A000166(seqlen int64) ([]*big.Int, int64) {
	a := utils.CreateSlice(seqlen)
	a[0] = New(1)

	for i := int64(1); i < seqlen; i++ {
		a[i] = Add(Mul(New(i), a[i - 1]), Pow(New(-1), New(i)))
	}

	return a, 0
}

/**
 * A000168 computes a[n] = 2*3^n*(2n)!/(n!*(n+2)!)
 * Date: December 10, 2021	Confirmed working: December 10, 2021
 * Link: https://oeis.org/A000168
 */
func A000168(seqlen int64) ([]*big.Int, int64) {
	a := utils.CreateSlice(seqlen)
	for i := int64(0); i < seqlen; i++ {
		twon := utils.Factorial(Mul(New(2), New(i)))	// (2n)!
		threen := Pow(New(3), New(i))					// 3^n
		numer := Mul(New(2), Mul(twon, threen))			// 2*3^n*(2n)!
		nplus2 := utils.Factorial(Add(New(i), New(2)))	// (n+2)!
		denom := Mul(utils.Factorial(New(i)), nplus2)	// n!*(n+2)!
		a[i] = Div(numer, denom)						// 2*3^n*(2n)!/(n!*(n+2)!)
	}
	return a, 0
}

/**
 * A000169 computes a[n]=n^(n-1)
 * Date: December 10, 2021	Confirmed working: December 10, 2021
 * Link: https://oeis.org/A000169
 */
func A000169(seqlen int64) ([]*big.Int, int64) {
	a := utils.CreateSlice(seqlen)
	for i := int64(1); i <= seqlen; i++ {
		a[i-1] = Pow(New(i), Sub(New(i), New(1)))
	}
	return a, 1
}

/**
 * A000172 computes the Franel #s; a[n] = sum_{k=0..n} binomial(n,k)^3
 * Date: December 10, 2021	Confirmed working: December 10, 2021
 * Link: https://oeis.org/A000172
 */
func A000172(seqlen int64) ([]*big.Int, int64) {
	a := utils.CreateSlice(seqlen)
	for i := int64(0); i < seqlen; i++ {
		for j := int64(0); j <= i; j++ {
			a[i] = Add(a[i], Pow(New(utils.Binomial(New(i).Int64(), New(j).Int64())), New(3)))
		}
	}
	return a, 0
}

/**
 * A000174 computes the # of partitions of n into 5 squares
 * Date: December 10, 2021	Confirmed working: December 10, 2021
 * Link: https://oeis.org/A000174
 */
func A000174(seqlen int64) ([]int64, int64) {
	if seqlen > LONG_A000174 {
		utils.LongCalculationWarning("A000174", LONG_A000174)
	}

	a := make([]int64, seqlen)
	for n := int64(0); n < seqlen; n++ {
		count := int64(0)
		for x5 := int64(0); x5 <= n; x5++ {
			xf5 := float64(x5)
			for x4 := int64(0); x4 <= x5; x4++ {
				xf4 := float64(x4)
				for x3 := int64(0); x3 <= x4; x3++ {
					xf3 := float64(x3)
					for x2 := int64(0); x2 <= x3; x2++ {
						xf2 := float64(x2)
						for x1 := int64(0); x1 <= x2; x1++ {
							xf1 := float64(x1)
							if int64(math.Pow(xf1, 2) +
									math.Pow(xf2, 2) +
									math.Pow(xf3, 2) +
									math.Pow(xf4, 2) +
									math.Pow(xf5, 2)) == n {
								count++
							}
						}
					}
				}
			}
		}
		a[n] = count
	}
	return a, 0
}

/**
 * A000177 computes the # of partitions of n into 6 squares
 * Date: December 10, 2021	Confirmed working: December 10, 2021
 * Link: https://oeis.org/A000177
 */
func A000177(seqlen int64) ([]int64, int64) {
	if seqlen > LONG_A000174 {
		utils.LongCalculationWarning("A000174", LONG_A000174)
	}

	// this is a really nasty algorithm
	a := make([]int64, seqlen)
	for n := int64(0); n < seqlen; n++ {
		count := int64(0)
		for x6 := int64(0); x6 <= n; x6++ {
			for x5 := int64(0); x5 <= x6; x5++ {
				xf5 := float64(x5)
				for x4 := int64(0); x4 <= x5; x4++ {
					xf4 := float64(x4)
					for x3 := int64(0); x3 <= x4; x3++ {
						xf3 := float64(x3)
						for x2 := int64(0); x2 <= x3; x2++ {
							xf2 := float64(x2)
							for x1 := int64(0); x1 <= x2; x1++ {
								xf1 := float64(x1)
								if int64(math.Pow(xf1, 2) +
										math.Pow(xf2, 2) +
										math.Pow(xf3, 2) +
										math.Pow(xf4, 2) +
										math.Pow(xf5, 2)) == n {
									count++
								}
							}
						}
					}
				}
			}
		}		
		a[n] = count
	}
	return a, 0
}

/**
 * A000178 computes the superfactorials, or the product of the first n factorials
 * Date: December 10, 2021	Confirmed working: December 10, 2021
 * Link: https://oeis.org/A000178
 */
func A000178(seqlen int64) ([]*big.Int, int64) {
	a := utils.CreateSlice(seqlen)
	a[0] = utils.Factorial(New(0))
	facts, _ := A000142(seqlen)
	for i := int64(1); i < seqlen; i++ {
		a[i] = Mul(a[i-1], facts[i])
	}
	return a, 0
}
