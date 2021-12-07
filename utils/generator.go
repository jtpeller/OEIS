// ============================================================================
// = generator.go															  =
// = 	Description: Useful generator functions like primes.				  =
// = 	Date: December 07, 2021												  =
// ============================================================================

package utils

import "errors"

// ########################## GENERATOR FUNCTIONS #############################
// ### given a number, it will generate a sequence with some quality up to that
// ### number. things like primes, evens, odds, etc.

// generates the sequence of primes; count = num
func Primes(seqlen int64) ([]int64, error) {
	// check valid num
	if seqlen <= 0 {
		return nil, errors.New("utils error in Primes(): num cannot be negative")
	}

	primes := make([]int64, 0)
	num := int64(0)
	for i := int64(0); i < seqlen; {
		if IsPrime(int64(num)) {
			primes = append(primes, num)
			i++
		}
		num++
	}
	return primes, nil
}

