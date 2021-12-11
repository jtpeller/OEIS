// ============================================================================
// = calc.go																  =
// = 	Description: Useful calculation functions like divisors, GCD, etc.	  =
// = 	Date: December 07, 2021												  =
// ============================================================================

package utils

import (
	"errors"
	"math"
	"math/big"
)

// ########################## CALCULATIONS #############################
// ### this section calculates some property of a number (factor count,
// ### prime factorization, etc.)

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

// counts the partitions of a given integer n
func CountParts(n int64) int64 {
	if n == 0 {
		return 1
	}

	// initializations
	parts := make([]int64, n)
	k := 0
	parts[k] = n

	// loop to generate partitions
	partcount := int64(0)
	for {
		// update partition count
		partcount++

		// find rightmost non-one value
		remval := int64(0)			// holds how much val can be changed
		for k >= 0 && parts[k] == 1 {
			remval += parts[k]
			k--
		}

		// if k < 0, all vals = 1; no more partitions
		if k < 0 {
			return partcount
		}
		
		// update values
		parts[k]--
		remval++

		// resort array. modify remval based on sort
		for remval > parts[k] {
			parts[k+1] = parts[k]
			remval -= parts[k]
			k++
		}

		// copy remval to next position & increment k
		parts[k+1] = remval
		k++
	}
}


// given a number num, it will compute Euler's Totient of the number
func EulerTotient(num int64) int64 {
	val := int64(0)
	for i := int64(0); i < num; i++ {
		if GreatestCommonDivisor(i, num) == 1 {
			val++
		}
	}
	return val
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

// finds the first digit of the number
func GetFirstDigit(n int64) int64 {
	return int64(float64(n) / math.Pow(10, float64(GetDigits(n) - 1)))
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

// given two numbers, it will calculate the greatest common divisor
func GreatestCommonDivisor(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return GreatestCommonDivisor(b, a % b)
}

// computes the factorial of num
func Fact(num int64) int64 {
	if num < 0 {
		HandleError(errors.New("factorial of a negative number is undefined"))
	}
	prod := int64(1)
	for i := int64(1); i < num; i++ {
		prod *= i
	}
	return prod
}

// compute the factorial of a num (big.Int)
func Factorial(num *big.Int) *big.Int {
	if num.Cmp(big.NewInt(0)) == -1 {
		HandleError(errors.New("factorial of a negative number is undefined"))
	}

	prod := big.NewInt(1)
	for i := big.NewInt(1); i.Cmp(num) == -1 || i.Cmp(num) == 0; i.Add(i, big.NewInt(1)) {
		prod.Mul(prod, i)
	}
	return prod
}

// Calculate the integer square root of a number
func Isqrt(num int64) (int64) {
	return int64(math.Floor(math.Sqrt(float64(num))))
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

// computes the product of the terms in the array, like Sum(), but for multiplication
func Prod(a []*big.Int) *big.Int {
	prod := big.NewInt(1)
	for i := 0; i < len(a); i++ {
		prod.Mul(prod, a[i])
	}
	return prod
}

// calculates the sum of a given array; essentially,
// this computes Sigma(a_i), 0 <= i < len(a)
func Sum(a []int64) int64 {
	sum := int64(0)
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	return sum
}

