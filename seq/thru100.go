// ============================================================================
// = thru100.go
// = 	Description		OEIS sequences from A000001-A000100
// = 	Note			Not all sequences in this range have been programmed
// = 	Date			October 08, 2021
// ============================================================================

package seq

import (
	"OEIS/utils"
	"math"
	"strconv"
)

/**
 * A000002 returns the Kolakoski sequence, given a sequence length
 * Date		October 08, 2021
 * Link		https://oeis.org/A000002
 */
func A000002(seqlen int64) ([]int64, int64) {
	return utils.Kolakoski(seqlen+1, 2)[:seqlen], 1
}

/**
 * A000004 returns a slice of length seqlen (default init'd to 0)
 * Date		October 08, 2021
 * Link		https://oeis.org/A000004
 */
func A000004(seqlen int64) ([]int64, int64) {
	return make([]int64, seqlen), 0
}

/**
 * A000005 returns the # of divisors of n, given a seq len
 * Date		October 08, 2021
 * Link		https://oeis.org/A000005
 */
func A000005(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	for i := int64(0); i < seqlen; i++ {
		count := utils.GetFactorCount(i + 1)
		a[i] = count
	}
	return a, 1
}

/**
 * A000006 returns the isqrt of numbers, given a seq len
 * Date		October 08, 2021
 * Link		https://oeis.org/A000006
 */
func A000006(seqlen int64) ([]int64, int64) {
	primes := utils.Primes(seqlen)
	a := utils.Isqrtarray(primes)
	return a, 1
}

/**
 * A000007 returns a sequence of len seqlen, where a(n) = 0^n
 * Date		October 08, 2021
 * Link		https://oeis.org/A000007
 */
func A000007(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	a[0] = 1
	return a, 0
}

/**
 * A000008 returns the # of ways of making change for n cents using coins of 1, 2, 5, 10 cents.
 * Date		October 08, 2021
 * Link		https://oeis.org/A000008
 */
func A000008(seqlen int64) ([]int64, int64) {
	utils.LongCalculationWarning("A000008")

	denoms := []int64{1, 2, 5, 10}
	a := make([]int64, 0)
	coins := int64(len(denoms))

	for i := int64(0); i < seqlen; i++ {
		a = append(a, utils.MakeChange(coins, i, denoms))
	}
	return a, 0
}

/**
 * A000010 computes the Euler totient function phi(n):
 * 		count numbers <= n and prime to n.
 * Date		October 08, 2021
 * Link		https://oeis.org/A000010
 */
func A000010(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	for i := int64(1); i <= seqlen; i++ {
		a[i-1] = utils.EulerTotient(i)
	}
	return a, 1
}

/**
 * A000011 returns the # of n-bead necklaces where turning over is allowed.
 * Date 	October 08, 2021
 * Fixed  	2025.02.01
 * Link 	https://oeis.org/A000011
 */
func A000011(seqlen int64) ([]*bint, int64) {
	// generate euler phi
	euler, _ := A000010(seqlen * 2)

	// generate a sequence
	a := iSlice(seqlen)
	a[0] = inew(1)
	for n := int64(1); n < seqlen; n++ {
		divisors := utils.Factors(n)
		sum := fpow(fnew(2), int64(n/2))

		// use the divisors to calculate the sequence
		for _, d := range divisors {
			// phi(2*d) (-1 for 0 indexing)
			phi := itof(inew(euler[2*d-1]))

			// 2^(n/d)
			b := fnew(math.Pow(2, float64(n)/float64(d)))

			// (phi*b) / (2*n)
			numer := fmul(phi, b)
			frac := fdiv(numer, fmul(fnew(2), itof(inew(n))))

			// sum += frac
			sum = fadd(sum, frac)
		}

		// return s/2
		a[n] = round(fdiv(sum, fnew(2)))

	}
	return a, 0
}

/**
 * A000012 returns a seq of all 1s, of len seqlen
 * Date		October 08, 2021
 * Link		https://oeis.org/A000012
 */
func A000012(seqlen int64) ([]int64, int64) {
	a := make([]int64, 0)
	for i := int64(0); i < seqlen; i++ {
		a = append(a, 1)
	}
	return a, 0
}

/**
 * A000013 computes # of n-bead binary necklaces with beads of 2 colors where
 *  the colors may be swapped but turning over is not allowed.
 * Date		December 10, 2021
 * Fix		2025.02.01
 * Link		https://oeis.org/A000013
 */
func A000013(seqlen int64) ([]*bint, int64) {
	a := iSlice(seqlen)
	a[0] = inew(1)
	phi, _ := A000010(seqlen * 2)
	for n := int64(1); n < seqlen; n++ {
		// calculate the sum
		sum := fnew(0)
		for d := int64(1); d <= n; d++ {
			if n%d == 0 { // d divides n
				temp := itof(inew(phi[2*d-1]))
				pow := math.Pow(2, float64(n)/float64(d))
				numer := fmul(temp, fnew(pow))
				n2 := mul(inew(2), inew(n))
				bigDiv := fdiv(numer, itof(n2))
				// the following computes a[n] = Sum_{d divides n} (phi(2*d)*2^(n/d))/(2*n)
				sum = fadd(sum, bigDiv)
			}
		}
		a[n] = round(sum)
	}
	return a, 0
}

/**
 * A000018 computes the # of positive integers <= 2^n of form x^2 + 16y^2
 * Date		December 12, 2021	Confirmed working: December 12, 2021
 * Link		https://oeis.org/A000018
 */
func A000018(seqlen int64) ([]*bint, int64) {
	utils.LongCalculationWarning("A000018")
	a := utils.Repr(seqlen, 1, 16, 1)
	return a, 0
}

/**
 * A000021 computes the # of positive integers <= 2^n of form x^2 + 12y^2
 * Date		December 12, 2021	Confirmed working: December 12, 2021
 * Link		https://oeis.org/A000021
 */
func A000021(seqlen int64) ([]*bint, int64) {
	utils.LongCalculationWarning("A000021")
	a := utils.Repr(seqlen, 1, 12, 1)
	return a, 0
}

/**
 * A000024 computes the # of positive integers <= 2^n of form x^2 + 10y^2
 * Date		December 12, 2021	Confirmed working: December 12, 2021
 * Link		https://oeis.org/A000024
 */
func A000024(seqlen int64) ([]*bint, int64) {
	utils.LongCalculationWarning("A000024")
	a := utils.Repr(seqlen, 1, 10, 1)
	return a, 0
}

/**
 * A000027 returns a seq of positive integers, of len seqlen
 * Date		October 09, 2021
 * Link		https://oeis.org/
 */
func A000027(seqlen int64) ([]int64, int64) {
	a := make([]int64, 0)
	for i := int64(0); i < seqlen+1; i++ {
		a = append(a, i+1)
	}
	return a, 1
}

/**
 * A000030 returns the sequence of the first digit of n, of len seqlen
 * Date		October 09, 2021
 * Link		https://oeis.org/A000030
 */
func A000030(seqlen int64) ([]int64, int64) {
	a := make([]int64, 0)
	for i := int64(0); i < seqlen; i++ {
		a = append(a, utils.GetFirstDigit(i))
	}
	return a, 0
}

/**
 * A000032 computes the Lucas numbers, beginning at 2: L(n) = L(n-1) + L(n-2),
 * 		L(0) = 2, L(1) = 1.
 * Date		October 09, 2021
 * Link		https://oeis.org/A000032
 */
func A000032(seqlen int64) ([]*bint, int64) {
	a := iSlice(seqlen)
	a[0] = inew(2) // a(0)=2
	a[1] = inew(1) // a(1)=1
	for i := int64(2); i < seqlen; i++ {
		a[i].Add(a[i-2], a[i-1])
	}
	return a, 0
}

/**
 * A000034 returns a(n) = 1 + (n mod 2), or 1 + A000035(n)
 * Date		October 08, 2021
 * Link		https://oeis.org/A000034
 */
func A000034(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	for i := int64(0); i < seqlen; i++ {
		a[i] = i%2 + 1
	}
	return a, 0
}

/**
 * A000035 computes the parity of n (basically, n mod 2)
 * Date		October 09, 2021
 * Link		https://oeis.org/A000035
 */
func A000035(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	for i := int64(0); i < seqlen; i++ {
		a[i] = i % 2
	}
	return a, 0
}

/**
 * A000037 computes the nonsquares
 * Date		October 09, 2021
 * Link		https://oeis.org/A000037
 */
func A000037(seqlen int64) ([]int64, int64) {
	a := make([]int64, 0)
	for i := int64(1); i < seqlen; i++ {
		// a(n) = n + floor(1/2 + sqrt(n))
		a = append(a, i+int64(math.Floor(0.5+math.Sqrt(float64(i)))))
	}
	return a, 1
}

/**
 * A000038 computes 2*A000007
 * Date		October 09, 2021
 * Link		https://oeis.org/A000038
 */
func A000038(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	a[0] = 2
	return a, 0
}

/**
 * A000040 computes prime numbers using Golang's built-in
 * Date		October 09, 2021
 * Link		https://oeis.org/A000040
 */
func A000040(seqlen int64) ([]int64, int64) {
	a := make([]int64, 0)
	prime := int64(0)
	for i := int64(0); i < seqlen; {
		if utils.IsPrime(prime) {
			a = append(a, prime)
			i++
		}
		prime++
	}
	return a, 1
}

/**
 * A000041 generates the # of partitions of n
 * Date		October 09, 2021
 * Link		https://oeis.org/A000041
 */
func A000041(seqlen int64) ([]int64, int64) {
	utils.LongCalculationWarning("A000041")

	a := make([]int64, seqlen)
	for i := int64(0); i < seqlen; i++ {
		a[i] = utils.CountParts(i)
	}
	return a, 1
}

/**
 * A000042 generates the unary representation of the natural numbers
 * Date		October 09, 2021
 * Link		https://oeis.org/A000042
 */
func A000042(seqlen int64) ([]*bint, int64) {
	a := iSlice(seqlen)
	a[0] = inew(1)
	for i := int64(1); i < seqlen; i++ {
		a[i] = add(mul(a[i-1], inew(10)), inew(1))
	}
	return a, 1
}

/**
 * A000043 generates the Mersenne exponents.
 * Date		December 07, 2021
 * Link		https://oeis.org/A000043
 */
func A000043(seqlen int64) ([]int64, int64) {
	utils.LongCalculationWarning("A000043")

	a := make([]int64, seqlen)
	prime := int64(0)
	i := int64(0)
	for i < seqlen {
		//mersenne := int64(math.Pow(2, float64(prime))) - 1
		merprime := zero()
		merprime.Exp(inew(2), inew(prime), inew(0))
		merprime.Sub(merprime, inew(1))
		if merprime.ProbablyPrime(20) && utils.IsPrime(prime) {
			a[i] = prime
			i++
		}
		prime++
	}
	return a, 1
}

/**
 * A000044 generates the Dying rabbits sequence, where a[0] = 1, a[1:12] =
 *		Fibonacci(n), and a[13:] = a[n-1] + a[n-2] - a[n-13].
 * Date December 07, 2021
 * Link: https://oeis.org/A000044
 */
func A000044(seqlen int64) ([]*bint, int64) {
	if seqlen <= 12 {
		utils.PrintWarning("For best results, sequence A000044 should have more than 12 elements")
	}

	a := iSlice(seqlen + 1)
	a[0] = inew(1)

	// for [1:12], a(n) = Fibonacci(n)
	a[1] = inew(1)
	a[2] = inew(1)
	bound := int64(12)
	if seqlen <= 12 {
		bound = seqlen
	}
	for i := int64(3); i <= bound; i++ {
		a[i].Add(a[i-1], a[i-2])
	}

	// for n >= 13, a(n) = a(n-1) + a(n-2) - a(n-13)
	for i := int64(13); i <= seqlen; i++ {
		a[i].Add(a[i-1], a[i-2])
		a[i].Sub(a[i], a[i-13])
	}

	return a, 0
}

/**
 * A000045 returns the Fibonacci numbers, of len seqlen
 * Date		December 07, 2021
 * Link		https://oeis.org/A000045
 */
func A000045(seqlen int64) ([]*bint, int64) {
	a := utils.Nacci(seqlen, 2, true)
	return a, 0
}

/**
 * A000047 computes the # of positive integers <= 2^n of form x^2 + 2y^2
 * Date		December 12, 2021	Confirmed working: December 12, 2021
 * Link		https://oeis.org/A000047
 */
func A000047(seqlen int64) ([]*bint, int64) {
	utils.LongCalculationWarning("A000047")

	a := utils.Repr(seqlen, 1, -2, 1)
	return a, 0
}

/**
 * A000049 computes the # of positive integers <= 2^n of form 3x^2 + 4y^2
 * Date: December 12, 2021	Confirmed working: December 12, 2021
 * Link: https://oeis.org/A000049
 */
func A000049(seqlen int64) ([]*bint, int64) {
	utils.LongCalculationWarning("A000049")
	a := utils.Repr(seqlen, 3, 4, 0)
	return a, 0
}

/**
 * A000050 computes the # of positive integers <= 2^n of form x^2 + y^2
 * Date		December 12, 2021	Confirmed working: December 12, 2021
 * Link		https://oeis.org/A000050
 */
func A000050(seqlen int64) ([]*bint, int64) {
	utils.LongCalculationWarning("A000050")
	a := utils.Repr(seqlen, 1, 1, 1)
	return a, 0
}

/**
 * A000051 computes a(n)=2^n + 1
 * Date		December 12, 2021	Confirmed working: December 12, 2021
 * Link		https://oeis.org/A000051
 */
func A000051(seqlen int64) ([]*bint, int64) {
	a, _ := A000079(seqlen)
	for i := int64(0); i < seqlen; i++ {
		a[i] = add(a[i], inew(1))
	}
	return a, 0
}

/**
 * A000058 returns Sylvester's sequence: a(n+1) = a(n)^2 - a(n) + 1
 * Date		December 07, 2021
 * Link		https://oeis.org/A000058
 */
func A000058(seqlen int64) ([]*bint, int64) {
	a := iSlice(seqlen)
	a[0] = inew(2)
	for i := int64(0); i < seqlen-1; i++ {
		a[i+1].Exp(a[i], inew(2), inew(0))
		a[i+1].Sub(a[i+1], a[i])
		a[i+1].Add(a[i+1], inew(1))
	}
	return a, 0
}

/**
 * A000059 returns the sequence a(n) such that (2n)^4 + 1 is prime
 * Date		December 07, 2021
 * Link		https://oeis.org/A000059
 */
func A000059(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	prime := int64(0)
	for i := int64(0); i < seqlen; {
		val := int64(math.Pow(float64(2*prime), 4)) + 1 // (2n)^4 + 1
		if utils.IsPrime(val) {
			a[i] = prime
			i++
		}
		prime++
	}
	return a, 1
}

/**
 * A000062 generates a Beatty sequence; where a(n) = floor(n/(e-2)).
 * Date		December 07, 2021
 * Link		https://oeis.org/A000062
 */
func A000062(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	for i := int64(1); i <= seqlen; i++ {
		a[i-1] = int64(math.Floor(float64(i) / (float64(math.E) - 2)))
	}
	return a, 1
}

/**
 * A000064 generates the partial sums of A000008
 * Date		December 07, 2021
 * Link		https://oeis.org/A000064
 */
func A000064(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	a8, _ := A000008(seqlen)
	for i := int64(1); i <= seqlen; i++ {
		a[i-1] = utils.Sum(a8[:i])
	}
	return a, 0
}

/**
 * A000065 computes -1 + the # of partitions of n
 * Date		December 07, 2021
 * Link		https://oeis.org/A000065
 */
func A000065(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	a10, _ := A000041(seqlen)
	for i := int64(0); i < seqlen; i++ {
		a[i] = a10[i] - 1
	}
	return a, 0
}

/**
 * A000068 returns a sequence such that n^4 + 1 is prime.
 * Date		December 07, 2021
 * Link		https://oeis.org/A000068
 */
func A000068(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	prime := int64(0)
	for i := int64(0); i < seqlen; {
		val := int64(math.Pow(float64(prime), 4)) + 1 // (2n)^4 + 1
		if utils.IsPrime(val) {
			a[i] = prime
			i++
		}
		prime++
	}
	return a, 1
}

/**
 * A000069: the Odious numbers; #s with an odd # of ones in their binary expansion
 * Date		December 07, 2021
 * Link		https://oeis.org/A000069
 */
func A000069(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	num := int64(0)
	for i := int64(0); i < seqlen; {
		binary := strconv.FormatInt(num, 2) // convert to binary
		count := int64(0)
		for _, bit := range binary {
			if bit == '1' {
				count++
			}
		}
		if count%2 != 0 {
			a[i] = num
			i++
		}
		num++
	}
	return a, 1
}

/**
 * A000070: Series of the number of partitions, i.e.
 * 	a[n] = Sum(p[:i]), where p[k] = # of partitions of k
 * Date		December 07, 2021
 * Link		https://oeis.org/A000070
 */
func A000070(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	p, _ := A000041(seqlen)
	for i := int64(1); i <= seqlen; i++ {
		a[i-1] = utils.Sum(p[:i])
	}
	return a, 0
}

/**
 * A000071 generates a(n), where a(n) = Fibonacci(n) - 1.
 * 		For some reason, the offset is 1 here, but 0 for A000045
 * Date		December 07, 2021
 * Link		https://oeis.org/A000071
 */
func A000071(seqlen int64) ([]*bint, int64) {
	a := iSlice(seqlen)
	F, _ := A000045(seqlen + 1)
	for i := int64(1); i <= seqlen; i++ {
		a[i-1].Sub(F[i], inew(1))
	}
	return a, 1
}

/**
 * A000073: Tribonacci #s: a[n] = a[n-1] + a[n-2] + a[n-3]
 *	a[0] = a[1] = 0 and a[2] = 1
 * Date		December 07, 2021
 * Link		https://oeis.org/A000073
 */
func A000073(seqlen int64) ([]*bint, int64) {
	a := iSlice(seqlen)
	a[0], a[1], a[2] = inew(0), inew(0), inew(1)
	for i := int64(3); i < seqlen; i++ {
		a[i].Add(a[i-1], a[i-2])
		a[i].Add(a[i], a[i-3])
	}
	return a, 0
}

/**
 * A000078: Tetranacci #s: a(n) = a(n-1) + a(n-2) + a(n-3) + a(n-4)
 *  for n >= 4 with a(0) = a(1) = a(2) = 0 and a(3) = 1.
 * Date		December 07, 2021
 * Link		https://oeis.org/A000078
 */
func A000078(seqlen int64) ([]*bint, int64) {
	a := utils.InitBslice(seqlen, []*bint{inew(0), inew(0), inew(0), inew(1)})
	for i := int64(4); i < seqlen; i++ {
		a[i].Add(a[i-1], a[i-2])
		a[i].Add(a[i], a[i-3])
		a[i].Add(a[i], a[i-4])
	}
	return a, 0
}

/**
 * A000079: Powers of 2: a(n) = 2^n
 * Date		December 07, 2021
 * Link		https://oeis.org/A000079
 */
func A000079(seqlen int64) ([]*bint, int64) {
	a := utils.Powers(seqlen, inew(2))
	return a, 0
}

/**
 * A000082: a(n) = n^2*Product_{p|n} (1 + 1/p)
 * Note		There may be some rounding error due to float64 <-> int64 conversions
 * Date		December 07, 2021
 * Link		https://oeis.org/A000082
 */
func A000082(seqlen int64) ([]int64, int64) {
	// this computes the Product_{p|n} (1 + 1/p) part
	gen := func(num int64) float64 {
		prod := 1.0
		arr := utils.Factors(num)

		// compute the product
		for i := 0; i < len(arr); i++ {
			if utils.IsPrime(arr[i]) {
				prod *= (1 + float64(1)/float64(arr[i]))
			}
		}
		return prod
	}

	a := make([]int64, seqlen)
	for i := int64(1); i <= seqlen; i++ {
		a[i-1] = int64(math.Round(math.Pow(float64(i), 2) * gen(i)))
	}
	return a, 1
}

/**
 * A000086 returns the # of solutions to x^2 - x + 1 == 0 (mod n)
 * Date		December 12, 2021	Confirmed working: December 12, 2021
 * Link		https://oeis.org/A000086
 */
func A000086(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	for n := int64(1); n < seqlen; n++ {
		// count soln's to x^2 - x + 1 == 0
		count := int64(0)
		for x := int64(1); x <= n; x++ {
			out := int64(math.Pow(float64(x), 2)) - x + 1
			if out%n == 0 {
				count++
			}
		}
		a[n-1] = count
	}
	return a, 1
}

/**
 * A000093 calculates a(n) = floor(n^(3/2))
 * Date		December 07, 2021
 * Link		https://oeis.org/A000093
 */
func A000093(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	for i := int64(0); i < seqlen; i++ {
		a[i] = int64(math.Floor(math.Pow(float64(i), 1.5)))
	}
	return a, 0
}

/**
 * A000094 computes the # of trees of diameter 4
 *  Or: a(n+1) = A000041(n) - n, n > 0
 * Date		December 07, 2021
 * Link		https://oeis.org/A000094
 */
func A000094(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	a41, _ := A000041(seqlen)
	for i := int64(1); i < seqlen; i++ {
		a[i] = a41[i] - i
	}
	return a, 1
}

/**
 * A000096 computes a(n) = n*(n+3)/2
 * Date		December 07, 2021
 * Link		https://oeis.org/A000096
 */
func A000096(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	for i := int64(0); i < seqlen; i++ {
		a[i] = int64(i * (i + 3.0) / 2.0)
	}
	return a, 0
}

/**
 * A000097 computes the # of partitions of n if there are two kinds of 1s and
 *  two kinds of 2s
 * Date		December 07, 2021
 * Link		https://oeis.org/A000097
 */
func A000097(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	a70, _ := A000070(seqlen)
	for i := int64(0); i < seqlen; i++ {
		bound := int64(math.Floor(float64(i) / 2.0))
		for j := int64(0); j <= bound; j++ {
			a[i] += a70[i-2*j]
		}
	}
	return a, 0
}

/**
 * A000098 computes the # of partitions of n if there are two kinds of 1s,
 *  two kinds of 2s, and two kinds of 3s
 * Date		December 07, 2021
 * Link		https://oeis.org/A000098
 */
func A000098(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	a97, _ := A000097(seqlen)
	for i := int64(0); i < seqlen; i++ {
		bound := int64(math.Floor(float64(i) / 3.0))
		for j := int64(0); j <= bound; j++ {
			a[i] += a97[i-3*j]
		}
	}
	return a, 0
}

/**
 * A000100 computes the # of compositions of n in which the maximal part is 3
 * Date		December 07, 2021
 * Link		https://oeis.org/A000100
 */
func A000100(seqlen int64) ([]*bint, int64) {
	a := iSlice(seqlen)
	Fib, _ := A000045(seqlen)
	a[0], a[1], a[2] = zero(), zero(), zero()
	a[3], a[4] = inew(1), inew(2)
	for i := int64(5); i < seqlen; i++ {
		a[i] = addall(Fib[i-2], a[i-3], a[i-2], a[i-1])
	}
	return a, 0
}
