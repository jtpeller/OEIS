// ============================================================================
// = calc.go
// = 	Description		Functions like divisors, GCD, etc.
// = 	Date 			December 07, 2021
// ============================================================================

package utils

import (
	"errors"
	"math"
	"math/big"

	gb "github.com/jtpeller/gobig"
)

// ####################### COMMON CALCULATIONS #########################
// ### this section contains any common calculations
// ### Examples: Summation, Product, Factorials etc.

// compute the factorial of a num (big.Int)
func Fact(num *big.Int) *big.Int {
	if gb.Less(num, gb.Zero()) {
		HandleError(errors.New("factorial of a negative number is undefined"))
	}

	prod := gb.New(1)
	for i := gb.New(1); i.Cmp(num) == -1 || i.Cmp(num) == 0; i.Add(i, gb.New(1)) {
		prod.Mul(prod, i)
	}
	return prod
}

// computes the factorial of num (int64)
func IFact(num int64) int64 {
	if num < 0 {
		HandleError(errors.New("factorial of a negative number is undefined"))
	}
	prod := int64(1)
	for i := int64(1); i < num; i++ {
		prod *= i
	}
	return prod
}

// calculate the integer log of a number & the given base
func ILog(num int64, base int64) int64 {
	return int64(math.Log(float64(num)) / math.Log(float64(base)))
}

// Calculate the integer square root of a number
func Isqrt(num int64) (int64) {
	return int64(math.Floor(math.Sqrt(float64(num))))
}

// computes the product of the terms in the array, like Sum(), but for multiplication
func Prod(a []*big.Int) *big.Int {
	prod := gb.New(1)
	for i := 0; i < len(a); i++ {
		prod.Mul(prod, a[i])
	}
	return prod
}

// calculates the sum of a given array
func Sum(a []int64) int64 {
	sum := int64(0)
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	return sum
}

// computes Sum(), except with *big.Int
func SumBig(a []*big.Int) *big.Int {
	sum := gb.New(0)
	for i := 0; i < len(a); i++ {
		sum.Add(sum, a[i])
	}
	return sum
}

// this computes Sigma_e(n), which computes the sum of the divisors of n
// where the divisors are raised to the power of e
func Sigma(n, e int64) *big.Int {
	divisors := Factors(n)
	bigdiv := CreateSlice(int64(len(divisors)))
	
	// raise each divisor to the power of e
	for i := 0; i < len(divisors); i++ {
		bigdiv[i] = gb.Pow(gb.New(divisors[i]), gb.New(e))
	}

	return SumBig(bigdiv)
}

// ================= PROBABILITY & COMBINATIONS =================
// ### Any computations related to probability/combinations
// ### Examples: Binomial, permutations, combinations

// computes the nCr(n, r) (binomial coefficient)
func Binomial(n, r int64) int64 {
	// C(n,r) = n!/((n-r)!r!), but this is inefficient
	if n < 0 || r < 0 {
		HandleError(errors.New("can't be negative"))
	} else if n < r {
		HandleError(errors.New("n cannot be less than r"))
	}

	// do the shortcut, where you can modify r based on n
	if r > n/2 {
		r = n - r
	}
	c := int64(1)
	for i := int64(1); i <= r; i++ {
		c = (n - r + i) * c / i
	}
	return c
}

// computes all combinations of n (0 <= k <= n)
func Combinations(n *big.Int) []*big.Int {
	a := CreateSlice(0)
	for k := gb.New(0); k.Cmp(n) == -1 || k.Cmp(n) == 0; k.Add(k, gb.New(1)) {
		a = append(a, C(n, k))
	}
	return a
}

// performs the same calculation as Binomial(n, k), except
// this uses big.Int instead of int64
func C(n, k *big.Int) *big.Int {
	// C(n,r) = n!/((n-r)!r!)
	if gb.Less(n, gb.New(0)) || gb.Less(k, gb.New(0)) {
		HandleError(errors.New("can't be negative"))
	} else if gb.Less(n, k) {
		HandleError(errors.New("n cannot be less than k"))
	}

	// do the shortcut, where you can modify k based on n
	if gb.Greater(k, gb.Div(n, gb.New(2))) {	// k > n/2
		k = gb.New(0).Sub(n, k)	// k = n - k
	}
	c := gb.New(1)
	for i := gb.New(1); i.Cmp(k) == -1 || i.Cmp(k) == 0; i.Add(i, gb.New(1)) {
		// c = (n - k + i) * c / i
		foo := gb.New(0).Sub(n, k)
		foo.Add(foo, i)
		c.Mul(c, foo)
		c.Div(c, i)
	}
	return c
}

// computes k-permutations of n (P(n, k) or nPr(n, r))
func P(n, k *big.Int) *big.Int {
	if gb.Greater(k, n) {
		return gb.New(0)
	}
	// returns n!/(n-k)!
	return gb.Div(Fact(n), Fact(gb.Sub(n, k)))
}

// computes all permutations of n (0 <= k <= n)
func Permutation(n *big.Int) []*big.Int {
	a := CreateSlice(0)
	for k := gb.New(0); gb.LessEqual(k, n); k.Add(k, gb.New(1)) {
		a = append(a, P(n, k))
	}
	return a
}

// ##################### DIVISORS & FACTORS #########################
// given a number num, it will compute Euler's Totient of the number
func EulerTotient(num int64) int64 {
	val := int64(0)
	for i := int64(0); i < num; i++ {
		if GCD(i, num) == 1 {
			val++
		}
	}
	return val
}

// computes Euler's Totient, but with arbitrary precision
func EulerTotientBig(num *big.Int) *big.Int {
	val := gb.Zero()
	for i := gb.Zero(); gb.Less(i, num); i = gb.Add(i, gb.New(1)) {
		if gb.Equals(GCD_big(i, num), gb.New(1)) {
			val = gb.Add(val, gb.New(1))
		}
	}
	return val
}

// given two numbers, it will calculate the greatest common divisor
func GCD(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return GCD(b, a % b)
}

// given two numbers, compute the greatest common divisor (with big.int!)
func GCD_big(a, b *big.Int) *big.Int {
	if gb.Equals(b, gb.Zero()) {
		return a
	}
	return GCD_big(b, gb.Mod(a, b))
}

// Compute the number of digits of the given number
func GetDigits(n int64) int64 {
	count := int64(0)
	for n != 0 {
		count++
		n /= 10
	}
	return count
}

// Calculate the number of factors of num
func GetFactorCount(num int64) int64 {
	// calculate factor count
	count := int64(0)
	for i := int64(1); i <= num; i++ {
		if num % i == 0 {
			count++
		}
	}
	return count
}

// finds the first digit of the number
func GetFirstDigit(n int64) int64 {
	return int64(float64(n) / math.Pow(10, float64(GetDigits(n) - 1)))
}

// ##################### PRIME CALCULATIONS #########################
// ### Any calculations related to prime numbers
// ### Examples: prime factorization, prime floor, etc.

// Calculates the prime factorization of num
func PrimeFactorization(num int64) []int64 {
	// initialize array
	primefact := make([]int64, 0)

	// the number of 2s
	for num % 2 == 0 {
		primefact = append(primefact, 2)
		num /= 2
	}

	// num is now odd. check 3s and beyond
	for i := int64(3); i * i <= num; i += 2 {
		// get factors until n is zero
		for num % i == 0 {
			primefact = append(primefact, i)
			num /= i
		}
	}

	// what if num is prime?
	if num > 2 {
		primefact = append(primefact, num)
	}

	return primefact
}

// computes the prime floor of a given number
func PrimeFloor(arr []int64, n int64) int64 {
	// lazy case
	if n == 1 {
		return n
	}

	// loop thru k (counter for exponent)
	foo := int64(0)
	found := false
	for k := int64(2); k < n; k++ {
		// check if k is a valid prime power
		for i := int64(0); i <= k; i++ {
			if arr[i] == k {
				found = true
				break
			}
			found = false
		}

		// if k is a valid prime power, perform the algo
		if found {
			foo = n
			if IsPrimePower(n, k) {
				return foo
			}
		} else {
			continue
		}
	}
	// leaving the loop means it isn't a prime power of k > 1
	// check for prime power k == 1
	if IsPrime(n) {
		return n
	} else {
		return PrimeFloor(arr, n + 1)
	}
}

// ################### MISCELLANEOUS CALCULATIONS ###################
// ### Any calculations that don't fit in the other sections

// Computes the Bernoulli Numbers using an explicit definition
// uses big.Rat b/c OEIS has a sequence of the numerators & of the denominators
// plus, there's no nonsense about float precision
func Bernoulli(n int64) *big.Rat {
	var f big.Rat
	a := make([]big.Rat, n+1)
	for m := range a {
		a[m].SetFrac64(1, int64(m+1))
		for j := m; j >= 1; j-- {
			d := &a[j-1]
			d.Mul(f.SetInt64(int64(j)), d.Sub(d, &a[j]))
		}
	}
	return f.Set(&a[0])
}

// counts the partitions of a given integer n
func CountParts(n int64) int64 {
	if n == 0 {
		return 1
	}

	// init
	p := make([]int64, n)	// stores the partitions
	k := 0					// index of last element in a partition
	p[k] = n				// first partition is n

	// loop to compute
	count := int64(0)
	for {
		// update count
		count++
		remval := int64(0)		// holds how much val can be changed

		for k >= 0 && p[k] == 1 {
			remval += p[k]
			k--
		}

		if k < 0 {	// all vals = 1 if k < 0; no more parts
			return count
		}
		p[k]--			// decr; found non-one value
		remval++		// adjust remval

		// resort array & modify remval based on the sort
		for remval > p[k] {
			p[k+1] = p[k]
			remval = remval - p[k]
			k++
		}

		// copy remval to next position
		p[k+1] = remval
		k++
	}
}

// Computes the Harmonic Number of n (i.e., H_n)
func Harmonic(n int64) *big.Float {
	sum := fzero()
	for k := int64(1); k <= n; k++ {
		kf := itof(inew(k))
		sum = fadd(sum, fdiv(fnew(1), kf))
	}
	return sum
}


// Computes the Harmonic Number of n of order k (i.e., H^(k)_n)
func HarmonicOrder(n, k int64) *big.Float {
	binom := nCr(inew(n+k-1), inew(k-1))
	kth_order_H := fmul(itof(binom), fsub(Harmonic(n+k-1), Harmonic(k-1)))
	return kth_order_H
}

// Generates the Kolakoski sequence of length seqlen
func Kolakoski(seqlen int64, numcount int64) []int64 {
	// declarations n stuff
	a := make([]int64, seqlen)
	nums := make([]int64, numcount)
	for i := int64(1); i <= numcount; i++ {
		nums[i - 1] = i
	}

	// special cases
	a[0] = nums[0]
	for i := int64(0); i < nums[0]; i++ {
		a[i+1] = nums[1]
	}

	// compute other values
	numidx := int64(1)
	count := int64(1)
	for i, j := 1 + nums[0], a[count]; i < seqlen; i, j = i+1, j-1 {
		// checks for when to reset counters/indexes
		if j <= 0 {
			i--
			count++
			j = a[count] + 1
			numidx++
			continue
		}
		if numidx >= numcount {
			numidx = 0
		}

		// assign the value
		a[i - 1] = nums[numidx]
	}
	return a
}

// calculates the change given len (# of coins), val (the value to make change for),
// and denom, the array of coin values, e.g. {1, 2, 5, 10}
func MakeChange(len int64, val int64, denom []int64) int64 {
	if val < 0 {
		return 0
	} else if val == 0 {
		return 1			// 1 way to make 0 change
	} else if len <= 0 && val >= 1 {
		return 0			// combo doesn't work
	}
	return MakeChange(len - 1, val, denom) + MakeChange(len, val - denom[len - 1], denom)
}

// computes the stirling numbers of the first kind (i.e., s(n, k))
func Stirling1(n, k int64) *big.Int {
	// handle instances for s(n, 0)
	if k == 0 {
		if n > 0 { 
			return zero() 
		} else if n == 0 {
			return inew(1)
		}
	}
	if k > n {
		return zero()
	} else {
		// otherwise, compute s(n, k) = (n-1)*s(n-1, k) + s(n-1, k-1)
		return add(mul(inew(n-1), Stirling1(n-1, k)), Stirling1(n-1, k-1))
	}
}

// Computes the stirling numbers of the second kind for n, k
func Stirling2(n, k int64) *big.Int {
	nb := inew(n)
	stir := fzero()
	for i := int64(0); i <= k; i++ {
		ib := inew(i)
		numer := mul(pow(inew(-1), inew(k-i)), pow(ib, nb))
		denom := mul(fact(inew(k-i)), fact(ib))
		stir = fadd(stir, fdiv(itof(numer), itof(denom)))
	}
	return round(stir)
}
