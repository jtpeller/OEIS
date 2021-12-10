// ============================================================================
// = generator.go															  =
// = 	Description: Useful generator functions like primes.				  =
// = 	Date: December 07, 2021												  =
// ============================================================================

package utils

// ########################## GENERATOR FUNCTIONS #############################
// ### given a number, it will generate a sequence with some quality up to that
// ### number. things like primes, evens, odds, etc.

// Computes the factors of num
func Factors(num int64) []int64 {
	factors := make([]int64, 0)
	for i := int64(1); i <= num; i++ {
		if num % i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}


// Calculates the isqrt of an array
func Isqrtarray(arr []int64) []int64 {
	a := make([]int64, len(arr))
	for i := 0; i < len(arr); i++ {
		val := Isqrt(arr[i])
		a[i] = val
	}
	return a
}

// generates the sequence of primes; count = num
func Primes(seqlen int64) []int64 {
	primes := make([]int64, 0)
	num := int64(0)
	for i := int64(0); i < seqlen; {
		if IsPrime(num) {
			primes = append(primes, num)
			i++
		}
		num++
	}
	return primes
}
