// ============================================================================
// = thru300.go
// = 	Description		All OEIS sequences from A000201-A000300
// = 	Note			Not all sequences in this range have been programmed
// = 	Date			December 12, 2021
// ============================================================================

package seq

import (
	"OEIS/utils"
	"fmt"
	"math"
	"math/big"
)

const (
	LONG_A000205 = 10
)

/**
 * A000201 computes the Lower Wythoff sequence (a Beatty sequence):
 *  a(n) = floor(n*phi), where phi = (1+sqrt(5))/2 = A001622
 * Date: December 12, 2021
 * Link: https://oeis.org/A000201
 */
func A000201(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	phi := (1.0 + math.Sqrt(5)) / 2.0
	for n := int64(1); n <= seqlen; n++ {
		a[n-1] = int64(math.Floor(float64(n) * phi))
	}
	return a, 1
}

/**
 * A000202 computes a(8i+j) = 13i + a(j), where 1<=j<=8. 
 * Date: December 12, 2021
 * Link: https://oeis.org/A000202
 */
func A000202(seqlen int64) ([]int64, int64) {
	b := []int64{1, 3, 4, 6, 8, 9, 11, 12}
	a := utils.InitIslice(seqlen, b)
	
	if seqlen < 8 {
		return a, 1
	}
	
	foo := int64(8)
	for i := int64(1); i < seqlen; i++ {
		for j := int64(1); j <= foo && foo*i+j <= seqlen; j++ {
			// -1 b/c offset = 1
			a[foo * i + j-1] = 13 * i + a[j-1]
		}
	}
	return a, 1
}

/**
 * A000203 computes a(n) = sigma(n), the sum of the divisors of n
 * Also, sigma_1(n)
 * Date: December 12, 2021
 * Link: https://oeis.org/A000203
 */
func A000203(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	for i := int64(1); i <= seqlen; i++ {
		a[i-1] = utils.Sum(utils.Factors(i))
	}
	return a, 1
}

/**
 * A000204 computes the Lucas #s beginning with 1, that is:
 *  L(n) = L(n-1) + L(n-2), with L(1) = 1, L(2) = 3
 * Date: December 12, 2021 	
 * Link: https://oeis.org/A000204
 */
func A000204(seqlen int64) ([]*big.Int, int64) {
	a, _ := A000032(seqlen+1)
	return a[1:], 1
}

/**
 * A000205 computes the # of positive integers <= 2^n of the form x^2 + 3y^2
 * Date: December 12, 2021 	
 * Link: https://oeis.org/A000205
 */
func A000205(seqlen int64) ([]*big.Int, int64) {
	if seqlen > LONG_A000205 {
		utils.LongCalculationWarning("A000205", LONG_A000205)
	}

	a := utils.Repr(seqlen, 1, 3, 1)
	return a, 0
}

/**
 * A000207 returns the # of inequivalent ways of dissecting a regular (n+2)-gon
 *  into n triangles by n-1 non-intersecting diagonals under rotations and 
 *  reflections; also the number of planar 2-trees.
 * NOTE: this has some odd rounding errors
 * Date: December 13, 2021	Accurate as of: 
 * Link: https://oeis.org/A000207
 */
func A000207(seqlen int64) ([]*big.Int, int64) {
	utils.AccuracyWarning("A000207")

	a := utils.InitBslice(seqlen, 
		[]*big.Int{inew(1), inew(1)})
	
	C, _ := A000108(seqlen+3)		// catalan numbers
	C = utils.ShiftBig(C, 2)	// C(n)=A000108(n-2)
	fmt.Println(C)
	for n := int64(3); n <= seqlen+2; n++ {
		k := (n+1)/2		// n is odd
		if n % 2 == 0 {		// n is even
			k = n/2+1
		}
		// a(n) = C(n)/(2*n) + C(n/2+1)/4 + C(k)/2 + C(n/3+1)/3
		// where C(n) = A000108(n-2)
		part1 := fadd(
			fdiv(tofloat(C[n]), fnew(2.0*float64(n))), 
			fdiv(tofloat(C[n/2+1]), fnew(4)))
		part2 := fadd(
			fdiv(tofloat(C[k]), fnew(2)), 
			fdiv(tofloat(C[n/3+1]), fnew(3)))
		a[n-3] = floor(fadd(part1, part2))
	}
	return a, 1
}

/**
 * A000208 computes the # of even sequences with period 2n
 *  Somehow this sequence is different than A000206
 * Date: December 14, 2021
 * Link: https://oeis.org/A000208
 */
func A000208(seqlen int64) ([]*big.Int, int64) {
	utils.AccuracyWarning("A000208 (which computes A000013)")

	a := utils.CreateSlice(seqlen)
	a13, _ := A000013(seqlen*2)
	for i := int64(0); i < seqlen; i++ {
		if i % 2 == 0 {
			a[i] = div(add(a13[2*i], a13[i]), inew(2))
		} else {
			a[i] = div(a13[2*i], inew(2))
		}
	}
	return a, 0
}

/**
 * A000209 computes the nearest integer to tan n
 * Date: December 14, 2021
 * Link: https://oeis.org/A000209
 */
func A000209(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	for i := int64(0); i < seqlen; i++ {
		a[i] = int64(math.Round(math.Tan(float64(i))))
	}
	return a, 0
}

/**
 * A000210 computes a Beatty sequence: floor(n*(e-1))
 * Date: December 14, 2021
 * Link: https://oeis.org/A000210
 */
func A000210(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	for i := int64(1); i <= seqlen; i++ {
		a[i-1] = int64(math.Floor(float64(i) * (math.E-1.0)))
	}
	return a, 1
}

/**
 * A000211 computes a(n) = a(n-1) + a(n-2) - 2, a(0) = 4, a(1) = 3. 
 * Date: December 14, 2021
 * Link: https://oeis.org/A000211
 */
func A000211(seqlen int64) ([]*big.Int, int64) {
	a := utils.InitBslice(seqlen, 
		[]*big.Int{inew(4), inew(3)})
	
	for i := int64(2); i < seqlen; i++ {
		a[i] = sub(add(a[i-1], a[i-2]), inew(2))
	}
	return a, 0
}

/**
 * A000212 computes a(n) = floor(n^2/3)
 * Date: December 14, 2021
 * Link: https://oeis.org/A000212
 */
func A000212(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	for i := int64(0); i < seqlen; i++ {
		a[i] = int64(math.Floor( math.Pow(float64(i), 2)/3.0 ))
	}
	return a, 0
}

/**
 * A000213 computes Tribonacci #s: a(n) = a(n-1) + a(n-2) + a(n-3)
 *  with a(0)=a(1)=a(2)=1. 
 * Date: December 14, 2021
 * Link: https://oeis.org/A000213
 */
func A000213(seqlen int64) ([]*big.Int, int64) {
	a := utils.InitBslice(seqlen, 
		[]*big.Int{inew(1), inew(1), inew(1)})
	
	for i := int64(3); i < seqlen; i++ {
		a[i] = add(add(a[i-1], a[i-2]), a[i-3])
	}
	return a, 0
}

/**
 * A000215 computes the Fermat #s: a(n)=2^(2^n)+1
 * Date: December 14, 2021
 * Link: https://oeis.org/A000215
 */
func A000215(seqlen int64) ([]*big.Int, int64) {
	a := utils.CreateSlice(seqlen)
	for i := int64(0); i < seqlen; i++ {
		a[i] = add(pow(inew(2), pow(inew(2), inew(i))), inew(1))
	}
	return a, 0
}

/**
 * A000216 computes the sum of squares of digits of previous term, starting with 2. 
 * Date: December 14, 2021
 * Link: https://oeis.org/A000216
 */
func A000216(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	a[0] = 2
	for n := int64(1); n < seqlen; n++ {
		a[n] += utils.SumSquares(a[n-1])
	}
	return a, 1
}

/**
 * A000217 computes the triangle numbers (0+1+2+...+n)
 * Date: December 14, 2021
 * Link: https://oeis.org/A000217
 */
func A000217(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	for i := int64(0); i < seqlen; i++ {
		a[i] = i*(i+1)/2
	}
	return a, 0
}

/**
 * A000218 computes the sum of squares of digits of previous term, starting with 3
 * Date: December 14, 2021
 * Link: https://oeis.org/A000218
 */
 func A000218(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	a[0] = 3
	for n := int64(1); n < seqlen; n++ {
		a[n] += utils.SumSquares(a[n-1])
	}
	return a, 1
}

/**
 * A000219 computes the # of planar partitions (or plane partitions) of n. 
 * Date: December 14, 2021
 * Link: https://oeis.org/A000219
 */
func A000219(seqlen int64) ([]*big.Int, int64) {
	a := utils.CreateSlice(seqlen)
	a[0] = inew(1)
	for n := int64(1); n < seqlen; n++ {
		inv := fdiv(fnew(1), fnew(float64(n)))
		sum := fnew(0)
		for k := int64(1); k <= n; k++ {
			sum = fadd(sum, tofloat(mul(a[n-k], utils.Sigma(k, 2))))
		}
		a[n] = floor(fmul(inv, sum))
	}
	return a, 0
}

/**
 * A000221 computes the sum of squares of digits of previous term, starting with 5
 * Date: December 14, 2021
 * Link: https://oeis.org/A000221
 */
func A000221(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	a[0] = 5
	for n := int64(1); n < seqlen; n++ {
		a[n] += utils.SumSquares(a[n-1])
	}
	return a, 1
}

/**
 * A000225 computes a(n) = 2^n - 1
 * Date: December 14, 2021
 * Link: https://oeis.org/A000225
 */
func A000225(seqlen int64) ([]*big.Int, int64) {
	a := utils.CreateSlice(seqlen)
	sq, _ := A000079(seqlen)
	for n := int64(1); n < seqlen; n++ {
		a[n] = sub(sq[n], inew(1))
	}
	return a, 0
}

/**
 * A000227 computes the nearest integer to e^n
 * Date: December 14, 2021
 * Link: https://oeis.org/A000227
 */
func A000227(seqlen int64) ([]*big.Int, int64) {
	a := utils.CreateSlice(seqlen)
	for i := int64(0); i < seqlen; i++ {
		a[i] = round(fpow(fnew(math.E), i))
	}
	return a, 0
}

/**
 * A000230 computes the smallest prime p such that there is a gap
 * of exactly 2n between p and the next prime, or -1 if no such prime exists
 * Date: December 14, 2021
 * Link: https://oeis.org/A000230
 */
func A000230(seqlen int64) ([]*big.Int, int64) {
	a := utils.CreateSlice(seqlen)
	a[0] = inew(2)

	primes := utils.PrimesBig(seqlen)
	lenfactor := 25		// how many times more primes to compute

	// loop to compute a
	for n := int64(1); n < seqlen; {
		oldn := n
		for k := 0; k < len(primes)-1; k++ {
			test := add(primes[k], mul(inew(2), inew(n)))
			if equals(test, primes[k+1]) {
				a[n] = primes[k]
				n++
				k = len(primes)+1
				//fmt.Println("for n =", n, "primes computed = ", len(primes), "\nwith", primes[len(primes)-1], "as the last prime")
			}
		}

		// this will only be true if the # of primes generated wasn't enough
		// and the value of a[n] wasn't found
		if oldn == n {
			primes = utils.PrimesBig(int64(len(primes)*lenfactor))
		}
	}
	return a, 0
}

/**
 * A000231 computes the # of inequivalent Boolean functions of n
 *  variables under action of complementing group.
 * Date: December 14, 2021
 * Link: https://oeis.org/A000231
 */
func A000231(seqlen int64) ([]*big.Int, int64) {
	a := utils.CreateSlice(seqlen)
	for n := int64(1); n <= seqlen; n++ {
		// a(n) = (2^(2^n)+(2^n-1)*2^(2^(n-1)))/2^n
		pt1 := pow(inew(2), pow(inew(2), inew(n)))	// 2^(2^n)
		pt2 := sub(pow(inew(2), inew(n)), inew(1))	// 2^n-1
		pt3 := pow(inew(2), pow(inew(2), inew(n-1)))	// 2^(2^(n-1))
		numer := add(pt1, mul(pt2, pt3))	// (2^(2^n)+(2^n-1)*2^(2^(n-1)))
		a[n-1] = div(numer, pow(inew(2), inew(n)))
	}
	return a, 1
}

/**
 * A000240 computes the Rencontres #s: # of permutations of [n] with
 *  exactly one fixed point. 
 * Date: December 14, 2021
 * Link: https://oeis.org/A000240
 */
func A000240(seqlen int64) ([]*big.Int, int64) {
	a := utils.CreateSlice(seqlen)
	for n := int64(1); n <= seqlen; n++ {
		for k := int64(0); k < n; k++ {
			a[n-1] = add(a[n-1], 
				mul(
					pow(inew(-1), inew(k)), 
					div(utils.Fact(inew(n)), utils.Fact(inew(k)))))
		}
	}
	return a, 1
}

/**
 * A000244 computes the powers of 3
 * Date: December 14, 2021
 * Link: https://oeis.org/A000244
 */
func A000244(seqlen int64) ([]*big.Int, int64) {
	a := utils.CreateSlice(seqlen)
	for n := int64(0); n < seqlen; n++ {
		a[n] = pow(inew(3), inew(n))
	}
	return a, 0
}

/**
 * A000245 computes a(n) = 3*(2*n)!/((n+2)!*(n-1)!). 
 * Date: December 14, 2021
 * Link: https://oeis.org/A000245
 */
func A000245(seqlen int64) ([]*big.Int, int64) {
	a := utils.CreateSlice(seqlen)
	for n := int64(1); n < seqlen; n++ {
		numer := mul(inew(3), utils.Fact(inew(2*n)))
		denom := mul(utils.Fact(inew(n+2)), utils.Fact(inew(n-1)))
		a[n] = div(numer, denom)
	}
	return a, 0
}

/**
 * A000246 computes the # of permutations in the symmetric group S_n that have odd order. 
 * Date: December 14, 2021
 * Link: https://oeis.org/A000246
 */
func A000246(seqlen int64) ([]*big.Int, int64) {
	a := utils.CreateSlice(seqlen)
	a[0] = inew(1)
	for n := int64(1); n < seqlen; n++ {
		//a(n) = Sum_{k=0..floor((n-1)/2)} (2k)! * C(n-1, 2k) * a(n-2k-1)
		sum := inew(0)
		for k := int64(0); k <= (n-1)/2; k++ {
			sum = add(sum, mul(mul(
					utils.Fact(inew(2*k)),
					utils.C(inew(n-1),inew(2*k))), 
					a[n-2*k-1]))		
		}
		a[n] = sum
	}
	return a, 0
}

/**
 * A000247 computes a(n) = 2^n-n-2
 * Date: December 14, 2021
 * Link: https://oeis.org/A000247
 */
func A000247(seqlen int64) ([]*big.Int, int64) {
	a := utils.CreateSlice(seqlen)
	for i := int64(2); i <= seqlen+1; i++ {
		a[i-2] = sub(pow(inew(2), inew(i)), inew(i+2))
	}
	return a, 2
}

/**
 * A000248 computes the expansion of e.g.f. exp(x*exp(x)). 
 * Date: December 14, 2021
 * Link: https://oeis.org/A000248
 */
func A000248(seqlen int64) ([]*big.Int, int64) {
	// a(n) = Sum_{k=0..n} C(n,k)*(n-k)^k
	a := utils.CreateSlice(seqlen)
	for n := int64(0); n < seqlen; n++ {
		sum := inew(0)
		for k := int64(0); k <= n; k++ {
			sum = add(sum, mul(utils.C(inew(n), inew(k)), pow(inew(n-k), inew(k))))
		}
		a[n] = sum
	}
	return a, 0
}

/**
 * A000252 computes # of invertible 2 X 2 matrices mod n. 
 * Date: December 14, 2021
 * Link: https://oeis.org/A000252
 */
func A000252(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	for n := int64(1); n <= seqlen; n++ {
		prod := float64(1)
		for k := int64(1); k <= n; k++ {
			// primes dividing n
			if utils.IsPrime(k) && n % k == 0 {
				//(1 - 1/p^2)*(1 - 1/p)
				prod *= (1.0 - 1.0/math.Pow(float64(k), 2)) * (1.0 - 1.0/float64(k))
			}
		}
		a[n-1] = int64(math.Floor(math.Pow(float64(n), 4.0) * prod))
	}
	return a, 1
}

/**
 * A000253 computes a(n) = 2*a(n-1) - a(n-2) + a(n-3) + 2^(n-1). 
 * Date: December 14, 2021
 * Link: https://oeis.org/A000253
 */
func A000253(seqlen int64) ([]*big.Int, int64) {
	a := utils.InitBslice(seqlen, 
		[]*big.Int{inew(0), inew(1), inew(4), inew(11)})

	for n := int64(3); n < seqlen; n++ {
		a[n] = add(add(sub(mul(inew(2), a[n-1]), a[n-2]), a[n-3]), pow(inew(2), inew(n-1)))
	}

	return a, 0
}

/**
 * A000254 computes unsigned Stirling numbers of first kind, s(n+1,2):
 *  a(n+1) = (n+1)*a(n) + n!. 
 * Date: December 14, 2021
 * Link: https://oeis.org/A000254
 */
func A000254(seqlen int64) ([]*big.Int, int64) {
	a := utils.CreateSlice(seqlen)
	for n := int64(1); n < seqlen; n++ {
		a[n] = add(mul(inew(n), a[n-1]), utils.Fact(inew(n-1)))
	}
	return a, 0
}

/**
 * A000255 computes a(n) = n*a(n-1) + (n-1)*a(n-2), a(0) = 1, a(1) = 1
 * Date: December 14, 2021
 * Link: https://oeis.org/A000255
 */
func A000255(seqlen int64) ([]*big.Int, int64) {
	a := utils.InitBslice(seqlen, 
		[]*big.Int{inew(1), inew(1)})
	
	for n := int64(2); n < seqlen; n++ {
		a[n] = add(mul(inew(n), a[n-1]), mul(inew(n-1), a[n-2]))
	}
	return a, 0
}

/**
 * A000256 computes the # of simple triangulations of the plane with n nodes. 
 * Date: December 14, 2021
 * Link: https://oeis.org/A000256
 */
func A000256(seqlen int64) ([]*big.Int, int64) {
	a := utils.CreateSlice(seqlen)
	bound := int64(3)
	if seqlen <= 3 {
		bound = seqlen
	}

	// init
	init := []int64{1, 1, 0, 1}
	for i := int64(0); i < bound; i++ {
		a[i] = inew(init[i])
	}
	
	// loop to generate
	for n := int64(5); n <= seqlen+2; n++ {
		// a(n) = (1/4)*(7*binomial(3n-9, n-4)-(8*n^2-43n+57)*a(n-1)) / (8*n^2-51n+81), n>4
		frac := fdiv(fnew(1), fnew(4))
		bino := mul(inew(7), utils.C(inew(3*n-9), inew(n-4)))
		poly := add(sub(mul(inew(8), pow(inew(n), inew(2))), inew(43*n)), inew(57))
		mess := sub(bino, mul(poly, a[n-1-3]))
		numer := fmul(frac, tofloat(mess))
		denom := add(sub(mul(inew(8), pow(inew(n), inew(2))), inew(51*n)), inew(81))
		a[n-3] = floor(fdiv(numer, tofloat(denom)))
	}
	return a, 3
}

/**
 * A000257 computes the # of rooted bicubic maps: a(n) = (8n-4)*a(n-1)/(n+2)
 * Date: December 14, 2021
 * Link: https://oeis.org/A000257
 */
func A000257(seqlen int64) ([]*big.Int, int64) {
	a := utils.CreateSlice(seqlen)
	a[0] = inew(1)
	for n := int64(1); n < seqlen; n++ {
		a[n] = div(mul(inew(8*n-4), a[n-1]), inew(n+2))
	}
	return a, 0
}

/**
 * A000259 computes the # of certain rooted planar maps
 * Date: December 14, 2021
 * Link: https://oeis.org/A000259
 */
func A000259(seqlen int64) ([]*big.Int, int64) {
	a := utils.CreateSlice(seqlen)
	F, _ := A000045(seqlen)

	// a(n) = Sum_{k = 1..n} (-1)^(k-1)*C(3n, n-k)*k/n*F(k-2)
	for n := int64(1); n <= seqlen; n++ {
		sum := fnew(0)
		for k := int64(1); k <= n; k++ {
			m1 := fpow(fnew(-1), k-1)
			m2 := utils.C(inew(3*n), inew(n-k))
			m3 := fdiv(fnew(float64(k)), fnew(float64(n)))
			m4 := inew(1)
			if k-2 != -1 {
				m4 = F[k-2]
			}
			m := fmul(m1, fmul(tofloat(m2), fmul(m3, tofloat(m4))))
			sum = fadd(sum, m)
		}
		a[n-1] = floor(sum)
	}
	return a, 1
}

/**
 * A000260 computes the # of rooted simplicial 3-polytopes with n+3
 *  nodes; or rooted 3-connected triangulations with 2n+2 faces; or
 *  rooted 3-connected trivalent maps with 2n+2 vertices. 
 * Date: December 14, 2021
 * Link: https://oeis.org/A000260
 */
func A000260(seqlen int64) ([]*big.Int, int64) {
	a := utils.CreateSlice(seqlen)
	a[0] = inew(1)
	for n := int64(1); n < seqlen; n++ {
		a[n] = sub(utils.C(inew(4*n+1), inew(n+1)), mul(inew(9), utils.C(inew(4*n+1), inew(n-1))))
	}
	return a, 0
}

/**
 * A000261 computes a(n) = n*a(n-1) + (n-3)*a(n-2), with a(1) = 0, a(2) = 1. 
 * Date: December 14, 2021
 * Link: https://oeis.org/A000261
 */
func A000261(seqlen int64) ([]*big.Int, int64) {
	a := utils.InitBslice(seqlen, 
		[]*big.Int{inew(0), inew(1)})
	
	for n := int64(3); n <= seqlen; n++ {
		// a(n) = n*a(n-1) + (n-3)*a(n-2), with a(1) = 0, a(2) = 1. 
		a[n-1] = add(mul(inew(n), a[n-2]), mul(inew(n-3), a[n-3]))
	}
	return a, 1
}

/**
 * A000262 computes the # of "sets of lists": # of partitions of 
 *  {1,...,n} into any number of lists, where a list means an 
 *  ordered subset. 
 * Date: December 14, 2021
 * Link: https://oeis.org/A000262
 */
func A000262(seqlen int64) ([]*big.Int, int64) {
	a := utils.CreateSlice(seqlen)
	a[0] = inew(1)
	for n := int64(1); n < seqlen; n++ {
		// a(n) = (n-1)! * Sum_{k=1..n} (a(n-k)*k!)/((n-k)!*(k-1)!)
		sum := fnew(0)
		for k := int64(1); k <= n; k++ {
			num := mul(a[n-k], utils.Fact(inew(k)))
			den := mul(utils.Fact(inew(n-k)), utils.Fact(inew(k-1)))
			sum = fadd(sum, fdiv(tofloat(num), tofloat(den)))
		}
		a[n] = floor(fmul(tofloat(utils.Fact(inew(n-1))), sum))
	}
	return a, 0
}

/**
 * A000263 computes the # of partitions into non-integral powers
 * Date: December 14, 2021
 * Link: https://oeis.org/A000263
 */
func A000263(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	for n := int64(3); n <= seqlen+2; n++ {
		nf := float64(n)
		count := int64(0)
		for x1 := int64(1); x1 <= n*n; x1++ {
			x1f := float64(x1)
			for x2 := x1+1; x2 <= n*n; x2++ {
				x2f := float64(x2)
				if math.Sqrt(x1f) + math.Sqrt(x2f) <= nf {
					count++
				}
			}
		}
		a[n-3] = count
	}
	return a, 3
}

/**
 * A000265 computes the largest odd divisor of n
 * Date: December 14, 2021
 * Link: https://oeis.org/A000265
 */
func A000265(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	for n := int64(1); n <= seqlen; n++ {
		if n % 2 != 0 {
			a[n-1] = n
		} else {
			a[n-1] = a[n/2-1]
		}
	}
	return a, 1
}

/**
 * A000266 computes the expansion of e.g.f exp(-x^2/2)/(1-x)
 * Date: December 14, 2021
 * Link: https://oeis.org/A000266
 */
func A000266(seqlen int64) ([]*big.Int, int64) {
	a := utils.CreateSlice(seqlen)
	for n := int64(0); n < seqlen; n++ {
		//a(n) = n! * Sum_{i=0..floor(n/2)} (-1)^i /(i! * 2^i)
		sum := fnew(0)
		for i := int64(0); i <= n/2; i++ {
			num := fpow(fnew(-1), i)
			den := fmul(tofloat(utils.Fact(inew(i))), fpow(fnew(2), i))
			sum = fadd(sum, fdiv(num, den))
		}
		a[n] = floor(fmul(tofloat(utils.Fact(inew(n))), sum))
	}
	return a, 0
}

/**
 * A000267 computes the integer part of square root of 4n+1
 * Date: December 14, 2021
 * Link: https://oeis.org/A000267
 */
func A000267(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	for n := int64(0); n < seqlen; n++ {
		a[n] = utils.Isqrt(4*n+1)
	}
	return a, 0
}

/**
 * A000270 computes: For n >= 2, a(n) = b(n+1)+b(n)+b(n-1), where
 *  b(i) are the ménage numbers A000179; a(0)=a(1)=1
 * Date: December 14, 2021
 * Link: https://oeis.org/A000270
 */
func A000270(seqlen int64) ([]*big.Int, int64) {
	a := utils.InitBslice(seqlen, 
		[]*big.Int{inew(1), inew(1)})
	
	b, _ := A000179(seqlen+1)
	for n := int64(2); n < seqlen; n++ {
		a[n] = add(add(b[n+1], b[n]), b[n-1])
	}
	return a, 0
}


/*
 * A000271 computes the sums of ménage #s
 * Date: December 14, 2021
 * Link: https://oeis.org/A000271
 */
func A000271(seqlen int64) ([]*big.Int, int64) {
	a := utils.CreateSlice(seqlen)
	for n := int64(0); n < seqlen; n++ {
		// a(n) = Sum_{k=0..n} (-1)^(n-k)*binomial(n+k,2*k)*k!
		sum := inew(0)
		for k := int64(0); k <= n; k++ {
			p1 := pow(inew(-1), inew(n-k))
			p2 := utils.C(inew(n+k), inew(2*k))
			p3 := utils.Fact(inew(k))
			eqn := mul(mul(p1, p2), p3)
			sum = add(sum, eqn)
		}
		a[n] = sum
	}
	return a, 0
}

/**
 * A000272 computes the # of trees on n labeled nodes: n^(n-2) with a(0)=1. 
 * Date: December 14, 2021
 * Link: https://oeis.org/A000272
 */
func A000272(seqlen int64) ([]*big.Int, int64) {
	a := utils.InitBslice(seqlen, 
		[]*big.Int{inew(1), inew(1)})
		
	for n := int64(1); n < seqlen-1; n++ {
		// a(n+1)= Sum_{i=1..n} n^(n-i)*binomial(n-1,i-1)
		sum := inew(0)
		for i := int64(1); i <= n; i++ {
			p1 := pow(inew(n), inew(n-i))		// n^(n-1-i)
			p2 := utils.C(inew(n-1), inew(i-1))	// binomial(n, i)
			sum = add(sum, mul(p1, p2))
		}
		a[n+1] = sum
	}
	return a, 0
}

/**
 * A000274 computes the # of permutations of length n with 2 
 *  consecutive ascending pairs.
 * Date: December 14, 2021
 * Link: https://oeis.org/A000274
 */
func A000274(seqlen int64) ([]*big.Int, int64) {
	a := utils.CreateSlice(seqlen)
	a166, _ := A000166(seqlen)
	for n := int64(2); n < seqlen; n++ {
		if n % 2 == 0 {
			a[n] = mul(a166[n], inew((n+1)/2))
		} else {
			a[n] = div(mul(a166[n], inew(n)), inew(2))
		}	
	}
	return a, 1
}

/** 
 * A000275 computes the coefficients of a Bessel function (reciprocal
 *  of J_0(z)); also pairs of permutations with rise/rise forbidden. 
 * Date: December 14, 2021
 * Link: https://oeis.org/A000275
 */
func A000275(seqlen int64) ([]*big.Int, int64) {
	// a(n) = Sum_{r=0..n-1} (-1)^(r+n+1) * binomial(n, r)^2 * a(r)
	a := utils.CreateSlice(seqlen)
	a[0] = inew(1)
	for n := int64(1); n < seqlen; n++ {
		sum := inew(0)
		for r := int64(0); r < n; r++ {
			t1 := pow(inew(-1), inew(r+n+1))
			t2 := pow(utils.C(inew(n), inew(r)), inew(2))
			t := mul(mul(t1, t2), a[r])
			sum = add(sum, t)
		}
		a[n] = sum
	}
	return a, 0
}

/**
 * A000276 computes associated Stirling numbers. 
 * Date: December 14, 2021
 * Link: https://oeis.org/A000276
 */
func A000276(seqlen int64) ([]*big.Int, int64) {
	a := utils.CreateSlice(seqlen)
	a254, _ := A000254(seqlen+3)
	for n := int64(4); n <= seqlen+3; n++ {
		a[n-4] = sub(sub(a254[n-1], utils.Fact(inew(n-1))), utils.Fact(inew(n-2)))
	}
	return a, 4
}

/**
 * A000277 computes 3*n - 2*floor(sqrt(4*n+5)) + 5
 * Date: December 14, 2021
 * Link: https://oeis.org/A000277
 */
func A000277(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	for n := int64(0); n < seqlen; n++ {
		a[n] = 3*n - 2*utils.Isqrt(4*n+5) + 5
	}
	return a, 0
}

/**
 * A000278 computes a(n) = a(n-1) + a(n-2)^2 for n >= 2
 *  with a(0) = 0 and a(1) = 1
 * Date: December 14, 2021
 * Link: https://oeis.org/A000278
 */
func A000278(seqlen int64) ([]*big.Int, int64) {
	a := utils.InitBslice(seqlen, 
		[]*big.Int{inew(0), inew(1)})
	
	for n := int64(2); n < seqlen; n++ {
		a[n] = add(a[n-1], pow(a[n-2], inew(2)))
	}
	return a, 0
}

/**
 * A000279 computes card matching: coefficients B[n,1] of t in the
 *  reduced hit polynomial A[n,n,n](t). 
 * Date: December 14, 2021
 * Link: https://oeis.org/A000279
 */
func A000279(seqlen int64) ([]*big.Int, int64) {
	a := utils.CreateSlice(seqlen)
	a172, _ := A000172(seqlen+1)
	for n := int64(1); n <= seqlen; n++ {
		//a(n) = n^2*(A000172(n)+4*A000172(n-1))/(n+1)
		num := add(a172[n], mul(inew(4), a172[n-1]))
		num = mul(pow(inew(n), inew(2)), num)
		a[n-1] = floor(fdiv(tofloat(num), fnew(float64(n+1))))
	}
	return a, 1
}

/**
 * A000280 computes a(n) = a(n-1) + a(n-2)^3. 
 * Date: December 14, 2021
 * Link: https://oeis.org/A000280
 */
func A000280(seqlen int64) ([]*big.Int, int64) {
	a := utils.InitBslice(seqlen, 
		[]*big.Int{inew(0), inew(1)})
	
	for n := int64(2); n < seqlen; n++ {
		a[n] = add(a[n-1], pow(a[n-2], inew(3)))
	}
	return a, 0
}

/**
 * A000283 computes a(n) = a(n-1)^2 + a(n-2)^2 for n >= 2 
 *  with a(0) = 0 and a(1) = 1
 * Date: December 14, 2021
 * Link: https://oeis.org/A000283
 */
 func A000283(seqlen int64) ([]*big.Int, int64) {
	a := utils.InitBslice(seqlen, 
		[]*big.Int{inew(0), inew(1)})
	
	for n := int64(2); n < seqlen; n++ {
		a[n] = add(pow(a[n-1], inew(2)), pow(a[n-2], inew(2)))
	}
	return a, 0
}

/**
 * A000284 computes a(n) = a(n-1)^3 + a(n-2)
 *  with a(0)=0, a(1)=1
 * Date: December 14, 2021
 * Link: https://oeis.org/A000284
 */
 func A000284(seqlen int64) ([]*big.Int, int64) {
	a := utils.InitBslice(seqlen, 
		[]*big.Int{inew(0), inew(1)})
	
	for n := int64(2); n < seqlen; n++ {
		a[n] = add(a[n-2], pow(a[n-1], inew(3)))
	}
	return a, 0
}

/**
 * A000285 computes a(0) = 1, a(1) = 4, and 
 *  a(n) = a(n-1) + a(n-2) for n >= 2. 
 * Date: December 14, 2021
 * Link: https://oeis.org/A000285
 */
func A000285(seqlen int64) ([]*big.Int, int64) {
	a := utils.InitBslice(seqlen, 
		[]*big.Int{inew(1), inew(4)})

	// loop to generate
	for n := int64(2); n < seqlen; n++ {
		a[n] = add(a[n-1], a[n-2])
	}
	return a, 0
}

/**
 * A000286 computes the # of positive integers <= 2^n of form 2 x^2 + 5 y^2
 * Date: December 15, 2021
 * Link: https://oeis.org/A000286
 */
func A000286(seqlen int64) ([]*big.Int, int64) {
	a := utils.Repr(seqlen, 2, 5, 0)
	return a, 0
}

/**
 * A000287 computes the # of rooted polyhedral graphs with n edges
 * Date: December 15, 2021
 * Link: https://oeis.org/A000287
 */
func A000287(seqlen int64) ([]*big.Int, int64) {
	// b(n) = ( 2*(2*n)!/(n!)^2 - (27*n^2+9*n-2)*b(n-1) ) / (54*n^2-90*n+32)
	b := utils.CreateSlice(seqlen+10)
	b[0] = inew(2)
	for n := 1; n < len(b); n++ {
		ni := int64(n) + 3
		nf := float64(n) + 3
		num1 := fmul(fnew(2), tofloat(utils.Fact(inew(2*ni))))	// 2*(2n)!
		den1 := fpow(tofloat(utils.Fact(inew(ni))), 2)		// (n!)^2
		frac1 := fdiv(num1, den1)							// 2*(2n)!/(n!)^2
		poly := fsub(fadd(fmul(fnew(27), fpow(fnew(nf), 2)), fmul(fnew(9), fnew(nf))), fnew(2))
		num2 := fsub(frac1, fmul(poly, tofloat(b[n-1])))	
		den2 := fadd(fsub(fmul(fnew(54), fpow(fnew(nf), 2)), fmul(fnew(90), fnew(nf))), fnew(32))
		b[n] = floor(fdiv(num2, den2))
	}

	// a(n) = b(n-1) + 2*(-1)^n
	a := utils.InitBslice(seqlen, []*big.Int{inew(1), inew(0), inew(4), inew(6)})
	for n := int64(9); n <= seqlen+4; n++ {
		fmt.Println("using", b[n-3], "for n =", n)
		a[n-5] = add(b[n-3], mul(inew(2), pow(inew(-1), inew(n-5))))
	}

	return a, 6
}

/**
 * A000288 compute the Tetranacci #s: a(n) = a(n-1) + a(n-2) + a(n-3) + a(n-4)
 *  w/ a(0) = a(1) = a(2) = a(3) = 1. 
 * Date: December 15, 2021
 * Link: https://oeis.org/A000288
 */
func A000288(seqlen int64) ([]*big.Int, int64) {
	a := utils.InitBslice(seqlen, []*big.Int{inew(1), inew(1), inew(1), inew(1)})
	for i := int64(4); i < seqlen; i++ {
		a[i].Add(a[i-1], a[i-2])
		a[i].Add(a[i], a[i-3])
		a[i].Add(a[i], a[i-4])
	}
	return a, 0
}

/**
 * A000289 computes a nonlinear recurrence: a(n) = a(n-1)^2 - 3*a(n-1) + 3 (for n>1). 
 * Date: December 15, 2021
 * Link: https://oeis.org/A000289
 */
func A000289(seqlen int64) ([]*big.Int, int64) {
	a := utils.InitBslice(seqlen, []*big.Int{inew(1), inew(4)})
	for n := int64(2); n < seqlen; n++ {
		a[n] = add(sub(pow(a[n-1], inew(2)), mul(inew(3), a[n-1])), inew(3))
	}
	return a, 0
}

/**
 * A000290 computes the squares; a[n] = n^2
 * Date: December 15, 2021
 * Link: https://oeis.org/A000290
 */
func A000290(seqlen int64) ([]*big.Int, int64) {
	a := utils.Power(seqlen, inew(2))
	return a, 0
}

/**
 * A000291 computes # of bipartite partitions of n white objects and 2 black ones
 * Date: December 15, 2021
 * Link: https://oeis.org/A000291
 */
func A000291(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	a70, _ := A000070(seqlen)
	a97, _ := A000097(seqlen)
	for n := int64(0); n < seqlen; n++ {
		a[n] = a70[n] + a97[n]
	}
	return a, 0
}

/**
 * A000292 computes tetrahedral (or triangular pyramidal) #s:
 *  a(n) = C(n+2,3) = n*(n+1)*(n+2)/6. 
 * Date: December 15, 2021
 * Link: https://oeis.org/A000292
 */
func A000292(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	for n := int64(0); n < seqlen; n++ {
		a[n] = n * (n+1) * (n+2)/6
	}
	return a, 0
}

/**
 * A000294 expansion of g.f. Product_{k >= 1} (1 - x^k)^(-k*(k+1)/2). 
 * Date: December 15, 2021
 * Link: https://oeis.org/A000294
 */
func A000294(seqlen int64) ([]*big.Int, int64) {
	// a(n) = (1/(2*n))*Sum_{k=1..n} (sigma[2](k)+sigma[3](k))*a(n-k)
	a := utils.CreateSlice(seqlen)
	a[0] = inew(1)
	for n := int64(1); n < seqlen; n++ {
		nf := float64(n)
		sum := inew(0)
		for k := int64(1); k <= n; k++ {
			sum = add(sum, mul(add(utils.Sigma(k, 2), utils.Sigma(k, 3)), a[n-k]))
		}
		a[n] = floor(fmul(tofloat(sum), fdiv(fnew(1), fnew(2.0*nf))))
	}
	return a, 0
}

/**
 * A000295 computes the Eulerian numbers (Euler's triangle: column k=2 of A008292, column k=1 of A173018). 
 * Date: December 15, 2021
 * Link: https://oeis.org/A000295
 */
func A000295(seqlen int64) ([]*big.Int, int64) {
	a := utils.CreateSlice(seqlen)
	for n := int64(0); n < seqlen; n++ {
		a[n] = sub(pow(inew(2), inew(n)), inew(n+1))
	}
	return a, 0
}

/**
 * A000296 computes set partitions without singletons: number of partitions of an n-set into blocks of size > 1. Also number of cyclically spaced (or feasible) partitions. 
 * Date: December 15, 2021
 * Link: https://oeis.org/A000296
 */
func A000296(seqlen int64) ([]*big.Int, int64) {
	a := utils.CreateSlice(seqlen)
	for n := int64(0); n < seqlen; n++ {
		ksum := fnew(0)
		for k := int64(0); k <= n; k++ {
			jsum := fnew(0)
			for j := int64(0); j <= k; j++ {
				jf := float64(j)
				p1 := fpow(fnew(-1), j)
				p2 := tofloat(utils.C(inew(k), inew(j)))
				p3 := fpow(fnew(1-jf), n)
				num := fmul(fmul(p1, p2), p3)
				jsum = fadd(jsum, fdiv(num, tofloat(utils.Fact(inew(k)))))
			}
			ksum = fadd(ksum, fmul(fpow(fnew(-1), n-k), jsum))
		}
		a[n] = floor(ksum)
	}
	return a, 0
}

/**
 * A000297 computes a(n) = (n+1)*(n+3)*(n+8)/6. 
 * Date: December 15, 2021
 * Link: https://oeis.org/
 */
func A000297(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	for n := int64(0); n < seqlen-1; n++ {
		a[n+1] = (n+1)*(n+3)*(n+8)/6
	}
	return a, -1
}

