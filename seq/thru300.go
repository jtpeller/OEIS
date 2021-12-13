// ============================================================================
// = thru200.go																  =
// = 	Description: All OEIS sequences from A000201-A000300				  =
// = 	Note: Not all sequences in this range have been programmed			  =
// = 	Date: December 12, 2021												  =
// = 	Last Update: December 12, 2021										  =
// ============================================================================

package seq

import (
	"OEIS/utils"
	"math"
	"math/big"
)

const (
	LONG_A000205 = 10
)

/**
 * A000201 computes the Lower Wythoff sequence (a Beatty sequence):
 *  a(n) = floor(n*phi), where phi = (1+sqrt(5))/2 = A001622
 * Date: December 12, 2021 	Confirmed working: December 12, 2021
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
 * Date: December 12, 2021 	Confirmed working: December 12, 2021
 * Link: https://oeis.org/A000202
 */
func A000202(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	b := []int64{1, 3, 4, 6, 8, 9, 11, 12}
	
	// otherwise, we can populate without worry
	for n := 0; n < len(b); n++ {
		a[n] = b[n]
	}
	
	// handle if the user gives a seqlen < 8
	if seqlen < 8 {
		return a, 1
	}
	
	// loop to populate
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
 * Date: December 12, 2021	Confirmed working: December 12, 2021
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
 * Date: December 12, 2021 	Confirmed working: December 12, 2021
 * Link: https://oeis.org/A000204
 */
func A000204(seqlen int64) ([]*big.Int, int64) {
	a, _ := A000032(seqlen+1)
	return a[1:], 1
}

/**
 * A000205 computes the # of positive integers <= 2^n of the form x^2 + 3y^2
 * Date: December 12, 2021 	Confirmed working: December 12, 2021
 * Link: https://oeis.org/A000205
 */
func A000205(seqlen int64) ([]*big.Int, int64) {
	if seqlen > LONG_A000205 {
		utils.LongCalculationWarning("A000205", LONG_A000205)
	}

	a := utils.CreateSlice(seqlen)
	a[0] = New(1)
	for n := int64(1); n < seqlen; n++ {
		nf := float64(n)
		count := a[n-1]
		for k := int64(math.Pow(2, nf-1) + 1); k <= int64(math.Pow(2, nf)); k++ {
			if utils.IsRepr(k, 1, 3) {
				count = Add(count, New(1))
			}
		}
		a[n] = count
	}
	return a, 0
}