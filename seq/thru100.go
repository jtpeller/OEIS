// ============================================================================
// = thru100.go																  =
// = 	Description: All OEIS sequences from A000001-A000100				  =
// = 	Note: Not all sequences in this range have been programmed			  =
// = 	Date: October 08, 2021												  =
// = 	Last Update: December 07, 2021										  =
// ============================================================================

package seq

import (
	"OEIS/utils"
	"math"
	"strconv"
)

const (
	LONG_A000008 = 350
	OVERFLOW_A000011 = 63
	OVERFLOW_A000032 = 91
	LONG_A000041 = 50
	OVERFLOW_A000042 = 19
	OVERFLOW_A000043 = 9
	OVERFLOW_A000044 = 92
	OVERFLOW_A000045 = 93
	OVERFLOW_A000058 = 7
	OVERFLOW_A000073 = 75
)

/**
 * A000002 returns the Kolakoski sequence, given a sequence length
 * Date: October 08, 2021
 * Link: https://oeis.org/A000002
 */
 func A000002(seqlen int64) ([]int64, int64) {
	return utils.Kolakoski(seqlen+1, 2)[:seqlen], 1
}

/**
 * A000004 returns a slice of length seqlen (default init'd to 0)
 * Date: October 08, 2021
 * Link: https://oeis.org/A000004
 */
 func A000004(seqlen int64) ([]int64, int64) {
	return make([]int64, seqlen), 0
}

/**
 * A000005 returns the # of divisors of n, given a seq len
 * Date: October 08, 2021
 * Link: https://oeis.org/A000005
 */
 func A000005(seqlen int64) ([]int64, int64) {
	a := make([]int64, 0)
	for i := int64(1); i < seqlen; i++ {
		count := utils.GetFactorCount(i)
		a = append(a, count)
	}
	return a, 1
}

/**
 * A000006 returns the isqrt of numbers, given a seq len
 * Date: October 08, 2021
 * Link: https://oeis.org/A000006
 */
 func A000006(seqlen int64) ([]int64, int64) {
	primes, err := utils.Primes(seqlen)
	utils.CheckError(err)
	a := utils.Isqrtarray(primes)
	return a, 1
}

/**
 * A000007 returns a sequence of len seqlen, where a(n) = 0^n
 * Date: October 08, 2021
 * Link: https://oeis.org/A000007
 */
 func A000007(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	a[0] = 1
	return a, 0
}

/** 
 * A000008 returns the # of ways of making change for n cents using coins of 1, 2, 5, 10 cents.
 * Date: October 08, 2021
 * Link: https://oeis.org/A000008
 */
func A000008(seqlen int64) ([]int64, int64) {
	if seqlen > LONG_A000008 {
		utils.LongCalculationWarning("A000008", LONG_A000008)
	}

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
 * Date: October 08, 2021
 * Link: https://oeis.org/A000010
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
 * Note: this implementation is slightly inaccurate due to rounding errors
 * Date: October 08, 2021
 * Link: https://oeis.org/A000011
 */
 func A000011(seqlen int64) ([]int64, int64) {
	if seqlen > OVERFLOW_A000011 {
		utils.OverflowError("A000011", OVERFLOW_A000011)
	}

	// generate euler phi
	eulerlen := seqlen * 2
	euler := make([]int64, 0)
	for i := int64(0); i < eulerlen; i++ {
		euler = append(euler, utils.EulerTotient(i))
	}

	// generate even sequence
	even := make([]int64, seqlen)
	for i := int64(0); i < seqlen; i++ {
		if i == 0 {
			even[i] = 1
		} else {
			divisors := utils.Factors(int64(i))
			euleridx := int64(0)

			// use the divisors to calculate the sequence
			factorcount := len(divisors)
			for j := 0; j < factorcount; j++ {
				euleridx = 2 * divisors[j] - 1
				b := math.Pow(2, float64(i) / float64(divisors[j]))
				even[i] += int64(float64(euler[euleridx]) * b)
			}
			even[i] = even[i] / (2.0 * i)
			foo := math.Pow(2, float64(i / 2))
			even[i] += int64(foo)
			even[i] = even[i] / 2
		}
	}

	return even, 0
}

/**
 * A000012 returns a seq of all 1s, of len seqlen
 * Date: October 08, 2021
 * Link: https://oeis.org/A000012
 */
 func A000012(seqlen int64) ([]int64, int64) {
	a := make([]int64, 0)
	for i := int64(0); i < seqlen; i++ {
		a = append(a, 1)
	}
	return a, 0
}

/** 
 * A000027 returns a seq of positive integers, of len seqlen
 * Date: October 09, 2021
 * Link: https://oeis.org/
 */
 func A000027(seqlen int64) ([]int64, int64) {
	if seqlen <= 0 {
		utils.PositiveError("A000027")
	}

	a := make([]int64, 0)
	for i := int64(0); i < seqlen + 1; i++ {
		a = append(a, i+1)
	}
	return a, 1
}

/**
 * A000030 returns the sequence of the first digit of n, of len seqlen
 * Date: October 09, 2021
 * Link: https://oeis.org/A000030
 */
 func A000030(seqlen int64) ([]int64, int64) {
	if seqlen <= 0 {
		utils.PositiveError("A000030")
	}

	a := make([]int64, 0)
	for i := int64(0); i < seqlen; i++ {
		a = append(a, utils.GetFirstDigit(i))
	}
	return a, 0
}

/**
 * A000032 computes teh Lucas numbers, beginning at 2: L(n) = L(n-1) + L(n-2),
 * 		L(0) = 2, L(1) = 1.
 * Date: October 09, 2021
 * Link: https://oeis.org/A000032
 */
 func A000032(seqlen int64) ([]int64, int64) {
	if seqlen <= 0 {
		utils.PositiveError("A000032")
	} else if seqlen > OVERFLOW_A000032 {
		utils.OverflowError("A000032", OVERFLOW_A000032)
	}

	a := make([]int64, seqlen)
	a[0] = 2		// a(0)=2
	a[1] = 1		// a(1)=1	
	for i := int64(2); i < seqlen; i++ {
		a[i] = a[i - 2] + a[i - 1]
	}
	return a, 0
}

/**
 * A000035 returns a(n) = 1 + (n mod 2), or 1 + A000035(n)
 * Date: October 08, 2021
 * Link: https://oeis.org/A000034
 */
 func A000034(seqlen int64) ([]int64, int64) {
	if seqlen <= 0 {
		utils.PositiveError("A000034")
	}

	a := make([]int64, seqlen)
	for i := int64(0); i < seqlen; i++ {
		a[i] = i % 2 + 1
	}
	return a, 0
}

/**
 * A000035 computes the parity of n (basically, n mod 2)
 * Date: October 09, 2021
 * Link: https://oeis.org/A000035
 */
 func A000035(seqlen int64) ([]int64, int64) {
	if seqlen <= 0 {
		utils.PositiveError("A000035")
	}

	a := make([]int64, seqlen)
	for i := int64(0); i < seqlen; i++ {
		a[i] = i % 2
	}
	return a, 0
}

/**
 * A000037 computes the nonsquares
 * Date: October 09, 2021
 * Link: https://oeis.org/A000037
 */
 func A000037(seqlen int64) ([]int64, int64) {
	if seqlen <= 0 {
		utils.PositiveError("A000037")
	}

	a := make([]int64, 0)
	for i := int64(1); i < seqlen; i++ {
		a = append(a, i + int64(math.Floor(0.5 + math.Sqrt(float64(i)))))
	}
	return a, 1
}

/**
 * A000038 computes 2*A000007
 * Date: October 09, 2021
 * Link: https://oeis.org/A000038
 */
 func A000038(seqlen int64) ([]int64, int64) {
	if seqlen <= 0 {
		utils.PositiveError("A000038")
	}
	
	a := make([]int64, seqlen)
	a[0] = 2
	return a, 0
}

/**
 * A000040 computes prime numbers using Golang's built-in
 * Date: October 09, 2021
 * Link: https://oeis.org/A000040
 */
 func A000040(seqlen int64) ([]int64, int64) {
	if seqlen <= 0 {
		utils.PositiveError("A000040")
	}

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
 * Date: October 09, 2021
 * Link: https://oeis.org/A000041
 */
 func A000041(seqlen int64) ([]int64, int64) {
	if seqlen <= 0 {
		utils.PositiveError("A000041")
	} else if seqlen > LONG_A000041 {
		utils.LongCalculationWarning("A000041", LONG_A000041)
	}

	a := make([]int64, 0)
	for i := int64(0); i < seqlen; i++ {
		a = append(a, utils.CountParts(i))
	}
	return a, 1
}

/**
 * A000042 generates the unary representation of the natural numbers
 * Date: October 09, 2021
 * Link: https://oeis.org/A000042
 */
 func A000042(seqlen int64) ([]int64, int64) {
	if seqlen >= OVERFLOW_A000042 {
		utils.OverflowError("A000042", OVERFLOW_A000042)
	}

	a := make([]int64, 0)
	for i := int64(1); i < seqlen; i++ {
		temp := int64(1)
		for j := int64(1); j < i; j++ {
			temp = temp + int64(math.Pow(10, float64(j)))
		}
		a = append(a, temp)
	}
	return a, 1
}

/**
 * A000043 generates a seqlen-long sequence of the Mersenne exponents.
 * Date: December 07, 2021
 * Link: https://oeis.org/A000043
 */
 func A000043(seqlen int64) ([]int64, int64) {
	if seqlen > OVERFLOW_A000043 {
		utils.OverflowError("A000043 (which computes 2^p-1)", OVERFLOW_A000043)
	}
	a := make([]int64, seqlen)
	prime := int64(0)
	i := int64(0)
	for i < seqlen {
		mersenne := int64(math.Pow(2, float64(prime))) - 1
		if utils.IsPrime(mersenne) && utils.IsPrime(prime) {
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
 func A000044(seqlen int64) ([]int64, int64) {
	if seqlen <= 12 {
		utils.PrintWarning("for best results, sequence A000044 should have more than 12 elements")
	} else if seqlen > OVERFLOW_A000044 {
		utils.OverflowError("A000044", OVERFLOW_A000044)
	}
	a := make([]int64, seqlen+1)
	a[0] = 1

	// for [1:12], a(n) = Fibonacci(n)
	a[1] = 1
	a[2] = 1
	bound := int64(12)
	if seqlen <= 12 {
		bound = seqlen
	}
	for i := int64(3); i <= bound; i++ {
		a[i] = a[i-1] + a[i-2]
	}

	// for n >= 13, a(n) = a(n-1) + a(n-2) - a(n-13)
	for i := int64(13); i <= seqlen; i++ {
		a[i] = a[i-1] + a[i-2] - a[i-13]
	}

	return a, 0
}

/**
 * A000045 returns the Fibonacci numbers, of len seqlen
 * Date: December 07, 2021
 * Link: https://oeis.org/A000045
 */
 func A000045(seqlen int64) ([]int64, int64) {
	// error checking
	if seqlen > OVERFLOW_A000045 {
		utils.OverflowError("A000045", OVERFLOW_A000045)
	}

	// generate fibonacci numbers
	a := make([]int64, seqlen)
	a[0] = 0
	a[1] = 1
	for i := int64(2); i < seqlen; i++ {
		a[i] = a[i-1] + a[i-2]
	}
	return a, 0
}

/**
 * A000058 returns Sylvester's sequence: a(n+1) = a(n)^2 - a(n) + 1
 * Date: December 07, 2021
 * Link: https://oeis.org/A000058
 */
 func A000058(seqlen int64) ([]int64, int64) {
	// error checking
	if seqlen > OVERFLOW_A000058 {
		utils.OverflowError("A000058", OVERFLOW_A000058)
	}

	a := make([]int64, seqlen)
	a[0] = 2
	for i := int64(0); i < seqlen - 1; i++ {
		a[i+1] = int64( math.Pow(float64(a[i]), 2) ) - a[i] + 1
	}

	return a, 0
}

/**
 * A000059 returns the sequence a(n) such that (2n)^4 + 1 is prime
 * Date: December 07, 2021
 * Link: https://oeis.org/A000059
 */
 func A000059(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	prime := int64(0)
	for i := int64(0); i < seqlen; {
		val := int64(math.Pow( float64(2*prime), 4)) + 1	// (2n)^4 + 1
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
 * Date: December 07, 2021
 * Link: https://oeis.org/A000062
 */
 func A000062(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	for i := int64(1); i <= seqlen; i++ {
		a[i-1] = int64(math.Floor(float64(i)/(float64(math.E)-2)))
	}
	return a, 1
}

/**
 * A000064 generates the partial sums of A000008
 * Date: December 07, 2021
 * Link: https://oeis.org/A000064
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
 * Date: December 07, 2021
 * Link: https://oeis.org/A000065
 */
 func A000065(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	a10, _ := A000010(seqlen)
	for i := int64(0); i < seqlen; i++ {
		a[i] = a10[i] - 1
	}
	return a, 0
}

/**
 * A000068 returns a sequence such that n^4 + 1 is prime.
 * Date: December 07, 2021
 * Link: https://oeis.org/A000068
 */
 func A000068(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	prime := int64(0)
	for i := int64(0); i < seqlen; {
		val := int64(math.Pow( float64(prime), 4)) + 1	// (2n)^4 + 1
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
 * Date: December 07, 2021
 * Link: https://oeis.org/A000069
 */
 func A000069(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	num := int64(0)
	for i := int64(0); i < seqlen; {
		binary := strconv.FormatInt(num, 2)	// convert to binary
		count := int64(0)
		for _, bit := range binary {
			if bit == '1' {
				count++
			}
		}
		if count % 2 != 0 {
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
 * Date: December 07, 2021
 * Link: https://oeis.org/A000070
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
 * Date: December 07, 2021
 * Link: https://oeis.org/A000071
 */
 func A000071(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	F, _ := A000045(seqlen+1)
	for i := int64(1); i <= seqlen; i++ {
		a[i-1] = F[i] - 1
	}
	return a, 1
}

/**
 * A000073: Tribonacci #s: a[n] = a[n-1] + a[n-2] + a[n-3]
 *	a[0] = a[1] = 0 and a[2] = 1
 * Date: December 07, 2021
 * Link: https://oeis.org/A000073
 */
func A000073(seqlen int64) ([]int64, int64) {
	if seqlen > OVERFLOW_A000073 {
		utils.OverflowError("A000073", OVERFLOW_A000073)
	}

	a := make([]int64, seqlen)
	a[0], a[1], a[2] = 0, 0, 1
	for i := int64(3); i < seqlen; i++ {
		a[i] = a[i-1] + a[i-2] + a[i-3]
	}
	return a, 0
}

