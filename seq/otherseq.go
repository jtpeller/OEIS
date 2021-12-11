package seq

import (
	"OEIS/utils"
	"math"
	"math/big"
)

/**
 * A003048 computes a[n+1]=n*a[n] - (-1)^n
 * Date: December 10, 2021	Confirmed working: December 10, 2021
 * Link: https://oeis.org/A003048
 */
func A003048(seqlen int64) ([]*big.Int, int64) {
	a := utils.CreateSlice(seqlen)
	a[0] = New(1)
	for i := int64(1); i < seqlen; i++ {
		a[i] = Sub(Mul(New(i), a[i - 1]), Pow(New(-1), New(i)))
	}
	return a, 0
}

func A007947(max int64) ([]int64, int64) {
	a := make([]int64, 0)
	for i := int64(0); i < max; i++ {
		// calculate the prime factorization of i
		pfact := utils.PrimeFactorization(i)

		// strip all non unique elements
		set := make(map[int64]bool)
		for j := int64(0); j < int64(len(pfact)); j++ {
			set[int64(pfact[j])] = true
		}

		// get the radical
		radical := int64(1)
		for key := range set {
			radical *= key
		}
		a = append(a, radical)
	}
	return a, 1
}

/**
 * A032346 essentially the same as A000108, except row starts at 1
 *  instead of 0
 * Date: December 07, 2021
 * Link: https://oeis.org/A032346
 */
 func A032346(seqlen int64) ([]int64, int64) {
	// init
	a := make([]int64, seqlen)		// the seq
	a[0] = 1
	old := make([]int64, seqlen)	// last row
	new := make([]int64, seqlen)	// new row
	old[0] = 1

	// compute each row & store into a
	row, col := int64(1), int64(0)
	for ; row < seqlen; row++ {
		col = 0

		// calculate new row
		for  ; col < row; col++ {
			new[col + 1] = new[col] + old[col]
		}

		// copy down
		if col > 0 {
			for i := int64(0); i < col + 1; i++ {
				old[i] = new[i]
				new[i] = 0		// erase new row
			}
		}

		// copy the last element
		new[0] = old[col]		// overwrite first elem
		a[row] = old[col]		// copy last elem of old
	}
	return a, 0
}

// works, up to seqlen = 11; after that the nums are too big
// discovered accidentally when attempting A000108
func A088218(seqlen int64) ([]int64, int64) {
	a := make([]int64, seqlen)
	for i := int64(0); i < seqlen; i++ {
		a[i] = utils.Fact(2*i)/(utils.Fact(i) * utils.Fact(i+1))
	}
	return a, 0
}

// inaccurate due to rounding errors in float64 <-> int64
// discovered accidentally when attemptiung A000111
func A137590 (seqlen int64) ([]int64, int64) {
	if seqlen > OVERFLOW_A000111 {
		utils.OverflowError("A000111", OVERFLOW_A000111)
	}
	
	a := make([]int64, seqlen)
	for i := int64(0); i < seqlen; i++ {
		a[i] = int64(math.Round(2.0 * math.Pow(2.0 / math.Pi, float64(i + 1)) * float64(utils.Fact(i))))
	}
	return a, 0
}
