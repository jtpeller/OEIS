// ============================================================================
// = builder.go
// = 	Description		Builds/converts things for big.Int.
// = 	Date			December 09, 2021
// ============================================================================

package utils

import "math/big"

// enables usage of big.Int by initializing a slice of len
func CreateSlice(len int64) []*big.Int {
	slice := make([]*big.Int, 0)
	for i := int64(0); i < len; i++ {
		slice = append(slice, big.NewInt(0))
	}
	return slice
}


// converts from []int64 slice to []*big.Int
func ToBigSlice(slice []int64) []*big.Int {
	a := CreateSlice(int64(len(slice)))
	for i := 0; i < len(slice); i++ {
		a[i] = big.NewInt(slice[i])
	}
	return a
}

// converts from []*big.Int to []int64
func ToIntSlice(slice []*big.Int) []int64 {
	a := make([]int64, 0)
	for i := 0; i < len(slice); i++ {
		a = append(a, slice[i].Int64())
	}
	return a
}
